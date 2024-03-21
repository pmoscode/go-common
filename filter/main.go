package filter

type Filter[T any] struct {
	items []*T
}

func (f *Filter[T]) Filter(filters ...func(T) bool) *[]T {
	filteredItems := make([]T, 0)

	for _, item := range f.items {
		skip := false
		for _, f := range filters {
			if !f(*item) {
				skip = true
				break
			}
		}
		if !skip {
			filteredItems = append(filteredItems, *item)
		}
	}

	return &filteredItems
}

func NewFilter[T any](items []*T) *Filter[T] {
	return &Filter[T]{items: items}
}
