package versia

import (
	"github.com/google/uuid"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
)

// Entity is the base type for all Lysand entities.  For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects#types
type Entity struct {
	// Type is the type of the entity
	Type string `json:"type"`

	// ID is a UUID for the entity
	ID uuid.UUID `json:"id"`

	// URI is the URL to the entity
	URI *versiautils.URL `json:"uri"`

	// CreatedAt is the time the entity was created
	CreatedAt versiautils.Time `json:"created_at"`

	// Extensions is a map of active extensions
	// https://lysand.org/objects/server-metadata#extensions
	Extensions Extensions `json:"extensions,omitempty"`
}

type Extensions map[string]any

// {
//   "org.lysand:custom_emojis": {
//     "emojis": [
//       {
//         "name": "neocat_3c",
//         "url": {
//           "image/webp": {
//             "content": "https://cdn.lysand.org/a97727158bf062ad31cbfb02e212ce0c7eca599a2f863276511b8512270b25e8/neocat_3c_256.webp"
//           }
//         }
//       }
//     ]
//   }
// }
