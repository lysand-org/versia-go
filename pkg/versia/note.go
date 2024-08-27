package versia

import (
	"encoding/json"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

// Note is a published message, similar to a tweet (from Twitter) or a toot (from Mastodon).
// For more information, see the [Spec].
//
// [Spec]: https://versia.pub/entities/note
type Note struct {
	Entity

	// Attachments is a list of attachment objects, keyed by their MIME type
	Attachments []versiautils.NoteAttachmentContentMap `json:"attachments,omitempty"`

	// Author is the URL to the user
	Author *versiautils.URL `json:"author"`

	// Category is the category of the note
	Category *CategoryType `json:"category,omitempty"`

	// Content is the content of the note
	Content versiautils.TextContentTypeMap `json:"content,omitempty"`

	// Device that created the note
	Device *Device `json:"device,omitempty"`

	// Group is the URL to a group
	// TODO: Properly parse these, can be "public" | "followers" as well
	Group *versiautils.URL `json:"group,omitempty"`

	// IsSensitive is a boolean indicating whether the note contains sensitive content
	IsSensitive *bool `json:"is_sensitive,omitempty"`

	// Mentions is a list of URLs to users
	Mentions []versiautils.URL `json:"mentions,omitempty"`

	// Previews is a list of URLs to preview images
	Previews []LinkPreview `json:"previews,omitempty"`

	// Quotes is the URL to the note being quoted
	Quotes *versiautils.URL `json:"quotes,omitempty"`

	// RepliesTo is the URL to the note being replied to
	RepliesTo *versiautils.URL `json:"replies_to,omitempty"`

	// Subject is the subject of the note
	Subject *string `json:"subject,omitempty"`
}

func (p Note) MarshalJSON() ([]byte, error) {
	type a Note
	n2 := a(p)
	n2.Type = "Note"
	return json.Marshal(n2)
}

// LinkPreview is a preview of a link. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/entities/note#entity-definition
type LinkPreview struct {
	Link        *versiautils.URL `json:"link"`
	Title       string           `json:"title"`
	Description *string          `json:"description,omitempty"`
	Image       *versiautils.URL `json:"image,omitempty"`
	Icon        *versiautils.URL `json:"icon,omitempty"`
}

// Device is the device that creates note. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/entities/note#entity-definition
type Device struct {
	Name    string           `json:"name"`
	Version string           `json:"version,omitempty"`
	URL     *versiautils.URL `json:"url,omitempty"`
}

// CategoryType is the type of note. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/entities/note#entity-definition
type CategoryType string

const (
	// CategoryMicroblog is similar to Twitter, Mastodon
	CategoryMicroblog CategoryType = "microblog"
	// CategoryForum is similar to Reddit
	CategoryForum CategoryType = "forum"
	// CategoryBlog is similar to Wordpress, WriteFreely
	CategoryBlog CategoryType = "blog"
	// CategoryImage is similar to Instagram
	CategoryImage CategoryType = "image"
	// CategoryVideo is similar to YouTube
	CategoryVideo CategoryType = "video"
	// CategoryAudio is similar to SoundCloud, Spotify
	CategoryAudio CategoryType = "audio"
	// CategoryMessaging is similar to Discord, Matrix, Signal
	CategoryMessaging CategoryType = "messaging"
)
