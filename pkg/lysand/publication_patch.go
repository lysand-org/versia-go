package lysand

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Patch is a type that represents a modification to a note. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/patch
type Patch struct {
	Note

	// PatchedID is the ID of the publication that was patched.
	// https://lysand.org/objects/patch#patched-id
	PatchedID uuid.UUID `json:"patched_id"`

	// PatchedAt is the time that the publication was patched.
	// https://lysand.org/objects/patch#patched-at
	PatchedAt Time `json:"patched_at"`
}

func (p Patch) MarshalJSON() ([]byte, error) {
	type patch Patch
	p2 := patch(p)
	p2.Type = "Patch"
	return json.Marshal(p2)
}
