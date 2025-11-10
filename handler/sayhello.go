package handler

import (
	"github.com/gofiber/fiber/v2"
)

func SayHello(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message":"Hello"})
}
