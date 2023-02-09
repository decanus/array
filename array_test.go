package array_test

import (
	"fmt"
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

func TestMap(t *testing.T) {
	data := []int{5, 4}

	mapped, err := array.Map(data, func(_, v int) (string, error) {
		return fmt.Sprintf("%d", v), nil
	})

	if err != nil {
		t.Error("unexpected error in mapping")
	}

	if !array.Equal(mapped, []string{"5", "4"}) {
		t.Error("mapping did not return expected result")
	}
}

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	filtered := array.Filter(data, func(_, n int) bool {
		return n%2 == 0
	})

	if !array.Equal(filtered, []int{2, 4, 6, 8, 10}) {
		t.Error("filter did not return expected result")
	}
}

func TestReduce(t *testing.T) {
	data := []int{1, 2, 3, 4}

	res := array.Reduce(0, data, func(x, y int) int {
		return x + y
	})

	if res != 10 {
		t.Error("reduce did not return expected result")
	}
}

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3, 4, 9, 12}
	b := []int{1, 2, 3, 4, 5, 6, 7}

	res := array.Intersection(a, b)

	if !array.Equal(res, []int{1, 2, 3, 4}) {
		t.Error("intersection did not return expected result")
	}
}
