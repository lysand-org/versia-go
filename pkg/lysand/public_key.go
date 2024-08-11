package lysand

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
)

var (
	ErrInvalidPublicKeyType = errors.New("invalid public key type")
)

// PublicKey represents a public key for a user. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/security/keys#public-key-cryptography
type PublicKey struct {
	PublicKey SPKIPublicKey `json:"public_key"`
	Actor     *URL          `json:"actor"`
}

// SPKIPublicKey is a type that represents a [ed25519.PublicKey] in the SPKI
// format.
type SPKIPublicKey ed25519.PublicKey

// UnmarshalJSON decodes the public key from a base64 encoded string and then unmarshals it from the SPKI form.
func (k *SPKIPublicKey) UnmarshalJSON(raw []byte) error {
	rawStr := ""
	if err := json.Unmarshal(raw, &rawStr); err != nil {
		return err
	}

	raw, err := base64.StdEncoding.DecodeString(rawStr)
	if err != nil {
		return err
	}

	parsed, err := x509.ParsePKIXPublicKey(raw)
	if err != nil {
		return err
	}

	edKey, ok := parsed.(ed25519.PublicKey)
	if !ok {
		return ErrInvalidPublicKeyType
	}

	*k = SPKIPublicKey(edKey)

	return nil
}

// MarshalJSON marshals the SPKI-encoded public key to a base64 encoded string.
func (k SPKIPublicKey) MarshalJSON() ([]byte, error) {
	raw, err := x509.MarshalPKIXPublicKey(ed25519.PublicKey(k))
	if err != nil {
		return nil, err
	}

	return json.Marshal(base64.StdEncoding.EncodeToString(raw))
}

func (k SPKIPublicKey) ToStd() ed25519.PublicKey {
	return ed25519.PublicKey(k)
}
