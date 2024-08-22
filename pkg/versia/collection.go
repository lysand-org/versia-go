package versia

import versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"

// Collection is a paginated group of entities. For more information, see the [Spec].
//
// [Spec]: https://versia.pub/structures/collection
type Collection[T any] struct {
	// Author represents the author of the collection. `nil` is used to represent the instance.
	Author *versiautils.URL `json:"author"`

	// First is a URI to the first page of the collection.
	First *versiautils.URL `json:"first"`

	// Last is a URI to the last page of the collection.
	// If the collection only has one page, this should be the same as First.
	Last *versiautils.URL `json:"last"`

	// Total is a count of all entities in the collection across all pages.
	Total uint64 `json:"total"`

	// Next is a URI to the next page of the collection. If there's no next page, this should be `nil`.
	Next *versiautils.URL `json:"next"`

	// Previous is a URI to the previous page of the collection. If there's no next page, this should be `nil`.
	// FIXME(spec): The spec uses `prev` instead of `previous` as the field name.
	Previous *versiautils.URL `json:"previous"`

	// Items is a list of T for the current page of the collection.
	Items []T `json:"items"`
}
