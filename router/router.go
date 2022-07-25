package router

import (
	noteRoutes "github.com/Muhammad5943/fiberLearnApi/internal/routes/note"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	// Group api calls with param '/api'
	api := app.Group("/api", logger.New())

	// Setup note routes, can use same syntax to add routes for more models
	noteRoutes.SetupNoteRoutes(api)
}
