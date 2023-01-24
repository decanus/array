package array_test

import (
	"testing"

	"github.com/decanus/array"
)

func TestMin(t *testing.T) {
	data := []int{4, 2, 3, 6, 1}

	if array.Min(data) != 1 {
		t.Error("incorrect max returned")
	}
}

func TestMax(t *testing.T) {
	data := []int{4, 2, 3, 6, 1}

	if array.Max(data) != 6 {
		t.Error("incorrect max returned")
	}
}

func TestAverage(t *testing.T) {
	data := []int{4, 2, 3, 6, 1}

	if array.Average(data) != 3.2 {
		t.Error("incorrect average returned")
	}
}
