package lysand

import "encoding/json"

type Undo struct {
	Entity

	// Author is the URL to the user that triggered the undo action
	Author *URL `json:"author"`
	// Object is the URL to the object that was undone
	Object *URL `json:"object"`
}

func (u Undo) MarshalJSON() ([]byte, error) {
	type undo Undo
	u2 := undo(u)
	u2.Type = "Undo"
	return json.Marshal(u2)
}
