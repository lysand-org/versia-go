package versiautils

import (
	"encoding/json"
	"time"
)

const ISO8601 = "2006-01-02T15:04:05.000Z"

func ParseTime(s string) (Time, error) {
	t, err := time.Parse(ISO8601, s)
	return Time(t), err
}

// Time is a type that represents a time in the ISO8601 format.
type Time time.Time

// String returns the time in the ISO8601 format.
func (t Time) String() string {
	return t.ToStd().Format(ISO8601)
}

// UnmarshalJSON decodes the time from a string in the ISO8601 format.
func (t *Time) UnmarshalJSON(data []byte) error {
	raw := ""
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	parsed, err := time.Parse(ISO8601, raw)
	if err != nil {
		return err
	}

	*t = Time(parsed)

	return nil
}

// MarshalJSON marshals the time to a string in the ISO8601 format.
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// ToStd converts the time to a [time.Time].
func (t Time) ToStd() time.Time {
	return time.Time(t)
}

// TimeFromStd converts a [time.Time] to a Time.
func TimeFromStd(u time.Time) Time {
	return Time(u)
}

func TimeNow() Time {
	return Time(time.Now())
}
