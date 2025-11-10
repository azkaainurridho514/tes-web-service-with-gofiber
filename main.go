package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)
func main() {
	app := fiber.New()

	// Ambil URL dari Railway
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL tidak ditemukan")
	}

	// Connect ke Neon
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Gagal open DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Gagal connect ke DB:", err)
	}

	log.Println("✅ Database connected!")

	// API test
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello! DB connected.")
	})

	// PORT Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server running on port:", port)
	log.Fatal(app.Listen(":" + port))
}