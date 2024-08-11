package lysand

import (
	"encoding/json"
	"fmt"
)

type inboxObject struct {
	Type string `json:"type"`
}

func ParseInboxObject(raw json.RawMessage) (any, error) {
	var i inboxObject
	if err := json.Unmarshal(raw, &i); err != nil {
		return nil, err
	}

	switch i.Type {
	case "Publication":
		m := Publication{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	case "Note":
		m := Note{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	case "Patch":
		m := Patch{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	case "Follow":
		m := Follow{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	case "FollowAccept":
		m := FollowAccept{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	case "FollowReject":
		m := FollowReject{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	case "Undo":
		m := Undo{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	default:
		return nil, ErrUnknownType{Type: i.Type}
	}
}

type ErrUnknownType struct {
	Type string
}

func (e ErrUnknownType) Error() string {
	return fmt.Sprintf("unknown inbox object type: %s", e.Type)
}
