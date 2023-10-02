package http

import (
	"context"
	"fmt"
	"time"

	"github.com/eonianmonk/go-rate-limit/backend/config"
	http_context "github.com/eonianmonk/go-rate-limit/backend/http/context"
	"github.com/eonianmonk/go-rate-limit/data"
	"github.com/gofiber/fiber/v2"
)

type Task func(context.Context, chan error)

func Run(cfg config.Config, ctx context.Context) {
	tasks := []Task{startFiber(cfg), startPurgeLoop(cfg)}
	errc := make(chan error)
	ctx, cancel := context.WithCancel(ctx)
	for _, task := range tasks {
		go task(ctx, errc)
	}
	err := <-errc
	fmt.Printf("failed at run: %s", err.Error())
	cancel()
	close(errc)
}

func startFiber(cfg config.Config) Task {
	q := data.New(cfg.DB())

	app := fiber.New()
	SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		http_context.SetLocal[*data.Queries](c, http_context.DbKey, q)
		http_context.SetLocal[*int16](c, http_context.LimitKey, cfg.Limit())
		return nil
	})

	return func(ctx context.Context, errc chan error) {
		err := app.Listener(cfg.Listen())
		if err != nil {
			errc <- fmt.Errorf("fiber listener failed: %s", err.Error())
		}
		<-ctx.Done()
		err = app.Shutdown()
		if err != nil {
			errc <- err
		}
	}
}

func startPurgeLoop(cfg config.Config) Task {
	q := data.New(cfg.DB())

	return func(ctx context.Context, errc chan error) {
		for {
			time.Sleep(time.Minute)
			err := q.DeleteOld(ctx)
			if err != nil {
				errc <- fmt.Errorf("old rates clearing failed: %s", err.Error())
			}
		}
	}
}
