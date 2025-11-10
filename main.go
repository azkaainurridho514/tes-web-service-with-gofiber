package main

import (
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
	router.SetupRoutes(app)
		app.Use(func(c *fiber.Ctx) error {
	return c.SendStatus(404) 
	})
	app.Listen(":8080")
}