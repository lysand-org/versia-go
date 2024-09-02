package versiacrypto

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

func TestFederationClient_ValidateSignatureHeader(t *testing.T) {
	var (
		bobURL = &url.URL{Scheme: "https", Host: "bob.com"}

		bobPrivBytes = versiautils.Must(base64.StdEncoding.DecodeString, "MC4CAQAwBQYDK2VwBCIEINOATgmaya61Ha9OEE+DD3RnOEqDaHyQ3yLf5upwskUU")
		bobPriv      = versiautils.Must(x509.ParsePKCS8PrivateKey, bobPrivBytes).(ed25519.PrivateKey)
		signer       = Signer{PrivateKey: bobPriv, UserURL: bobURL}

		bobPubBytes = versiautils.Must(base64.StdEncoding.DecodeString, "MCowBQYDK2VwAyEAQ08Z/FJ5f16o8mthLaFZMo4ssn0fJ7c+bipNYm3kId4=")
		bobPub      = versiautils.Must(x509.ParsePKIXPublicKey, bobPubBytes).(ed25519.PublicKey)
		verifier    = Verifier{PublicKey: bobPub}

		method = "POST"
		nonce  = "myrandomnonce"
		u      = &url.URL{Scheme: "https", Host: "bob.com", Path: "/a/b/c", RawQuery: "z=foo&a=bar"}
		body   = []byte("hello")
	)

	toSign := NewSignatureData(method, nonce, u, SHA256(body))
	assert.Equal(t, `post /a/b/c?z=foo&a=bar myrandomnonce LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=`, toSign.String())

	signed := signer.Sign(*toSign)
	assert.Equal(t, true, verifier.Verify(method, u, body, signed), "signature verification failed")

	assert.Equal(t, "myrandomnonce", signed.Nonce)
	assert.Equal(t, bobURL, signed.SignedBy)
	assert.Equal(t, "datQHNaqJ1jeKzK3UeReUVf+B65JPq5P9LxfqUUJTMv3QNqDu5KawosKoduIRk4/D/A+EKjDhlcw0c7GzUlMCA==", base64.StdEncoding.EncodeToString(signed.Signature))
}
