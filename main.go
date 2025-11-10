package main

import (
	"log"
	"os"

	"github.com/azkaainurridho514/tes-web-service/database"
	"github.com/azkaainurridho514/tes-web-service/router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)
func main() {
	app := fiber.New()

	database.Connect()

	router.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server running on port:", port)
	log.Fatal(app.Listen(":" + port))
}