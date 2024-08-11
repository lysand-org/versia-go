package user_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/api_schema"
)

func (i *Handler) GetLysandUser(c *fiber.Ctx) error {
	parsedRequestedUserID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	u, err := i.userService.GetUserByID(c.UserContext(), parsedRequestedUserID)
	if err != nil {
		i.log.Error(err, "Failed to query user", "id", parsedRequestedUserID)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query user",
			"id":    parsedRequestedUserID,
		})
	}
	if u == nil {
		return api_schema.ErrNotFound(map[string]any{"id": parsedRequestedUserID})
	}

	return c.JSON(u.ToLysand())
}
