package v1

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
)

func HiRouter(app fiber.Router, jwt *func(*fiber.Ctx) error) {
	app.Get("/", *jwt, sayHi())
}

func sayHi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    user,
			"error":   nil,
		})
	}
}
