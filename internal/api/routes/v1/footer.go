package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"
)

func FooterRouter(app fiber.Router, jwt *func(*fiber.Ctx) error) {
	app.Get("/", getFooter())
	app.Put("/", *jwt, putFooter())
}

func getFooter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.Instance().ReadFooter()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    fetched,
			"error":   nil,
		})
	}
}

func putFooter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		var requestBody model.Footer
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().UpdateFooter(user, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Footer has been successfully updated",
			"error":   nil,
		})
	}
}
