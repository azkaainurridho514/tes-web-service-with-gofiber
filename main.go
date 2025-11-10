package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Railway + Fiber!")
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