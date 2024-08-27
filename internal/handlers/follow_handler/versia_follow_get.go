package follow_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/internal/api_schema"
)

func (i *Handler) GetVersiaFollow(c *fiber.Ctx) error {
	parsedRequestedFollowID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return api_schema.ErrBadRequest(map[string]any{"reason": "Invalid follow ID"})
	}

	f, err := i.followService.GetFollow(c.UserContext(), parsedRequestedFollowID)
	if err != nil {
		i.log.Error(err, "Failed to query follow", "id", parsedRequestedFollowID)

		return api_schema.ErrInternalServerError(map[string]any{"reason": "Failed to query follow"})
	}
	if f == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Follow not found",
		})
	}

	return c.JSON(f.ToVersia())
}
