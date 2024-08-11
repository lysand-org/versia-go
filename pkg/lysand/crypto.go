package lysand

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func (c *FederationClient) ValidateSignatureHeader(req *http.Request) (bool, error) {
	date, sigHeader, err := ExtractFederationHeaders(req.Header)
	if err != nil {
		return false, err
	}

	// TODO: Fetch user from database instead of using the URI
	user, err := c.GetUser(req.Context(), sigHeader.KeyID)
	if err != nil {
		return false, err
	}

	body, err := copyBody(req)
	if err != nil {
		return false, err
	}

	v := Verifier{ed25519.PublicKey(user.PublicKey.PublicKey)}
	valid := v.Verify(req.Method, date, req.Host, req.URL.Path, body, sigHeader)

	return valid, nil
}

func ExtractFederationHeaders(h http.Header) (time.Time, *SignatureHeader, error) {
	gotDates := h.Values("date")
	var date *Time
	for i, raw := range gotDates {
		if parsed, err := ParseTime(raw); err != nil {
			log.Printf("invalid date[%d] header: %s", i, raw)
			continue
		} else {
			date = &parsed
			break
		}
	}
	if date == nil {
		return time.Time{}, nil, fmt.Errorf("missing date header")
	}

	gotSignature := h.Get("signature")
	if gotSignature == "" {
		return date.ToStd(), nil, fmt.Errorf("missing signature header")
	}
	sigHeader, err := ParseSignatureHeader(gotSignature)
	if err != nil {
		return date.ToStd(), nil, err
	}

	return date.ToStd(), sigHeader, nil
}

func hashSHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func must[In any, Out any](fn func(In) (Out, error), v In) Out {
	out, err := fn(v)
	if err != nil {
		panic(err)
	}

	return out
}

func copyBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if err := req.Body.Close(); err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}
