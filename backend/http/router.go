package http

import (
	handlers "github.com/eonianmonk/go-rate-limit/backend/http/api_handlers/v1"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/rate", handlers.GetRate)
}
