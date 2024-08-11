package api_schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type APIError struct {
	StatusCode  int            `json:"status_code"`
	Description string         `json:"description"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

func (e APIError) Error() string {
	if e.Metadata == nil || len(e.Metadata) == 0 {
		return fmt.Sprintf("APIError: %d - %s", e.StatusCode, e.Description)
	}

	return fmt.Sprintf("APIError: %d - %s, %s", e.StatusCode, e.Description, stringifyErrorMetadata(e.Metadata))
}

func stringifyErrorMetadata(m map[string]any) string {
	sb := strings.Builder{}
	for key, value := range m {
		sb.WriteString(fmt.Sprintf("%s=%v, ", key, value))
	}
	return strings.TrimSuffix(sb.String(), ", ")
}

func (e APIError) Equals(other any) bool {
	var err *APIError

	switch raw := other.(type) {
	case *APIError:
		err = raw
	default:
		return false
	}

	return e.StatusCode == err.StatusCode && e.Description == err.Description
}

func (e APIError) URLEncode() (string, error) {
	v := url.Values{}
	v.Set("status_code", fmt.Sprintf("%d", e.StatusCode))
	v.Set("description", e.Description)

	if e.Metadata != nil {
		b, err := json.Marshal(e.Metadata)
		if err != nil {
			return "", err
		}

		v.Set("metadata", string(b))
	}

	// Fix up spaces because Golang's net/url.URL encodes " " as "+" instead of "%20"
	// https://github.com/golang/go/issues/13982
	return strings.ReplaceAll(v.Encode(), "+", "%20"), nil
}

func NewAPIError(code int, description string) func(metadata map[string]any) *APIError {
	return func(metadata map[string]any) *APIError {
		return &APIError{StatusCode: code, Description: description, Metadata: metadata}
	}
}

type APIResponse[T any] struct {
	Ok    bool      `json:"ok"`
	Data  *T        `json:"data"`
	Error *APIError `json:"error"`
}

func NewFailedAPIResponse[T any](err error) APIResponse[T] {
	var e *APIError

	if errors.As(err, &e) {
	} else {
		e = NewAPIError(500, err.Error())(nil)
	}

	return APIResponse[T]{Ok: false, Error: e}
}
