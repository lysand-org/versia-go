package versiautils

import (
	"encoding/json"
	"net/url"
)

// URL is a type that represents a URL, represented by a string in JSON, instead of a JSON object.
type URL url.URL

func (u *URL) ResolveReference(ref *url.URL) *URL {
	return URLFromStd(u.ToStd().ResolveReference(ref))
}

func (u *URL) String() string {
	return u.ToStd().String()
}

func (u *URL) UnmarshalJSON(data []byte) error {
	raw := ""
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		return err
	}

	*u = URL(*parsed)

	return nil
}

func (u *URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *URL) ToStd() *url.URL {
	return (*url.URL)(u)
}

func URLFromStd(u *url.URL) *URL {
	return (*URL)(u)
}

func ParseURL(raw string) (*URL, error) {
	parsed, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	return URLFromStd(parsed), nil
}
