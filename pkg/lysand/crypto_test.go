package lysand

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFederationClient_ValidateSignatureHeader(t *testing.T) {
	var (
		bobPrivBytes = must(base64.StdEncoding.DecodeString, "MC4CAQAwBQYDK2VwBCIEINOATgmaya61Ha9OEE+DD3RnOEqDaHyQ3yLf5upwskUU")
		bobPubBytes  = must(base64.StdEncoding.DecodeString, "MCowBQYDK2VwAyEAQ08Z/FJ5f16o8mthLaFZMo4ssn0fJ7c+bipNYm3kId4=")
	)

	bobPub := must(x509.ParsePKIXPublicKey, bobPubBytes).(ed25519.PublicKey)
	bobPriv := must(x509.ParsePKCS8PrivateKey, bobPrivBytes).(ed25519.PrivateKey)

	date := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	body := []byte("hello")

	sigData := NewSignatureData("POST", date, "example2.com", "/users/bob", hashSHA256(body))

	sig := Signer{PrivateKey: bobPriv, UserURL: &url.URL{Scheme: "https", Host: "example.com", Path: "/users/bob"}}.
		Sign(*sigData)

	t.Run("validate against itself", func(t *testing.T) {
		v := Verifier{
			PublicKey: bobPub,
		}

		if !v.Verify("POST", date, "example2.com", "/users/bob", body, sig) {
			t.Error("signature verification failed")
		}
	})

	t.Run("validate against @lysand/api JS implementation", func(t *testing.T) {
		expectedSignedString := `(request-target): post /users/bob
host: example2.com
date: 1970-01-01T00:00:00.000Z
digest: SHA-256=LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=
`
		assert.Equal(t, expectedSignedString, sigData.String())

		expectedSignatureHeader := `keyId="https://example.com/users/bob",algorithm="ed25519",headers="(request-target) host date digest",signature="PbVicu1spnATYUznWn6N5ebNUC+w94U9k6y4dncLsr6hNfUD8CLInbUSkgR3AZrCWEZ+Md2+Lch70ofiSqXgAQ=="`
		assert.Equal(t, expectedSignatureHeader, sig.String())
	})
}

func TestSignatureInterop(t *testing.T) {
	var (
		bobPubBytes  = must(base64.StdEncoding.DecodeString, "MCowBQYDK2VwAyEAgKNt+9eyOXdb7MSrrmHlsFD2H9NGwC+56PjpWD46Tcs=")
		bobPrivBytes = must(base64.StdEncoding.DecodeString, "MC4CAQAwBQYDK2VwBCIEII+nkwT3nXwBp9FEE0q95RBBfikf6UTzPzdH2yrtIvL1")
	)

	bobPub := must(x509.ParsePKIXPublicKey, bobPubBytes).(ed25519.PublicKey)
	bobPriv := must(x509.ParsePKCS8PrivateKey, bobPrivBytes).(ed25519.PrivateKey)

	signedString := `(request-target): post /api/users/ec042557-8c30-492d-87d6-9e6495993072/inbox
host: lysand-test.i.devminer.xyz
date: 2024-07-25T21:03:24.866Z
digest: SHA-256=mPN5WKMoC4k3zor6FPTJUhDQ1JKX6zqA2QfEGh3omuc=
`
	method := "POST"
	dateHeader := "2024-07-25T21:03:24.866Z"
	date := must(ParseTime, dateHeader)
	host := "lysand-test.i.devminer.xyz"
	path := "/api/users/ec042557-8c30-492d-87d6-9e6495993072/inbox"
	body := []byte(`{"type":"Follow","id":"2265b3b2-a176-4b20-8fcf-ac82cf2efd7d","author":"https://lysand.i.devminer.xyz/users/0190d697-c83a-7376-8d15-0f77fd09e180","followee":"https://lysand-test.i.devminer.xyz/api/users/ec042557-8c30-492d-87d6-9e6495993072/","created_at":"2024-07-25T21:03:24.863Z","uri":"https://lysand.i.devminer.xyz/follows/2265b3b2-a176-4b20-8fcf-ac82cf2efd7d"}`)
	signatureHeader := `keyId="https://lysand.i.devminer.xyz/users/0190d697-c83a-7376-8d15-0f77fd09e180",algorithm="ed25519",headers="(request-target) host date digest",signature="KUkKYexLk2hOfE+NVIacLDHSJP2QpX4xJGclHhQIM39ce2or7UJauRtCL8eWrhpSgQdVPk11bYhvvi8fdCruBw=="`

	sigData := NewSignatureData(method, date.ToStd(), host, path, hashSHA256(body))
	assert.Equal(t, signedString, sigData.String())

	t.Run("signature header parsing", func(t *testing.T) {
		parsedSignatureHeader, err := ParseSignatureHeader(signatureHeader)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, "https://lysand.i.devminer.xyz/users/0190d697-c83a-7376-8d15-0f77fd09e180", parsedSignatureHeader.KeyID.String())
		assert.Equal(t, "ed25519", parsedSignatureHeader.Algorithm)
		assert.Equal(t, "(request-target) host date digest", parsedSignatureHeader.Headers)
		assert.Equal(t, sigData.Sign(bobPriv), parsedSignatureHeader.Signature)

		v := Verifier{PublicKey: bobPub}
		if !v.Verify(method, date.ToStd(), host, path, body, parsedSignatureHeader) {
			t.Error("signature verification failed")
		}
	})

	t.Run("signature header generation", func(t *testing.T) {
		sig := Signer{PrivateKey: bobPriv, UserURL: &url.URL{Scheme: "https", Host: "lysand.i.devminer.xyz", Path: "/users/0190d697-c83a-7376-8d15-0f77fd09e180"}}.
			Sign(*sigData)
		assert.Equal(t, signatureHeader, sig.String())
	})
}
