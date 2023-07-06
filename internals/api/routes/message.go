package routes

import (
	"hermes/internals/api/handlers"
	"hermes/internals/services"

	"github.com/gofiber/fiber/v2"
)

func MessageRouter(app fiber.Router, publisher *services.Relayer) {
	app.Post("/message", handlers.PostMessage(publisher))
}
