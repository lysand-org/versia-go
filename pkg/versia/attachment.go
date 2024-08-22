package versia

// Attachment is a file or other piece of content that is attached to a post. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/structures/content-format
type Attachment struct {
	// URL to the attachment
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Hash        DataHash `json:"hash"`
	Size        int      `json:"size"`

	// BlurHash is available when the content type is an image
	BlurHash *string `json:"blurhash,omitempty"`
	// BlurHash is available when the content type is an image
	Height *int `json:"height,omitempty"`
	// BlurHash is available when the content type is an image
	Width *int `json:"width,omitempty"`

	// TODO: Figure out when this is available
	FPS *int `json:"fps,omitempty"`
}

type DataHash struct {
	SHA256 string `json:"sha256"`
}
