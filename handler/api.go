package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Main(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func Healthcheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}
