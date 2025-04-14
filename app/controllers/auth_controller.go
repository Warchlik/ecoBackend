package controllers

import (
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

	var existing models.User
	database.DB.Where("email = ?", req.Email).First(&existing)
	if existing.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email już istnieje"})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := database.DB.Create(&user).Error; err != nil {
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

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowe dane logowania"})
		return
	}

	// Sprawdzenie hasła
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowe dane logowania"})
		return
	}

	// TODO: Wygeneruj JWT tutaj
	c.JSON(http.StatusOK, gin.H{"message": "Zalogowano pomyślnie"})
}
