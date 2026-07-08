// Package filter provides a generic way to filter a slice of items.
package filter

// The Filter contains the slice of items to filter.
type Filter[T any] struct {
	items []*T
}

// Filter executes the filter operation with the given filters functions.
// Returns a slice of the matching items.
func (f *Filter[T]) Filter(filters ...func(T) bool) []T {
	filteredItems := make([]T, 0, len(f.items))

outer:
	for _, item := range f.items {
		for _, filterFn := range filters {
			if !filterFn(*item) {
				continue outer
			}
		}
		filteredItems = append(filteredItems, *item)
	}

	return filteredItems
}

// NewFilter creates a new filter with the given type and item slice.
func NewFilter[T any](items []*T) *Filter[T] {
	return &Filter[T]{items: items}
}
