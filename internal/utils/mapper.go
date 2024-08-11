package utils

import "strings"

func MapSlice[T any, V any](obj []T, transform func(T) V) []V {
	vs := make([]V, 0, len(obj))

	for _, o := range obj {
		vs = append(vs, transform(o))
	}

	return vs
}

type CombinedError struct {
	Errors []error
}

func (e CombinedError) Error() string {
	sb := strings.Builder{}

	for i, err := range e.Errors {
		sb.WriteString(err.Error())

		if i < len(e.Errors)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func MapErrorSlice[T any, V any](obj []T, transform func(T) (V, error)) ([]V, error) {
	vs := make([]V, 0, len(obj))
	errs := make([]error, 0, len(obj))

	for _, o := range obj {
		v, err := transform(o)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		vs = append(vs, v)
	}

	if len(errs) > 0 {
		return nil, CombinedError{Errors: errs}
	}

	return vs, nil
}
