package meta_handler

import (
	"github.com/Masterminds/semver"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

func (i *Handler) GetLysandServerMetadata(c *fiber.Ctx) error {
	return c.JSON(lysand.ServerMetadata{
		// TODO: Get version from build linker flags
		Version: semver.MustParse("0.0.0-dev"),

		Name:        config.C.InstanceName,
		Description: config.C.InstanceDescription,
		Website:     lysand.URLFromStd(config.C.PublicAddress),

		// TODO: Get more info
		Moderators: nil,
		Admins:     nil,
		Logo:       nil,
		Banner:     nil,

		SupportedExtensions: []string{},
		Extensions:          map[string]any{},
	})
}
