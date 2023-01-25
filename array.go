// Package array provides generic functions for working with go slices
package array

// InsertAt inserts v into data at index i and returns a new slice
// Adapted from: https://stackoverflow.com/questions/42746972/golang-insert-to-a-sorted-slice
func InsertAt[T any](data []T, i int, v T) []T {
	if i == len(data) {
		return append(data, v)
	}

	data = append(data[:i+1], data[i:]...)
	data[i] = v

	return data
}

// Equal compares 2 arrays to ensure they are equal
func Equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// Contains returns whether an element is contained in an array
func Contains[T comparable](data []T, v T) bool {
	for _, d := range data {
		if d == v {
			return true
		}
	}

	return false
}

// Map returns an array containing the results of a mapping
func Map[V, T any](data []V, fn func(int, V) (T, error)) ([]T, error) {
	res := make([]T, len(data))

	for i, v := range data {
		m, err := fn(i, v)
		if err != nil {
			return nil, err
		}

		res[i] = m
	}

	return res, nil
}
