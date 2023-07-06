package handlers

import (
	"hermes/internals/api/presenter"
	"hermes/internals/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostMessage(publisher services.IPublisher) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody presenter.Message
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MessageErrorResponse(uuid.Nil, err))
		}

		transactionId, err := publisher.Publish(requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MessageErrorResponse(transactionId, err))
		}

		return c.JSON(presenter.MessageSuccessResponse(transactionId, &requestBody))
	}
}
