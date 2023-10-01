package http

import (
	"github.com/eonianmonk/go-rate-limit/backend/config"
	"github.com/eonianmonk/go-rate-limit/backend/http/context"
	"github.com/eonianmonk/go-rate-limit/data"
	"github.com/gofiber/fiber/v2"
)

func Run(cfg config.Config) {
	// TODO:
	// endpoints
	// delete old limits
}

func startFiber(cfg config.Config) error {
	q := data.New(cfg.DB())

	app := fiber.New()
	SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		context.SetLocal[*data.Queries](c, context.DbKey, q)
		context.SetLocal[*int16](c, context.LimitKey, cfg.Limit())
		return nil
	})

	return nil
}
