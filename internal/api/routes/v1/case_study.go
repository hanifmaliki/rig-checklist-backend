package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"
)

func CaseStudyRouter(app fiber.Router, jwt *func(*fiber.Ctx) error) {
	app.Get("/", getCaseStudies())
	app.Get("/:param", getCaseStudy())

	app.Use(*jwt)
	app.Post("/", postCaseStudy())
	app.Put("/:id", putCaseStudy())
	app.Delete("/:id", deleteCaseStudy())
}

func getCaseStudies() fiber.Handler {
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

		fetched, err := service.Instance().ReadCaseStudies(conds)
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

func getCaseStudy() fiber.Handler {
	return func(c *fiber.Ctx) error {
		conds := map[string]interface{}{}

		param := c.Params("param")
		id64, err := strconv.ParseUint(param, 10, 64)
		if err != nil {
			conds["slug"] = param
		} else {
			conds["id"] = uint(id64)
		}

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

		fetched, err := service.Instance().ReadCaseStudy(conds)
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

func postCaseStudy() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		var requestBody model.CaseStudy
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().CreateCaseStudy(user, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Case study has been successfully created",
			"error":   nil,
		})
	}
}

func putCaseStudy() fiber.Handler {
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
		var requestBody model.CaseStudy
		err = c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().UpdateCaseStudy(user, id, &requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Case study has been successfully updated",
			"error":   nil,
		})
	}
}

func deleteCaseStudy() fiber.Handler {
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
		err = service.Instance().DeleteCaseStudy(user, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    nil,
			"message": "Case study has been successfully deleted",
			"error":   nil,
		})
	}
}
