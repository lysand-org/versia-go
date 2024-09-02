package versia

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	versiacrypto "github.com/versia-pub/versia-go/pkg/versia/crypto"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

func TestUserPublicKey_UnmarshalJSON(t *testing.T) {
	expectedPk := versiautils.Must(x509.ParsePKIXPublicKey, versiautils.Must(base64.StdEncoding.DecodeString, "MCowBQYDK2VwAyEAgKNt+9eyOXdb7MSrrmHlsFD2H9NGwC+56PjpWD46Tcs="))

	pk := UserPublicKey{}
	raw := []byte(`{"algorithm":"ed25519","key":"MCowBQYDK2VwAyEAgKNt+9eyOXdb7MSrrmHlsFD2H9NGwC+56PjpWD46Tcs="}`)
	if err := json.Unmarshal(raw, &pk); err != nil {
		t.Error(err)
	}

	assert.Equal(t, "ed25519", pk.Algorithm)
	assert.Equal(t, "ed25519", pk.Key.Algorithm)
	assert.Equal(t, expectedPk, pk.Key.Key.(ed25519.PublicKey))
}

func TestUserPublicKey_MarshalJSON(t *testing.T) {
	expectedPk := versiautils.Must(x509.ParsePKIXPublicKey, versiautils.Must(base64.StdEncoding.DecodeString, "MCowBQYDK2VwAyEAgKNt+9eyOXdb7MSrrmHlsFD2H9NGwC+56PjpWD46Tcs=")).(ed25519.PublicKey)

	pk := UserPublicKey{
		Key: &versiacrypto.SPKIPublicKey{Key: expectedPk, Algorithm: "ed25519"},
	}
	if _, err := json.Marshal(pk); err != nil {
		t.Error(err)
	}
}
