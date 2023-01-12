package array_test

import (
	"testing"

	"github.com/decanus/array"
)

func TestInsertAt(t *testing.T) {
	data := []int{3}

	data = array.InsertAt(data, 1, 4)
	data = array.InsertAt(data, 0, 1)
	data = array.InsertAt(data, 1, 2)

	if !array.Equal(data, []int{1, 2, 3, 4}) {
		t.Error("items were not inserted as expected")
	}
}

func TestEqualWhenArraysDoNotMatchInLength(t *testing.T) {
	if array.Equal([]int{1, 2}, []int{1, 2, 3, 4}) {
		t.Error("equal returned true when arrays do not match")
	}
}

func TestEqualWhenArraysDoNotMatch(t *testing.T) {
	if array.Equal([]int{1, 2, 3, 5}, []int{1, 2, 3, 4}) {
		t.Error("equal returned true when arrays do not match")
	}
}

func TestEqualWhenArraysMatch(t *testing.T) {
	if !array.Equal([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}) {
		t.Error("equal returned false when arrays match")
	}
}

func TestContains(t *testing.T) {
	data := []int{5, 4}

	if array.Contains(data, 3) {
		t.Error("contains returned true when item is not in array")
	}

	if !array.Contains(data, 4) {
		t.Error("contains returned false when item is in array")
	}
}
