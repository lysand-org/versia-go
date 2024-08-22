package versia

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
	case "Note":
		m := Note{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	case "Group":
		m := Group{}
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
	case "Unfollow":
		m := Unfollow{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		return m, nil
	default:
		return nil, UnknownEntityTypeError{Type: i.Type}
	}
}

type UnknownEntityTypeError struct {
	Type string
}

func (e UnknownEntityTypeError) Error() string {
	return fmt.Sprintf("unknown entity type: %s", e.Type)
}
