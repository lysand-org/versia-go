package lysand

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
)

// UserPublicKey represents a public key for a user. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/security/keys#public-key-cryptography
type UserPublicKey struct {
	Actor *URL `json:"actor"`

	// Algorithm can only be `ed25519` for now
	Algorithm string `json:"algorithm"`

	Key    *SPKIPublicKey  `json:"-"`
	RawKey json.RawMessage `json:"key"`
}

func (k *UserPublicKey) UnmarshalJSON(raw []byte) error {
	type t UserPublicKey
	k2 := (*t)(k)

	if err := json.Unmarshal(raw, k2); err != nil {
		return err
	}

	var err error
	if k2.Key, err = unmarshalSPKIPubKey(k2.Algorithm, k2.RawKey); err != nil {
		return err
	}

	*k = UserPublicKey(*k2)

	return nil
}

func (k UserPublicKey) MarshalJSON() ([]byte, error) {
	type t UserPublicKey
	k2 := t(k)

	var err error
	if k2.RawKey, err = k2.Key.MarshalJSON(); err != nil {
		return nil, err
	}

	return json.Marshal(k2)
}

// SPKIPublicKey is a type that represents a [ed25519.PublicKey] in the SPKI
// format.
type SPKIPublicKey struct {
	Key       any
	Algorithm string
}

func unmarshalSPKIPubKey(algorithm string, raw []byte) (*SPKIPublicKey, error) {
	rawStr := ""
	if err := json.Unmarshal(raw, &rawStr); err != nil {
		return nil, err
	}

	raw, err := base64.StdEncoding.DecodeString(rawStr)
	if err != nil {
		return nil, err
	}

	return NewSPKIPubKey(algorithm, raw)
}

// NewSPKIPubKey decodes the public key from a base64 encoded string and then unmarshals it from the SPKI form.
func NewSPKIPubKey(algorithm string, raw []byte) (*SPKIPublicKey, error) {
	parsed, err := x509.ParsePKIXPublicKey(raw)
	if err != nil {
		return nil, err
	}

	return &SPKIPublicKey{
		Key:       parsed,
		Algorithm: algorithm,
	}, nil
}

// MarshalJSON marshals the SPKI-encoded public key to a base64 encoded string.
func (k SPKIPublicKey) MarshalJSON() ([]byte, error) {
	raw, err := x509.MarshalPKIXPublicKey(k.Key)
	if err != nil {
		return nil, err
	}

	return json.Marshal(base64.StdEncoding.EncodeToString(raw))
}

func (k SPKIPublicKey) ToKey() crypto.PublicKey {
	return k.Key
}
