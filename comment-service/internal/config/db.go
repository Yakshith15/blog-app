package config

import (
	"database/sql"
	"os"
	"log"
)
func ConnectDB() *sql.DB {
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("DB_URL is not set")
	}
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("failed to open database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping database:", err)
	}
	log.Println("connected to database")
	return db
}