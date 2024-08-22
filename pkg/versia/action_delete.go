package versia

import (
	"encoding/json"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
)

// Delete signals the deletion of an entity. For more information, see the [Spec].
// This entity does not have a URI.
//
// Implementations must ensure that the author of the Delete entity has the authorization to delete the target entity.
//
// [Spec]: https://versia.pub/entities/delete
type Delete struct {
	Entity

	// Author is the URL to the user that triggered the deletion
	Author *versiautils.URL `json:"author"`

	// DeletedType is the type of the object that is being deleted
	DeletedType string `json:"deleted_type"`

	// Deleted is the URL to the object that is being deleted
	Deleted *versiautils.URL `json:"deleted"`
}

func (d Delete) MarshalJSON() ([]byte, error) {
	type a Delete
	d2 := a(d)
	d2.Type = "Delete"
	return json.Marshal(d2)
}
