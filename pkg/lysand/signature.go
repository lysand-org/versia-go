package lysand

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type SignatureData struct {
	RequestMethod string
	Date          time.Time
	Host          string
	Path          string
	Digest        []byte
}

func NewSignatureData(method string, date time.Time, host, path string, digest []byte) *SignatureData {
	return &SignatureData{
		RequestMethod: method,
		Date:          date,
		Host:          host,
		Path:          path,
		Digest:        digest,
	}
}

func (s *SignatureData) String() string {
	return strings.Join([]string{
		fmt.Sprintf("(request-target): %s %s", strings.ToLower(s.RequestMethod), s.Path),
		fmt.Sprintf("host: %s", s.Host),
		fmt.Sprintf("date: %s", TimeFromStd(s.Date).String()),
		fmt.Sprintf("digest: SHA-256=%s", base64.StdEncoding.EncodeToString(s.Digest)),
		"",
	}, "\n")
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

func (s Signer) Sign(signatureData SignatureData) *SignatureHeader {
	return &SignatureHeader{
		KeyID:     s.UserURL,
		Algorithm: "ed25519",
		Headers:   "(request-target) host date digest",
		Signature: signatureData.Sign(s.PrivateKey),
	}
}

type Verifier struct {
	PublicKey ed25519.PublicKey
}

func (v Verifier) Verify(method string, date time.Time, host, path string, body []byte, sigHeader *SignatureHeader) bool {
	sigData := NewSignatureData(method, date, host, path, hashSHA256(body))

	return sigData.Validate(v.PublicKey, sigHeader.Signature)
}
