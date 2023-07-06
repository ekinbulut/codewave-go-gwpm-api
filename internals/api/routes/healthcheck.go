package routes

import "github.com/gofiber/fiber/v2"

func HealthCheck(app fiber.Router) {
	app.Get("/hc", func(c *fiber.Ctx) error {
		return c.JSON("OK")
	})
}
