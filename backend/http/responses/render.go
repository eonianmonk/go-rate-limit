package responses

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, val interface{}, status int) error {
	c.Set("content-type", "application/json")
	err := c.Status(status).JSON(val)
	if err != nil {
		return fmt.Errorf("failed to render response: %s", err.Error())
	}
	return nil
}

func RenderErr(c *fiber.Ctx, err error, status int) {
	body := newErrorResponseE(err)
	Render(c, body, status)
}
