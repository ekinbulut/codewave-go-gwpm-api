package presenter

import "github.com/gofiber/fiber/v2"

type Message struct {
	Message string `json:"message"`
}

func MessageSuccessResponse(message *Message) *fiber.Map {
	return &fiber.Map{
		"data":   message,
		"status": true,
		"error":  nil,
	}
}

func MessageErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":   nil,
		"status": false,
		"error":  err.Error(),
	}
}
