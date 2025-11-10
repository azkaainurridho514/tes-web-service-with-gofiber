package router

import (
	"github.com/azkaainurridho514/employee_presence/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", handler.SayHello)
}