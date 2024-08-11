package helpers

import (
	"crypto/sha256"
	"time"
)

func HashSHA256(data []byte) []byte {
	h := sha256.New()

	h.Write(data)

	return h.Sum(nil)
}

func ISO8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}

func ParseISO8601(s string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z", s)
}
