package error_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Handler(c *fiber.Ctx, err error) error {
	log.Error().
		Err(err).
		Msg("global api handler")

	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}
