package handlers

import (
	"hermes/internals/api/presenter"
	"hermes/internals/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PostMessage(publisher services.IPublisher) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody presenter.Message
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MessageErrorResponse(err))
		}

		err = publisher.Publish(requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.MessageErrorResponse(err))
		}

		return c.JSON(presenter.MessageSuccessResponse(&requestBody))
	}
}
