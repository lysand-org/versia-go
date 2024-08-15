package lysand

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"net/url"
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

func (s *SignatureData) Validate(pubKey ed25519.PublicKey, signature []byte) bool {
	return ed25519.Verify(pubKey, []byte(s.String()), signature)
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
	PublicKey ed25519.PublicKey
}

func (v Verifier) Verify(method string, u *url.URL, body []byte, fedHeaders *FederationHeaders) bool {
	sigData := NewSignatureData(method, fedHeaders.Nonce, u, hashSHA256(body))

	return sigData.Validate(v.PublicKey, fedHeaders.Signature)
}
