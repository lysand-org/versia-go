package lysand

import (
	"encoding/json"
	"slices"

	"github.com/rs/zerolog/log"
)

var (
	validTextContentTypes  = []string{"text/html", "text/plain"}
	validImageContentTypes = []string{"image/png", "image/jpeg", "image/gif", "image/svg+xml"}
)

// ContentTypeMap is a map of content types to their respective content.
type ContentTypeMap[T any] map[string]T

func (m *ContentTypeMap[T]) unmarshalJSON(raw []byte, valid []string) error {
	var cm map[string]json.RawMessage
	if err := json.Unmarshal(raw, &cm); err != nil {
		return err
	}

	*m = make(ContentTypeMap[T])

	for k, v := range cm {
		if !slices.Contains(valid, k) {
			// TODO: replace with logr
			log.Debug().Caller().Str("mimetype", k).Msg("unexpected content type, skipping")
			continue
		}

		var c T
		if err := json.Unmarshal(v, &c); err != nil {
			return err
		}
		(*m)[k] = c
	}

	return nil
}

func (m ContentTypeMap[T]) getPreferred(preferred []string) *T {
	for _, v := range preferred {
		if c, ok := m[v]; ok {
			return &c
		}
	}

	return nil
}

type TextContent struct {
	Content string `json:"content"`
}
type TextContentTypeMap ContentTypeMap[TextContent]

func (t *TextContentTypeMap) UnmarshalJSON(data []byte) error {
	return (*ContentTypeMap[TextContent])(t).unmarshalJSON(data, validTextContentTypes)
}

func (t TextContentTypeMap) String() string {
	if c := (ContentTypeMap[TextContent])(t).getPreferred(validTextContentTypes); c != nil {
		return c.Content
	}

	return ""
}

type ImageContent struct {
	Content *URL `json:"content"`
}
type ImageContentTypeMap ContentTypeMap[ImageContent]

func (i *ImageContentTypeMap) UnmarshalJSON(data []byte) error {
	return (*ContentTypeMap[ImageContent])(i).unmarshalJSON(data, validImageContentTypes)
}

func (i ImageContentTypeMap) String() string {
	if c := (ContentTypeMap[ImageContent])(i).getPreferred(validImageContentTypes); c != nil {
		return c.Content.String()
	}

	return ""
}
