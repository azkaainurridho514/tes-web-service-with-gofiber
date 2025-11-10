package router

import (
	"github.com/azkaainurridho514/tes-web-service/database"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/users", func(c *fiber.Ctx) error {
		rows, err := database.DB.Query("SELECT id, name, email FROM users")
		if err != nil {
			return c.Status(500).SendString("Query error: " + err.Error())
		}
		defer rows.Close()

		var users []map[string]interface{}

		for rows.Next() {
			var id int
			var name, email string

			if err := rows.Scan(&id, &name, &email); err != nil {
				return c.Status(500).SendString("Scan error: " + err.Error())
			}

			users = append(users, fiber.Map{
				"id":    id,
				"name":  name,
				"email": email,
			})
		}

		return c.JSON(users)
	})

}
