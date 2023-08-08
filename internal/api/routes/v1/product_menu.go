package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"
)

func ProductMenuRouter(app fiber.Router, jwt *func(*fiber.Ctx) error) {
	app.Get("/", getProductMenus())
	app.Get("/:id", getProductMenu())

	app.Use(*jwt)
	app.Post("/", postProductMenu())
	app.Put("/:id", putProductMenu())
	app.Delete("/:id", deleteProductMenu())
}

func getProductMenus() fiber.Handler {
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

		orderBy := c.Query("order_by")
		fetched, err := service.Instance().ReadProductMenus(conds, orderBy)
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

func getProductMenu() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		id := uint(id64)
		fetched, err := service.Instance().ReadProductMenu(id)
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

func postProductMenu() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		var requestBody model.ProductMenu
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().CreateProductMenu(user, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Product menu has been successfully created",
			"error":   nil,
		})
	}
}

func putProductMenu() fiber.Handler {
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
		var requestBody model.ProductMenu
		err = c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().UpdateProductMenu(user, id, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Product menu has been successfully updated",
			"error":   nil,
		})
	}
}

func deleteProductMenu() fiber.Handler {
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
		err = service.Instance().DeleteProductMenu(user, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    nil,
			"message": "Product menu has been successfully deleted",
			"error":   nil,
		})
	}
}
