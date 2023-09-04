package middleware

import (
	"crypto/rsa"
	"os"

	"github.com/hanifmaliki/rig-checklist-backend/internal/helper"
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func Jwt(privateKey *rsa.PrivateKey, devMode bool) fiber.Handler {
	if !devMode {
		return jwtware.New(jwtware.Config{
			ContextKey:    "claims",
			SigningMethod: "RS256",
			SigningKey:    privateKey.Public(),
			SuccessHandler: func(c *fiber.Ctx) error {
				claims := c.Locals("claims").(*jwt.Token)
				claimsObject := claims.Claims.(jwt.MapClaims)

				user := model.User{}
				user.Username = claimsObject["username"].(string)
				user.Name = claimsObject["name"].(string)
				user.Email = claimsObject["email"].(string)

				c.Locals("user", &user)
				return c.Next()
			},
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			},
		})
	} else {
		return func(c *fiber.Ctx) error {
			user := model.User{}
			username := "admin"
			if os.Getenv("ADMIN_USERNAME") != "" {
				username = os.Getenv("ADMIN_USERNAME")
			}
			user.Username = username
			user.Name = helper.UserDummy.Name
			user.Email = helper.UserDummy.Email

			c.Locals("user", &user)

			return c.Next()
		}
	}
}
