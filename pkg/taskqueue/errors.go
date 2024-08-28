package taskqueue

import "strings"

type CombinedError struct {
	Errors []error
}

func (e CombinedError) Error() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	for i, err := range e.Errors {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(err.Error())
	}
	sb.WriteRune(']')
	return sb.String()
}
