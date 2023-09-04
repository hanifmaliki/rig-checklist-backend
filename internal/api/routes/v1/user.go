package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hanifmaliki/rig-checklist-backend/internal/model"
	"github.com/hanifmaliki/rig-checklist-backend/internal/service"
)

func UserRouter(app fiber.Router) {
	app.Get("/", getUsers())
	app.Get("/:id", getUser())

	app.Post("/", postUser())
	app.Put("/:id", putUser())
	app.Delete("/:id", deleteUser())
}

func getUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		conds := map[string]interface{}{}

		isActiveString := c.Query("is_active")
		var isActive *bool
		if isActiveString != "" {
			isActiveBool, err := strconv.ParseBool(isActiveString)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}
			isActive = &isActiveBool
		}
		if isActive != nil {
			conds["is_active"] = isActive
		}

		fetched, err := service.Instance().ReadUsers(conds, c.Query("sort_by"))
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

func getUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		conds := map[string]interface{}{}

		id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		conds["id"] = uint(id64)

		fetched, err := service.Instance().ReadUser(conds)
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

func postUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		var requestBody model.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().CreateUser(user, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "User has been successfully created",
			"error":   nil,
		})
	}
}

func putUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		id := uint(id64)

		var requestBody model.User
		err = c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().UpdateUser(user, id, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "User has been successfully updated",
			"error":   nil,
		})
	}
}

func deleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		id := uint(id64)

		err = service.Instance().DeleteUser(user, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    nil,
			"message": "User has been successfully deleted",
			"error":   nil,
		})
	}
}
