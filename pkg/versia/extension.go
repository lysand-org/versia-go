package versia

import (
	"fmt"
	"strings"
)

// Extensions represents the active extensions on an entity. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/extensions#extension-definition
type Extensions map[ExtensionKey]any

// ExtensionKey represents an extension's key. For more information see the [Spec].
//
// [Spec]: https://versia.pub/types#extensions
type ExtensionKey [2]string

func (e *ExtensionKey) UnmarshalText(b []byte) (err error) {
	raw := string(b)

	spl := strings.Split(raw, ":")
	if len(spl) != 2 {
		return InvalidExtensionKeyError{raw}
	}

	*e = [2]string{spl[0], spl[1]}

	return
}

func (e ExtensionKey) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s:%s", e[0], e[1])), nil
}

type InvalidExtensionKeyError struct {
	Raw string
}

func (e InvalidExtensionKeyError) Error() string {
	return fmt.Sprintf("invalid extension key: %s", e.Raw)
}
