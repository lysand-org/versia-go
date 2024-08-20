package meta_handler

import (
	"github.com/gofiber/fiber/v2"
)

func (i *Handler) GetLysandInstanceMetadata(c *fiber.Ctx) error {
	m, err := i.instanceMetadataService.Ours(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(m.ToLysand())
}
