package versiacrypto

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFederationHeaders_String(t *testing.T) {
	one := SignatureData{
		RequestMethod: "POST",
		Nonce:         "1234567890",
		URL:           &url.URL{Scheme: "https", Host: "bob.com", Path: "/users/bob", RawQuery: "z=foo&a=bar"},
		Digest:        SHA256([]byte("hello")),
	}

	assert.Equal(t, "post /users/bob?z=foo&a=bar 1234567890 LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=", one.String())
}
