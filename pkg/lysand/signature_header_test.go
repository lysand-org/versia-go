package lysand

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseSignatureHeader(t *testing.T) {
	data := `keyId="https://example.com/users/bob",algorithm="ed25519",headers="(request-target) host date digest",signature="PbVicu1spnATYUznWn6N5ebNUC+w94U9k6y4dncLsr6hNfUD8CLInbUSkgR3AZrCWEZ+Md2+Lch70ofiSqXgAQ=="`
	expectedSignature := must(base64.StdEncoding.DecodeString, "PbVicu1spnATYUznWn6N5ebNUC+w94U9k6y4dncLsr6hNfUD8CLInbUSkgR3AZrCWEZ+Md2+Lch70ofiSqXgAQ==")

	sig, err := ParseSignatureHeader(data)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "https://example.com/users/bob", sig.KeyID.String())
	assert.Equal(t, "ed25519", sig.Algorithm)
	assert.Equal(t, "(request-target) host date digest", sig.Headers)
	assert.Equal(t, expectedSignature, sig.Signature)
}

func TestSignatureHeader_String(t *testing.T) {
	one := SignatureData{
		RequestMethod: "POST",
		Date:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		Host:          "example2.com",
		Path:          "/users/bob",
		Digest:        hashSHA256([]byte("hello")),
	}

	expected := `(request-target): post /users/bob
host: example2.com
date: 1970-01-01T00:00:00.000Z
digest: SHA-256=LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=
`

	assert.Equal(t, expected, one.String())
}
