package array

import "sort"

// InsertSorted inserts v into data sorted by function fn
func InsertSorted[T any](data []T, v T, fn func(a, b T) bool) []T {
	i := sort.Search(len(data), func(i int) bool {
		return fn(data[i], v)
	})

	return InsertAt(data, i, v)
}
