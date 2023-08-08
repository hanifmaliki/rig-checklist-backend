package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/helper"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/model"
	"gitlab.todcoe.com/todcoe/petros-website/corporate-website-minerva/internal/service"
)

func HomeContentRouter(app fiber.Router, jwt *func(*fiber.Ctx) error) {
	app.Get("/", getHomeContents())

	app.Use(*jwt)
	app.Put("/", putHomeContents())
	app.Delete("/:id", deleteHomeContent())
}

func getHomeContents() fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := &model.HomeContent{
			Section: c.Query("section"),
		}

		fetched, err := service.Instance().ReadHomeContents(query)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		fetchedMap, err := helper.HomeContentSliceToMapSlice(fetched)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    fetchedMap,
			"error":   nil,
		})
	}
}

func putHomeContents() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		var requestBody []*model.HomeContent
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		result, err := service.Instance().UpdateHomeContents(user, requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    result,
			"message": "Home content has been successfully updated",
			"error":   nil,
		})
	}
}

func deleteHomeContent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id64, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		id := uint(id64)
		err = service.Instance().DeleteHomeContent(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    nil,
			"message": "Home content has been successfully deleted",
			"error":   nil,
		})
	}
}
