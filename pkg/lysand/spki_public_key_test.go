package lysand

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSPKIPublicKey_UnmarshalJSON(t *testing.T) {
	expectedPk := must(x509.ParsePKIXPublicKey, must(base64.StdEncoding.DecodeString, "MCowBQYDK2VwAyEAgKNt+9eyOXdb7MSrrmHlsFD2H9NGwC+56PjpWD46Tcs="))

	pk := UserPublicKey{}
	raw := []byte(`{"public_key":"MCowBQYDK2VwAyEAgKNt+9eyOXdb7MSrrmHlsFD2H9NGwC+56PjpWD46Tcs="}`)
	if err := json.Unmarshal(raw, &pk); err != nil {
		t.Error(err)
	}

	assert.Equal(t, expectedPk, ed25519.PublicKey(pk.Key))
}

func TestSPKIPublicKey_MarshalJSON(t *testing.T) {
	expectedPk := must(x509.ParsePKIXPublicKey, must(base64.StdEncoding.DecodeString, "MCowBQYDK2VwAyEAgKNt+9eyOXdb7MSrrmHlsFD2H9NGwC+56PjpWD46Tcs=")).(ed25519.PublicKey)

	pk := UserPublicKey{
		Key: SPKIPublicKey(expectedPk),
	}
	if _, err := json.Marshal(pk); err != nil {
		t.Error(err)
	}
}
