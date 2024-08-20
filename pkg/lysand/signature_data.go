package lysand

import (
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	versiacrypto "github.com/lysand-org/versia-go/pkg/lysand/crypto"
	"net/url"
	"os"
	"strings"
)

type SignatureData struct {
	// RequestMethod is the *lowercase* HTTP method of the request
	RequestMethod string
	// Nonce is a random byte array, used to prevent replay attacks
	Nonce string
	// RawPath is the path of the request, without the query string
	URL *url.URL
	// Digest is the SHA-256 hash of the request body
	Digest []byte
}

func NewSignatureData(method, nonce string, u *url.URL, digest []byte) *SignatureData {
	return &SignatureData{
		RequestMethod: method,
		Nonce:         nonce,
		URL:           u,
		Digest:        digest,
	}
}

func (s *SignatureData) String() string {
	return fmt.Sprintf("%s %s?%s %s %s", strings.ToLower(s.RequestMethod), s.URL.Path, s.URL.RawQuery, s.Nonce, base64.StdEncoding.EncodeToString(s.Digest))
}

func (s *SignatureData) Validate(pubKey crypto.PublicKey, signature []byte) bool {
	data := []byte(s.String())

	verify, err := versiacrypto.NewVerify(pubKey)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return false
	}

	return verify(data, signature)
}

func (s *SignatureData) Sign(privKey ed25519.PrivateKey) []byte {
	return ed25519.Sign(privKey, []byte(s.String()))
}

type Signer struct {
	PrivateKey ed25519.PrivateKey
	UserURL    *url.URL
}

func (s Signer) Sign(signatureData SignatureData) *FederationHeaders {
	return &FederationHeaders{
		SignedBy:  s.UserURL,
		Nonce:     signatureData.Nonce,
		Signature: signatureData.Sign(s.PrivateKey),
	}
}

type Verifier struct {
	PublicKey crypto.PublicKey
}

func (v Verifier) Verify(method string, u *url.URL, body []byte, fedHeaders *FederationHeaders) bool {
	return NewSignatureData(method, fedHeaders.Nonce, u, versiacrypto.SHA256(body)).
		Validate(v.PublicKey, fedHeaders.Signature)
}
