package versia

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtensionKey_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		Raw      string
		Expected ExtensionKey
		Error    bool
	}{
		{"\"pub.versia:Emoji\"", ExtensionKey{"pub.versia", "Emoji"}, false},
		{"\"pub.versia\"", ExtensionKey{}, true},
	}

	for _, case_ := range cases {
		key := ExtensionKey{}
		err := json.Unmarshal([]byte(case_.Raw), &key)

		assert.Equal(t, case_.Error, err != nil, "error assertion should match")

		if !case_.Error {
			assert.Equal(t, case_.Expected, key)
		}
	}
}
