package v1

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/hanifmaliki/rig-checklist-backend/internal/service"
)

func FileRouter(app fiber.Router, jwt *func(*fiber.Ctx) error) {
	app.Get("/", getFileList())
	app.Get("/:filename", getFile())

	app.Use(*jwt)
	app.Post("/", postFile())
	app.Delete("/:filename", deleteFile())
}

func getFileList() fiber.Handler {
	return func(c *fiber.Ctx) error {
		prefix := c.Query("prefix")

		data, err := service.Instance().GetFileListFromMinio(prefix)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    data,
			"error":   nil,
		})
	}
}

func getFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		filename := c.Params("filename")
		filename, err := url.QueryUnescape(filename)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		contentDisposition := c.Query("content_disposition")
		if contentDisposition == "" {
			contentDisposition = "inline" // inline == open, attachment == download
		}

		file, err := service.Instance().GetFileFromMinio(filename)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		b, err := io.ReadAll(file)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		var mimeType string
		if filepath.Ext(filename) != ".svg" {
			mimeType = http.DetectContentType(b)
		} else {
			mimeType = "image/svg+xml"
		}

		c.Set("Accept-Ranges", "bytes")
		c.Set("Content-Description", "File Transfer")
		c.Set("Content-Type", mimeType)
		c.Set("Content-Disposition", contentDisposition+"; filename="+filename)
		c.Set("Content-Transfer-Encoding", "binary")
		c.Set("Expires", "0")
		c.Set("Cache-Control", "must-revalidate")
		c.Set("Pragma", "public")

		return c.Send(b)
	}
}

func postFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		filename := c.FormValue("filename")
		if filename == "" {
			filename = file.Filename
		}
		err = service.Instance().PutFileToMinio(filename, file)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"message": "File has been successfully uploaded",
			"error":   nil,
		})
	}
}

func deleteFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		filename := c.Params("filename")
		filename, err := url.QueryUnescape(filename)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		err = service.Instance().DeleteFileFromMinio(filename)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"message": fmt.Sprintf("File %s has been successfully deleted", filename),
			"error":   nil,
		})
	}
}
