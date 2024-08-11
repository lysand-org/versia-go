package note_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/internal/api_schema"
)

func (i *Handler) CreateNote(c *fiber.Ctx) error {
	req := api_schema.CreateNoteRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	if err := i.bodyValidator.Validate(req); err != nil {
		return err
	}

	n, err := i.noteService.CreateNote(c.UserContext(), req)
	if err != nil {
		return err
	}
	if n == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create note",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(api_schema.Note{
		ID: n.ID,
	})
}
