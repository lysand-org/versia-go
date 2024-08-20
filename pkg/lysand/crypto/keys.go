package versiacrypto

import (
	"crypto/ed25519"
	"fmt"
)

type UnknownPublicKeyTypeError struct {
	Got string
}

func (i UnknownPublicKeyTypeError) Error() string {
	return fmt.Sprintf("unknown public key type: \"%s\"", i.Got)
}

func ToTypedKey(algorithm string, raw []byte) (any, error) {
	switch algorithm {
	case "ed25519":
		return ed25519.PublicKey(raw), nil
	default:
		return nil, UnknownPublicKeyTypeError{algorithm}
	}
}
