package healthz_router

import (
	"github.com/gofiber/fiber/v2"
)

func HealthzRouter(app fiber.Router) {
	app.Get("/healthz", getHealth())
}

func getHealth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("ok")
	}
}
