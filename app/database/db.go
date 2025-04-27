package database

import (
	"database/sql"
	"eco-backend/config"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASSWORD", "")
	dbname := config.GetEnv("DB_NAME", "eco_db")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=public",
		user, password, host, port, dbname,
	)

	log.Printf("Connecting to database: %s\n", dsn)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("❌ Błąd otwarcia połączenia: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Nieudane pingowanie bazy: %v", err)
	}

	DB = db
	log.Println("✅ Połączono z bazą danych (pgx/stdlib)")
}
