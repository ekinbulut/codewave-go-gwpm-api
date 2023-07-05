package routes

import (
	"hermes/internals/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func MessageRouter(app fiber.Router) {
	app.Post("/message", handlers.PostMessage())
}
