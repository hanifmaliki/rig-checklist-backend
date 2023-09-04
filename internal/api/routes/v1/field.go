package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"
)

func FieldRouter(app fiber.Router) {
	app.Get("/", getFields())
	app.Get("/:id", getField())

	app.Post("/", postField())
	app.Put("/:id", putField())
	app.Delete("/:id", deleteField())
}

func getFields() fiber.Handler {
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

		fetched, err := service.Instance().ReadFields(conds, c.Query("sort_by"))
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

func getField() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		id := uint(id64)

		fetched, err := service.Instance().ReadField(id)
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

func postField() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		var requestBody model.Field
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().CreateField(user, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Field has been successfully created",
			"error":   nil,
		})
	}
}

func putField() fiber.Handler {
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

		var requestBody model.Field
		err = c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().UpdateField(user, id, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Field has been successfully updated",
			"error":   nil,
		})
	}
}

func deleteField() fiber.Handler {
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

		err = service.Instance().DeleteField(user, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    nil,
			"message": "Field has been successfully deleted",
			"error":   nil,
		})
	}
}
