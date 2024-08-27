package versiautils

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
)

var (
	validTextContentTypes  = []string{"text/html", "text/plain"}
	validImageContentTypes = []string{"image/png", "image/jpeg", "image/gif", "image/svg+xml"}
)

// ContentMap is a map of content types to their respective content.
type ContentMap[T any] map[string]T

type UnexpectedContentTypeError struct {
	MIMEType string
}

func (e UnexpectedContentTypeError) Error() string {
	return fmt.Sprintf("unexpected content type: %s", e.MIMEType)
}

func (m ContentMap[T]) unmarshalJSON(raw []byte, valid []string) error {
	var cm map[string]json.RawMessage
	if err := json.Unmarshal(raw, &cm); err != nil {
		return err
	}

	m = make(ContentMap[T])

	errs := make([]error, 0)
	for k, v := range cm {
		if valid != nil {
			if !slices.Contains(valid, k) {
				errs = append(errs, UnexpectedContentTypeError{k})
				continue
			}
		}

		var c T
		if err := json.Unmarshal(v, &c); err != nil {
			errs = append(errs, err)
			continue
		}
		m[k] = c
	}
	if len(errs) > 0 {
		return MultipleError{errs}
	}

	return nil
}

func (m ContentMap[T]) getPreferred(preferred []string) *T {
	for _, v := range preferred {
		if c, ok := m[v]; ok {
			return &c
		}
	}

	return nil
}

// TextContent is embedded string. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/structures/content-format
type TextContent struct {
	Content string `json:"content"`

	// Remote is always false
	Remote bool `json:"remote"`
}
type TextContentTypeMap ContentMap[TextContent]

func (t TextContentTypeMap) UnmarshalJSON(data []byte) error {
	return (ContentMap[TextContent])(t).unmarshalJSON(data, validTextContentTypes)
}

func (t TextContentTypeMap) String() string {
	if c := (ContentMap[TextContent])(t).getPreferred(validTextContentTypes); c != nil {
		return c.Content
	}

	return ""
}

// File is a file or other piece of content. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/structures/content-format
type File struct {
	// Remote is always true
	Remote bool `json:"remote"`

	// URL to the attachment
	Content     *URL     `json:"content"`
	Description string   `json:"description"`
	Hash        DataHash `json:"hash"`
	Size        int      `json:"size"`

	// BlurHash is available when the content type is an image
	BlurHash *string `json:"blurhash,omitempty"`
	// BlurHash is available when the content type is an image
	Height *int `json:"height,omitempty"`
	// BlurHash is available when the content type is an image
	Width *int `json:"width,omitempty"`

	FPS *int `json:"fps,omitempty"`
}
type DataHash struct {
	SHA256 string `json:"sha256"`
}

type ImageContentMap ContentMap[File]

func (i ImageContentMap) UnmarshalJSON(data []byte) error {
	return (ContentMap[File])(i).unmarshalJSON(data, validImageContentTypes)
}

func (i ImageContentMap) String() string {
	if c := (ContentMap[File])(i).getPreferred(validImageContentTypes); c != nil {
		return c.Content.String()
	}

	return ""
}

type NoteAttachmentContentMap ContentMap[File]

var ErrContentMapEntryNotRemote = errors.New("content map entry not remote")

func (i NoteAttachmentContentMap) UnmarshalJSON(data []byte) error {
	if err := (ContentMap[File])(i).unmarshalJSON(data, nil); err != nil {
		return err
	}

	errs := make([]error, 0)
	for _, a := range i {
		if !a.Remote {
			errs = append(errs, ErrContentMapEntryNotRemote)
		}
	}
	if len(errs) > 0 {
		return MultipleError{errs}
	}

	return nil
}
