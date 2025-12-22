package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL is not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	log.Println("connected to database")
	return db
}
