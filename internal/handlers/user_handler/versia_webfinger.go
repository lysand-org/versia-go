package user_handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/versia-pub/versia-go/config"
	"github.com/versia-pub/versia-go/internal/api_schema"
	"github.com/versia-pub/versia-go/internal/helpers"
	"github.com/versia-pub/versia-go/pkg/webfinger"
)

func (i *Handler) Webfinger(c *fiber.Ctx) error {
	userID, err := webfinger.ParseResource(c.Query("resource"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(webfinger.Response{
			Error: helpers.StringPtr(err.Error()),
		})
	}

	if userID.Domain != config.C.PublicAddress.Host {
		return c.Status(fiber.StatusBadRequest).JSON(webfinger.Response{
			Error: helpers.StringPtr("The requested user is a remote user"),
		})
	}

	wf, err := i.userService.GetWebfingerForUser(c.UserContext(), userID.ID)
	if err != nil {
		if errors.Is(err, api_schema.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(webfinger.Response{
				Error: helpers.StringPtr("User could not be found"),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(webfinger.Response{
			Error: helpers.StringPtr("Failed to query user"),
		})
	}

	return c.JSON(wf.WebFingerResource())
}
