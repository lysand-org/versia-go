package versia

import (
	"encoding/json"
	versiacrypto "github.com/lysand-org/versia-go/pkg/versia/crypto"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
)

// InstanceMetadata represents the metadata of a Lysand instance. For more information, see the [Spec].
//
// ! Unlike other entities, instance metadata is not meant to be federated.
//
// [Spec]: https://versia.pub/entities/instance-metadata
type InstanceMetadata struct {
	// Type is always "InstanceMetadata"
	Type string `json:"type"`

	// CreatedAt is the initial date when the instance was first created
	CreatedAt versiautils.Time `json:"created_at"`

	// Extensions is a map of active extensions
	Extensions Extensions `json:"extensions,omitempty"`

	// Name is the name of the instance
	Name string `json:"name"`

	// Description is a description of the instance
	Description *string `json:"description,omitempty"`

	// Host is the hostname of the instance, including the port
	Host string `json:"host"`

	// PublicKey is the public key of the instance
	PublicKey InstancePublicKey `json:"public_key"`

	// SharedInbox is the URL to the instance's shared inbox
	SharedInbox *versiautils.URL `json:"shared_inbox,omitempty"`

	// Moderators is a URL to a collection of moderators
	Moderators *versiautils.URL `json:"moderators,omitempty"`

	// Admins is a URL to a collection of administrators
	Admins *versiautils.URL `json:"admins,omitempty"`

	// Logo is the URL to the instance's logo
	Logo *versiautils.ImageContentTypeMap `json:"logo,omitempty"`

	// Banner is the URL to the instance's banner
	Banner *versiautils.ImageContentTypeMap `json:"banner,omitempty"`

	// Software is information about the instance software
	Software InstanceSoftware `json:"software"`

	// Compatibility is information about the instance's compatibility with different Lysand versions
	Compatibility InstanceCompatibility `json:"compatibility"`
}

func (s InstanceMetadata) MarshalJSON() ([]byte, error) {
	type instanceMetadata InstanceMetadata
	s2 := instanceMetadata(s)
	s2.Type = "InstanceMetadata"
	return json.Marshal(s2)
}

// InstanceSoftware represents the software of a Lysand instance.
type InstanceSoftware struct {
	// Name is the name of the instance software
	Name string `json:"name"`
	// Version is the version of the instance software
	Version string `json:"version"`
}

// InstanceCompatibility represents the compatibility of a Lysand instance.
type InstanceCompatibility struct {
	// Versions is a list of versions of Lysand the instance is compatible with
	Versions []string `json:"versions"`

	// Extensions is a list of extensions supported by the instance
	Extensions []string `json:"extensions"`
}

// InstancePublicKey represents the public key of a Versia instance.
type InstancePublicKey struct {
	// Algorithm can only be `ed25519` for now
	Algorithm string `json:"algorithm"`

	Key    *versiacrypto.SPKIPublicKey `json:"-"`
	RawKey json.RawMessage             `json:"key"`
}

func (k *InstancePublicKey) UnmarshalJSON(raw []byte) error {
	type t InstancePublicKey
	k2 := (*t)(k)
	if err := json.Unmarshal(raw, k2); err != nil {
		return nil
	}

	var err error
	if k2.Key, err = versiacrypto.UnmarshalSPKIPubKey(k2.Algorithm, k2.RawKey); err != nil {
		return nil
	}

	return nil
}

func (k InstancePublicKey) MarshalJSON() ([]byte, error) {
	type t InstancePublicKey
	k2 := t(k)

	var err error
	if k2.RawKey, err = k2.Key.MarshalJSON(); err != nil {
		return nil, err
	}

	return json.Marshal(k2)
}
