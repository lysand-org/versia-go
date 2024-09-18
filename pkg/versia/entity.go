package versia

import (
	"github.com/google/uuid"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

// Entity is the base type for all Versia entities.  For more information, see the [Spec].
//
// [Spec]: https://versia.pub/entities
type Entity struct {
	// ID is a UUID for the entity
	ID uuid.UUID `json:"id"`

	// Type is the type of the entity
	Type string `json:"type"`

	// CreatedAt is the time the entity was created
	CreatedAt versiautils.Time `json:"created_at"`

	// URI is the URL to the entity
	URI *versiautils.URL `json:"uri"`

	// Extensions is a map of active extensions for the entity
	Extensions Extensions `json:"extensions,omitempty"`
}
