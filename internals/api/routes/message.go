package routes

import (
	"hermes/internals/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func MessageRouter(router fiber.Router, handlers *handlers.MessageHandler) {
	router.Post("/message", handlers.PostMessage())
}
