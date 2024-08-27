package note_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/internal/api_schema"
)

func (i *Handler) GetNote(c *fiber.Ctx) error {
	parsedRequestedNoteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return api_schema.ErrBadRequest(map[string]any{
			"reason": "Invalid note ID",
		})
	}

	n, err := i.noteService.GetNote(c.UserContext(), parsedRequestedNoteID)
	if err != nil {
		i.log.Error(err, "Failed to query note", "id", parsedRequestedNoteID)

		return api_schema.ErrInternalServerError(map[string]any{"reason": "Failed to query note"})
	}
	if n == nil {
		return api_schema.ErrNotFound(nil)
	}

	if !n.Author.IsRemote {
		// For local authors we can also sign the note
		if err := i.requestSigner.SignAndSend(c, n.Author.Signer, n.ToVersia()); err != nil {
			i.log.Error(err, "Failed to sign response body", "id", parsedRequestedNoteID)

			return api_schema.ErrInternalServerError(map[string]any{
				"reason": "failed to sign response body",
			})
		}
	}

	return c.JSON(n.ToVersia())
}
