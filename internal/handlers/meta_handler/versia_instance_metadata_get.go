package meta_handler

import (
	"github.com/gofiber/fiber/v2"
)

func (i *Handler) GetVersiaInstanceMetadata(c *fiber.Ctx) error {
	m, err := i.instanceMetadataService.Ours(c.UserContext())
	if err != nil {
		return err
	}

	// TODO: Sign with the instance private key
	return c.JSON(m.ToVersia())
}
