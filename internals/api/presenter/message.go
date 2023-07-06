package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Message struct {
	Numbers  []string `json:"numbers"`
	Template struct {
		Name string `json:"name"`
	} `json:"template"`
}

func MessageSuccessResponse(id uuid.UUID, message *Message) *fiber.Map {
	return &fiber.Map{
		"id":     id,
		"data":   message,
		"status": true,
		"error":  nil,
	}
}

func MessageErrorResponse(id uuid.UUID, err error) *fiber.Map {
	return &fiber.Map{
		"id":     id,
		"data":   nil,
		"status": false,
		"error":  err.Error(),
	}
}
