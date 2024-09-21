package versiacrypto

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

// FederationHeaders represents the signature header of the Versia protocol. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/signatures#signature-definition
type FederationHeaders struct {
	// SignedBy is the URL to a user, or `nil` when the signature was created by the instance's privatekey
	SignedBy *url.URL
	// Nonce is a random string, used to prevent replay attacks
	Nonce string
	// Signature is the signature of the request
	Signature []byte
}

func (f *FederationHeaders) Inject(h http.Header) {
	signedBy := "instance"
	if f.SignedBy != nil {
		signedBy = f.SignedBy.String()
	}
	h.Set("x-signed-by", signedBy)
	h.Set("x-nonce", f.Nonce)
	h.Set("x-signature", base64.StdEncoding.EncodeToString(f.Signature))
}

func (f *FederationHeaders) Headers() map[string]string {
	signedBy := "instance"
	if f.SignedBy != nil {
		signedBy = f.SignedBy.String()
	}

	return map[string]string{
		"x-signed-by": signedBy,
		"x-nonce":     f.Nonce,
		"x-signature": base64.StdEncoding.EncodeToString(f.Signature),
	}
}

func ExtractFederationHeaders(h http.Header) (*FederationHeaders, error) {
	signedBy := h.Get("x-signed-by")
	if signedBy == "" {
		return nil, fmt.Errorf("missing x-signed-by header")
	}

	var u *url.URL
	if signedBy == "instance" {
		// Signed by the instance
	} else {
		var err error
		u, err = url.Parse(signedBy)
		if err != nil {
			return nil, err
		}
	}

	nonce := h.Get("x-nonce")
	if nonce == "" {
		return nil, fmt.Errorf("missing x-nonce header")
	}

	rawSignature := h.Get("x-signature")
	if rawSignature == "" {
		return nil, fmt.Errorf("missing x-signature header")
	}

	signature, err := base64.StdEncoding.DecodeString(rawSignature)
	if err != nil {
		return nil, err
	}

	return &FederationHeaders{
		SignedBy:  u,
		Nonce:     nonce,
		Signature: signature,
	}, nil
}
