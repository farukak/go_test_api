package router

import (
	"github.com/farukak/go_test_api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// api := app.Group("/", logger.New())
	// api.Get("/", handler.Main)
	// api.Get("/healthcheck", handler.Healthcheck)

	// GET /@v1
	app.Get("/", handler.Main)

}
