package user_handler

import (
	"github.com/gofiber/fiber/v2"
)

func (i *Handler) RobotsTXT(c *fiber.Ctx) error {
	return c.SendString("")
}
