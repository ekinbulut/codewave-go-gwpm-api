package handlers

import (
	"hermes/internals/api/presenter"
	"hermes/internals/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MessageHandler struct {
	publisher services.IPublisher
}

func NewMessageHandler(publisher services.IPublisher) *MessageHandler {
	return &MessageHandler{
		publisher: publisher,
	}
}

func (h *MessageHandler) PostMessage() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody presenter.Message
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MessageErrorResponse(uuid.Nil, err))
		}

		transactionId, err := h.publisher.Publish(requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MessageErrorResponse(transactionId, err))
		}

		return c.JSON(presenter.MessageSuccessResponse(transactionId, &requestBody))
	}
}
