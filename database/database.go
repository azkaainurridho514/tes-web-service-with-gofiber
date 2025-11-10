package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL tidak ditemukan")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Gagal open DB:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Gagal connect ke DB:", err)
	}

	log.Println("✅ Database connected!")
}
