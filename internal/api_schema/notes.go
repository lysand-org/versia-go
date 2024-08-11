package api_schema

import (
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

type Note struct {
	ID uuid.UUID `json:"id,string"`
}

type FetchNoteResponse = APIResponse[Note]

type CreateNoteRequest struct {
	Content    string                       `json:"content" validate:"required,min=1,max=1024"`
	Visibility lysand.PublicationVisibility `json:"visibility" validate:"required,oneof=public private direct"`
	Mentions   []lysand.URL                 `json:"mentions"`
}
