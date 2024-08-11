package lysand

import (
	"encoding/json"

	"github.com/Masterminds/semver"
)

// ServerMetadata represents the metadata of a Lysand server. For more information, see the [Spec].
//
// ! Unlike other objects, server metadata is not meant to be federated.
//
// [Spec]: https://lysand.org/objects/server-metadata
type ServerMetadata struct {
	// Type is always "ServerMetadata"
	// https://lysand.org/objects/server-metadata#type
	Type string `json:"type"`

	// Extensions is a map of active extensions
	// https://lysand.org/objects/server-metadata#extensions
	Extensions Extensions `json:"extensions,omitempty"`

	// Name is the name of the server
	// https://lysand.org/objects/server-metadata#name
	Name string `json:"name"`

	// Version is the version of the server software
	// https://lysand.org/objects/server-metadata#version
	Version *semver.Version `json:"version"`

	// Description is a description of the server
	// https://lysand.org/objects/server-metadata#description
	Description *string `json:"description,omitempty"`

	// Website is the URL to the server's website
	// https://lysand.org/objects/server-metadata#website
	Website *URL `json:"website,omitempty"`

	// Moderators is a list of URLs to moderators
	// https://lysand.org/objects/server-metadata#moderators
	Moderators []*URL `json:"moderators,omitempty"`

	// Admins is a list of URLs to administrators
	// https://lysand.org/objects/server-metadata#admins
	Admins []*URL `json:"admins,omitempty"`

	// Logo is the URL to the server's logo
	// https://lysand.org/objects/server-metadata#logo
	Logo *ImageContentTypeMap `json:"logo,omitempty"`

	// Banner is the URL to the server's banner
	// https://lysand.org/objects/server-metadata#banner
	Banner *ImageContentTypeMap `json:"banner,omitempty"`

	// SupportedExtensions is a list of supported extensions
	SupportedExtensions []string `json:"supported_extensions"`
}

func (s ServerMetadata) MarshalJSON() ([]byte, error) {
	type serverMetadata ServerMetadata
	s2 := serverMetadata(s)
	s2.Type = "ServerMetadata"
	return json.Marshal(s2)
}
