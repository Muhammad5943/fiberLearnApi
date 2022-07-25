package main

import (
	"github.com/Muhammad5943/fiberLearnApi/database"
	"github.com/Muhammad5943/fiberLearnApi/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// start new fiber App()
	app := fiber.New()

	// Connect data to connectDB
	database.ConnectDB()

	// Retrieve data example
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API in UP!")

		return err
	})

	// Setup the router
	router.SetupRoutes(app)

	// listen app on port 3000
	app.Listen(":3000")
}
