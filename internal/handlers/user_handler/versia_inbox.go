package user_handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lysand-org/versia-go/internal/validators/val_impls"
	"github.com/lysand-org/versia-go/pkg/versia"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/api_schema"
)

func (i *Handler) HandleVersiaInbox(c *fiber.Ctx) error {
	if err := i.requestValidator.ValidateFiberCtx(c.UserContext(), c); err != nil {
		if errors.Is(err, val_impls.ErrInvalidSignature) {
			i.log.Error(err, "Invalid signature")
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		i.log.Error(err, "Failed to validate signature")
		return err
	}

	var raw json.RawMessage
	if err := json.Unmarshal(c.Body(), &raw); err != nil {
		i.log.Error(err, "Failed to unmarshal inbox object")
		return api_schema.ErrBadRequest(nil)
	}

	obj, err := versia.ParseInboxObject(raw)
	if err != nil {
		i.log.Error(err, "Failed to parse inbox object")

		if errors.Is(err, versia.UnknownEntityTypeError{}) {
			return api_schema.ErrNotFound(map[string]any{
				"error": fmt.Sprintf("Unknown object type: %s", err.(versia.UnknownEntityTypeError).Type),
			})
		}

		return err
	}

	userId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	if err := i.inboxService.Handle(c.UserContext(), obj, userId); err != nil {
		i.log.Error(err, "Failed to handle inbox", "user", userId)
	}

	return c.SendStatus(fiber.StatusOK)
}
