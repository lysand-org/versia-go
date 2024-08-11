package lysand

// PublicationVisibility is the visibility of a publication. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/publications#visibility
type PublicationVisibility string

const (
	// PublicationVisiblePublic means that the publication is visible to everyone.
	PublicationVisiblePublic PublicationVisibility = "public"
	// PublicationVisibleUnlisted means that the publication is visible everyone, but should not appear in public timelines and search results.
	PublicationVisibleUnlisted PublicationVisibility = "unlisted"
	// PublicationVisibleFollowers means that the publication is visible to followers only.
	PublicationVisibleFollowers PublicationVisibility = "followers"
	// PublicationVisibleDirect means that the publication is a direct message, and is visible only to the mentioned users.
	PublicationVisibleDirect PublicationVisibility = "direct"
)

// Publication is a publication object. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/publications
type Publication struct {
	Entity

	// Author is the URL to the user
	// https://lysand.org/objects/publications#author
	Author *URL `json:"author"`

	// Content is the content of the publication
	// https://lysand.org/objects/publications#content
	Content TextContentTypeMap `json:"content,omitempty"`

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
	Group *URL `json:"group,omitempty"`

	// Attachments is a list of attachment objects, keyed by their MIME type
	// https://lysand.org/objects/publications#attachments
	Attachments []ContentTypeMap[Attachment] `json:"attachments,omitempty"`

	// RepliesTo is the URL to the publication being replied to
	// https://lysand.org/objects/publications#replies-to
	RepliesTo *URL `json:"replies_to,omitempty"`

	// Quoting is the URL to the publication being quoted
	// https://lysand.org/objects/publications#quotes
	Quoting *URL `json:"quoting,omitempty"`

	// Mentions is a list of URLs to users
	// https://lysand.org/objects/publications#mentionshttps://lysand.org/objects/publications#mentions
	Mentions []URL `json:"mentions,omitempty"`

	// Subject is the subject of the publication
	// https://lysand.org/objects/publications#subject
	Subject *string `json:"subject,omitempty"`

	// IsSensitive is a boolean indicating whether the publication contains sensitive content
	// https://lysand.org/objects/publications#is-sensitive
	IsSensitive *bool `json:"is_sensitive,omitempty"`

	// Visibility is the visibility of the publication
	// https://lysand.org/objects/publications#visibility
	Visibility PublicationVisibility `json:"visibility"`
}

// LinkPreview is a preview of a link. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/publications#types
type LinkPreview struct {
	Link        URL     `json:"link"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Image       *URL    `json:"image"`
	Icon        *URL    `json:"icon"`
}

// Device is the device that creates publications. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/publications#types
type Device struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	URL     *URL   `json:"url,omitempty"`
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
