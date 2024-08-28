package task_dtos

import "github.com/google/uuid"

const (
	FederateNote = "federate_note"
)

type FederateNoteData struct {
	NoteID uuid.UUID `json:"noteID"`
}
