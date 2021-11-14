package main

import (
	"log"

	"github.com/farukak/go_test_api/database"
	"github.com/farukak/go_test_api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Use an external setup function in order
	// to configure the app in tests as well
	app := Setup()

	// start the application on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {

	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// Initialize a new app
	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app)

	// Return the configured app
	return app
}
