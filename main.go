package main

import (
	"log"
	"os"

	"github.com/azkaainurridho514/employee_presence/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)
func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Route not found",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" 
	}

	log.Println("Server running on port:", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}