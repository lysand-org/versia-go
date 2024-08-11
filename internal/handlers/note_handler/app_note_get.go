package note_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/api_schema"
)

func (i *Handler) GetNote(c *fiber.Ctx) error {
	parsedRequestedNoteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return api_schema.ErrBadRequest(map[string]any{
			"reason": "Invalid note ID",
		})
	}

	u, err := i.noteService.GetNote(c.UserContext(), parsedRequestedNoteID)
	if err != nil {
		i.log.Error(err, "Failed to query note", "id", parsedRequestedNoteID)

		return api_schema.ErrInternalServerError(map[string]any{"reason": "Failed to query note"})
	}
	if u == nil {
		return api_schema.ErrNotFound(nil)
	}

	return c.JSON(u.ToLysand())
}
