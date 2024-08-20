package tasks

import (
	"context"
	"github.com/lysand-org/versia-go/internal/repository"

	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/entity"
)

type FederateNoteData struct {
	NoteID uuid.UUID `json:"noteID"`
}

func (t *Handler) FederateNote(ctx context.Context, data FederateNoteData) error {
	s := t.telemetry.StartSpan(ctx, "function", "tasks/Handler.FederateNote")
	defer s.End()
	ctx = s.Context()

	var n *entity.Note
	if err := t.repositories.Atomic(ctx, func(ctx context.Context, tx repository.Manager) error {
		var err error
		n, err = tx.Notes().GetByID(ctx, data.NoteID)
		if err != nil {
			return err
		}
		if n == nil {
			t.log.V(-1).Info("Could not find note", "id", data.NoteID)
			return nil
		}

		for _, uu := range n.Mentions {
			if !uu.IsRemote {
				t.log.V(2).Info("User is not remote", "user", uu.ID)
				continue
			}

			res, err := t.federationService.SendToInbox(ctx, n.Author, &uu, n.ToLysand())
			if err != nil {
				t.log.Error(err, "Failed to send note to remote user", "res", res, "user", uu.ID)
			} else {
				t.log.V(2).Info("Sent note to remote user", "res", res, "user", uu.ID)
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
