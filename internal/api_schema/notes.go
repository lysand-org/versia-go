package api_schema

import (
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/pkg/versia"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
)

type Note struct {
	ID uuid.UUID `json:"id,string"`
}

type FetchNoteResponse = APIResponse[Note]

type CreateNoteRequest struct {
	Content    string                `json:"content" validate:"required,min=1,max=1024"`
	Visibility versia.NoteVisibility `json:"visibility" validate:"required,oneof=public unlisted private direct"`
	Mentions   []versiautils.URL     `json:"mentions"`
}
