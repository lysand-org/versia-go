package lysand

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

// FederationHeaders represents the signature header of the Lysand protocol. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/signatures#signature-definition
type FederationHeaders struct {
	// SignedBy is the URL to a user
	SignedBy *url.URL
	// Nonce is a random string, used to prevent replay attacks
	Nonce string
	// Signature is the signature of the request
	Signature []byte
}

func (f *FederationHeaders) Inject(h http.Header) {
	h.Set("x-signed-by", f.SignedBy.String())
	h.Set("x-nonce", f.Nonce)
	h.Set("x-signature", base64.StdEncoding.EncodeToString(f.Signature))
}

func (f *FederationHeaders) Headers() map[string]string {
	return map[string]string{
		"x-signed-by": f.SignedBy.String(),
		"x-nonce":     f.Nonce,
		"x-signature": base64.StdEncoding.EncodeToString(f.Signature),
	}
}

func ExtractFederationHeaders(h http.Header) (*FederationHeaders, error) {
	signedBy := h.Get("x-signed-by")
	if signedBy == "" {
		return nil, fmt.Errorf("missing x-signed-by header")
	}

	u, err := url.Parse(signedBy)
	if err != nil {
		return nil, err
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
