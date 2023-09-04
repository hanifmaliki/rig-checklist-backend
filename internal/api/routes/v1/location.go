package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"
)

func LocationRouter(app fiber.Router) {
	app.Get("/", getLocations())
	app.Get("/:id", getLocation())

	app.Post("/", postLocation())
	app.Put("/:id", putLocation())
	app.Delete("/:id", deleteLocation())
}

func getLocations() fiber.Handler {
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

		fetched, err := service.Instance().ReadLocations(conds, c.Query("sort_by"))
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

func getLocation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		id := uint(id64)

		fetched, err := service.Instance().ReadLocation(id)
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

func postLocation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		var requestBody model.Location
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().CreateLocation(user, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Location has been successfully created",
			"error":   nil,
		})
	}
}

func putLocation() fiber.Handler {
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

		var requestBody model.Location
		err = c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().UpdateLocation(user, id, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Location has been successfully updated",
			"error":   nil,
		})
	}
}

func deleteLocation() fiber.Handler {
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

		err = service.Instance().DeleteLocation(user, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    nil,
			"message": "Location has been successfully deleted",
			"error":   nil,
		})
	}
}
