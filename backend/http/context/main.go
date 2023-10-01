package context

import "github.com/gofiber/fiber/v2"

const (
	DbKey    string = "db"
	LimitKey        = "limit"
)

func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}

func GetLocal[T any](c *fiber.Ctx, key string) T {
	return c.Locals(key).(T)
}
