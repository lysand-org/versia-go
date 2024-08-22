package versiacrypto

import (
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// SignatureData is a combination of HTTP method, URL (only url.URL#Path and url.URL#RawQuery are required),
// a nonce and the Base64 encoded SHA256 hash of the request body.
// For more information, see the [Spec].
//
// [Spec]: https://versia.pub/signatures
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

// String returns the payload to sign
func (s *SignatureData) String() string {
	return fmt.Sprintf("%s %s?%s %s %s", strings.ToLower(s.RequestMethod), s.URL.Path, s.URL.RawQuery, s.Nonce, base64.StdEncoding.EncodeToString(s.Digest))
}

// Validate validate that the SignatureData belongs to the provided public key and matches the provided signature.
func (s *SignatureData) Validate(pubKey crypto.PublicKey, signature []byte) bool {
	data := []byte(s.String())

	verify, err := NewVerify(pubKey)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return false
	}

	return verify(data, signature)
}

// Sign signs the SignatureData with the provided private key.
func (s *SignatureData) Sign(privKey ed25519.PrivateKey) []byte {
	return ed25519.Sign(privKey, []byte(s.String()))
}

// Signer is an object, with which requests can be signed with the user's private key.
// For more information, see the [Spec].
//
// [Spec]: https://versia.pub/signatures
type Signer struct {
	PrivateKey ed25519.PrivateKey
	UserURL    *url.URL
}

// Sign signs a signature data and returns the headers to inject into the response.
func (s Signer) Sign(signatureData SignatureData) *FederationHeaders {
	return &FederationHeaders{
		SignedBy:  s.UserURL,
		Nonce:     signatureData.Nonce,
		Signature: signatureData.Sign(s.PrivateKey),
	}
}

// Verifier is an object, with which requests can be verified against a user's public key.
// For more information, see the [Spec].
//
// [Spec]: https://versia.pub/signatures
type Verifier struct {
	PublicKey crypto.PublicKey
}

// Verify verifies a request against the public key provided to it duration object creation.
func (v Verifier) Verify(method string, u *url.URL, body []byte, fedHeaders *FederationHeaders) bool {
	return NewSignatureData(method, fedHeaders.Nonce, u, SHA256(body)).
		Validate(v.PublicKey, fedHeaders.Signature)
}
