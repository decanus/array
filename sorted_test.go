package array_test

import (
	"math/rand"
	"testing"

	"github.com/decanus/array"
)

var sortFunc = func(a, b int) bool {
	return a >= b
}

func TestInsertSorted(t *testing.T) {

	data := array.InsertSorted(make([]int, 0), 3, sortFunc)
	data = array.InsertSorted(data, 2, sortFunc)
	data = array.InsertSorted(data, 4, sortFunc)
	data = array.InsertSorted(data, 1, sortFunc)
	data = array.InsertSorted(data, 5, sortFunc)
	data = array.InsertSorted(data, 6, sortFunc)

	if !array.Equal(data, []int{1, 2, 3, 4, 5, 6}) {
		t.Error("array was not sorted correctly")
	}
}

func BenchmarkInsertSorted(b *testing.B) {
	data := make([]int, 0)

	for i := 0; i < b.N; i++ {
		v := b.N + 2
		data = array.InsertSorted(data, rand.Intn(v/2), sortFunc)
	}
}
