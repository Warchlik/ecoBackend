package database

import (
	"fmt"
	"log"

	"eco-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASSWORD", "")
	dbname := config.GetEnv("DB_NAME", "eco_db")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=public",
		host, port, user, password, dbname)

	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Nie można połączyć się z bazą danych: ", err)
	}

	log.Println(db)

	DB = db
	log.Println("✅ Połączono z bazą danych")
}
