package v1

import (
	"crypto/rsa"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/service"
)

func AuthRouter(app fiber.Router, privateKey *rsa.PrivateKey) {
	app.Post("/login", postLogin(privateKey))
}

func postLogin(privateKey *rsa.PrivateKey) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody model.Login
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		var validate = validator.New()
		err = validate.Struct(requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		token, err := service.Instance().Login(&requestBody)
		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		// Generate encoded token and send it as response.
		t, err := token.SignedString(privateKey)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"token":   t,
			"message": "Login success",
			"error":   nil,
		})
	}
}
