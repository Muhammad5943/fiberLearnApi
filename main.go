package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// start new fiber App()
	app := fiber.New()

	// Retrieve data example
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API in UP!")

		return err
	})

	// listen app on port 3000
	app.Listen(":3000")
}
