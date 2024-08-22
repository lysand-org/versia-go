package versiacrypto

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
)

// SPKIPublicKey is a type that represents a [ed25519.PublicKey] in the SPKI
// format.
type SPKIPublicKey struct {
	Key       any
	Algorithm string
}

func UnmarshalSPKIPubKey(algorithm string, raw []byte) (*SPKIPublicKey, error) {
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
