package versiautils

import (
	"strings"
)

type MultipleError struct {
	Errors []error
}

func (e MultipleError) Error() string {
	s := strings.Builder{}
	for i, err := range e.Errors {
		s.WriteString(err.Error())

		if i != len(e.Errors) {
			s.WriteRune('\n')
		}
	}

	return s.String()
}
