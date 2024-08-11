package meta_handler

import (
	"github.com/gofiber/fiber/v2"
)

func (i *Handler) GetHostMeta(c *fiber.Ctx) error {
	if c.Accepts(fiber.MIMEApplicationJSON) != "" {
		return i.GetHostMetaJSON(c)
	}

	if c.Accepts(fiber.MIMEApplicationXML) != "" {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationXMLCharsetUTF8)
		return c.Send(i.hostMeta.XML)
	}

	return c.Status(fiber.StatusNotAcceptable).SendString("Not Acceptable")
}

func (i *Handler) GetHostMetaJSON(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.Send(i.hostMeta.JSON)
}
