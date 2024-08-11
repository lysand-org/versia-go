package lysand

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var (
	ErrInvalidSignatureHeader = errors.New("invalid signature header")
)

type SignatureHeader struct {
	// URL to a user
	KeyID     *url.URL
	Headers   string
	Algorithm string
	Signature []byte
}

func (s SignatureHeader) String() string {
	return strings.Join([]string{
		fmt.Sprintf(`keyId="%s"`, s.KeyID.String()),
		fmt.Sprintf(`algorithm="%s"`, s.Algorithm),
		fmt.Sprintf(`headers="%s"`, s.Headers),
		fmt.Sprintf(`signature="%s"`, base64.StdEncoding.EncodeToString(s.Signature)),
	}, ",")
}

// ParseSignatureHeader parses strings in the form of
// `keyId="<URL>",algorithm="ed25519",headers="(request-target) host date digest",signature="<BASE64 SIGNATURE>"`
func ParseSignatureHeader(raw string) (*SignatureHeader, error) {
	parts := strings.Split(raw, ",")
	if len(parts) != 4 {
		return nil, ErrInvalidSignatureHeader
	}

	sig := &SignatureHeader{}

	for _, part := range parts {
		kv := strings.SplitN(part, "=", 2)
		kv[1] = strings.TrimPrefix(kv[1], "\"")
		kv[1] = strings.TrimSuffix(kv[1], "\"")

		var err error

		switch kv[0] {
		case "keyId":
			sig.KeyID, err = url.Parse(kv[1])
		case "algorithm":
			sig.Algorithm = kv[1]
		case "headers":
			sig.Headers = kv[1]
		case "signature":
			sig.Signature, err = base64.StdEncoding.DecodeString(kv[1])
		}

		if err != nil {
			return nil, err
		}
	}

	return sig, nil
}
