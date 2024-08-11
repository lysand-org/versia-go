package lysand

import "encoding/json"

type Note Publication

func (n Note) MarshalJSON() ([]byte, error) {
	type note Note
	n2 := note(n)
	n2.Type = "Note"
	return json.Marshal(n2)
}
