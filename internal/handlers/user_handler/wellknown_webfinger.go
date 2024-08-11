package user_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/internal/helpers"
	"github.com/lysand-org/versia-go/pkg/webfinger"
)

func (i *Handler) Webfinger(c *fiber.Ctx) error {
	userID, err := webfinger.ParseResource(c.Query("resource"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(webfinger.Response{
			Error: helpers.StringPtr(err.Error()),
		})
	}

	if userID.Domain != config.C.PublicAddress.Host {
		return c.Status(fiber.StatusNotFound).JSON(webfinger.Response{
			Error: helpers.StringPtr("The requested user is a remote user"),
		})
	}

	wf, err := i.userService.GetWebfingerForUser(c.UserContext(), userID.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(webfinger.Response{
			Error: helpers.StringPtr("Failed to query user"),
		})
	}

	return c.JSON(wf.WebFingerResource())
}
