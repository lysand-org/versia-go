package versia

import (
	"encoding/json"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

// Group is a way to organize users and notes into communities. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/entities/pub
type Group struct {
	Entity

	// Name is the group's name / title.
	Name versiautils.TextContentTypeMap `json:"name"`

	// Description is a description of the group's contents / purpose.
	Description versiautils.TextContentTypeMap `json:"description"`

	// Members is a list of URLs of the group's members.
	Members []versiautils.URL `json:"members"`

	// Notes is a URL to the collection of notes associated with this group.
	Notes *versiautils.URL `json:"notes"`
}

func (g Group) MarshalJSON() ([]byte, error) {
	type a Group
	g2 := a(g)
	g2.Type = "Group"
	return json.Marshal(g2)
}
