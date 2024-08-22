package versia

import (
	"encoding/json"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
)

// NoteVisibility is the visibility of a note. For more information, see the [Spec].
//
// TODO:
// [Spec]: https://lysand.org/objects/publications#visibility
type NoteVisibility string

const (
	// NoteVisiblePublic means that the Note is visible to everyone.
	NoteVisiblePublic NoteVisibility = "public"
	// NoteVisibleUnlisted means that the Note is visible everyone, but should not appear in public timelines and search results.
	NoteVisibleUnlisted NoteVisibility = "unlisted"
	// NoteVisibleFollowers means that the Note is visible to followers only.
	NoteVisibleFollowers NoteVisibility = "followers"
	// NoteVisibleDirect means that the Note is a direct message, and is visible only to the mentioned users.
	NoteVisibleDirect NoteVisibility = "direct"
)

// Note is a published message, similar to a tweet (from Twitter) or a toot (from Mastodon).
// For more information, see the [Spec].
//
// [Spec]: https://versia.pub/entities/note
type Note struct {
	Entity

	// Author is the URL to the user
	// https://lysand.org/objects/publications#author
	Author *versiautils.URL `json:"author"`

	// Content is the content of the publication
	// https://lysand.org/objects/publications#content
	Content versiautils.TextContentTypeMap `json:"content,omitempty"`

	// Category is the category of the publication
	// https://lysand.org/objects/publications#category
	Category *CategoryType `json:"category,omitempty"`

	// Device that created the publication
	// https://lysand.org/objects/publications#device
	Device *Device `json:"device,omitempty"`

	// Previews is a list of URLs to preview images
	// https://lysand.org/objects/publications#previews
	Previews []LinkPreview `json:"previews,omitempty"`

	// Group is the URL to a group
	// https://lysand.org/objects/publications#group
	Group *versiautils.URL `json:"group,omitempty"`

	// Attachments is a list of attachment objects, keyed by their MIME type
	// https://lysand.org/objects/publications#attachments
	Attachments []versiautils.ContentTypeMap[Attachment] `json:"attachments,omitempty"`

	// RepliesTo is the URL to the publication being replied to
	// https://lysand.org/objects/publications#replies-to
	RepliesTo *versiautils.URL `json:"replies_to,omitempty"`

	// Quoting is the URL to the publication being quoted
	// https://lysand.org/objects/publications#quotes
	Quoting *versiautils.URL `json:"quoting,omitempty"`

	// Mentions is a list of URLs to users
	// https://lysand.org/objects/publications#mentionshttps://lysand.org/objects/publications#mentions
	Mentions []versiautils.URL `json:"mentions,omitempty"`

	// Subject is the subject of the publication
	// https://lysand.org/objects/publications#subject
	Subject *string `json:"subject,omitempty"`

	// IsSensitive is a boolean indicating whether the publication contains sensitive content
	// https://lysand.org/objects/publications#is-sensitive
	IsSensitive *bool `json:"is_sensitive,omitempty"`

	// Visibility is the visibility of the publication
	// https://lysand.org/objects/publications#visibility
	Visibility NoteVisibility `json:"visibility"`
}

func (p Note) MarshalJSON() ([]byte, error) {
	type a Note
	n2 := a(p)
	n2.Type = "Note"
	return json.Marshal(n2)
}

// LinkPreview is a preview of a link. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/publications#types
type LinkPreview struct {
	Link        *versiautils.URL `json:"link"`
	Title       string           `json:"title"`
	Description *string          `json:"description"`
	Image       *versiautils.URL `json:"image"`
	Icon        *versiautils.URL `json:"icon"`
}

// Device is the device that creates publications. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/publications#types
type Device struct {
	Name    string           `json:"name"`
	Version string           `json:"version,omitempty"`
	URL     *versiautils.URL `json:"url,omitempty"`
}

// CategoryType is the type of publication. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/publications#types
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
