package api_schema

import (
	"github.com/google/uuid"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

type Note struct {
	ID uuid.UUID `json:"id,string"`
}

type FetchNoteResponse = APIResponse[Note]

type CreateNoteRequest struct {
	Content  string            `json:"content" validate:"required,min=1,max=1024"`
	Group    string            `json:"group" validate:"required"`
	Mentions []versiautils.URL `json:"mentions"`
}
