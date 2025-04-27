package controllers

import (
	"database/sql"
	"eco-backend/app/database"
	"eco-backend/app/models"
	"eco-backend/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req requests.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.DB

	var existingID int
	err := db.QueryRow(`SELECT id FROM users WHERE email = $1`, req.Email).Scan(&existingID)

	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd bazy danych"})
		return
	}

	if existingID >= 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email już istnieje"})
		return
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd hashowania hasła"})
		return
	}

	_, err = db.Exec(
		`INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`,
		req.Name, req.Email, string(hashedPwd),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się utworzyć użytkownika"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Zarejestrowano pomyślnie"})
}

func Login(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.DB

	var user models.User
	err := db.QueryRow(`SELECT id, name, email, password FROM users WHERE email = $1`, req.Email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowe dane logowania"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd bazy danych"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowe dane logowania"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Zalogowano pomyślnie"})
}
