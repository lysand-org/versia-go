package versiacrypto

import (
	"crypto"
	"crypto/ed25519"
	"fmt"
	"reflect"
)

type InvalidPublicKeyTypeError struct {
	Got reflect.Type
}

func (i InvalidPublicKeyTypeError) Error() string {
	return fmt.Sprintf("failed to convert public key of type \"%s\"", i.Got.String())
}

type Verify = func(data, signature []byte) bool

func NewVerify(pubKey crypto.PublicKey) (Verify, error) {
	switch pk := pubKey.(type) {
	case ed25519.PublicKey:
		return func(data, signature []byte) bool {
			return ed25519.Verify(pk, data, signature)
		}, nil
	default:
		return nil, InvalidPublicKeyTypeError{reflect.TypeOf(pk)}
	}
}
