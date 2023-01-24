package array

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

// Min returns the min Number in an array
func Min[T Number](data []T) T {
	var min T
	for i, v := range data {
		if v < min || i == 0 {
			min = v
		}
	}

	return min
}

// Max returns the max Number in an array
func Max[T Number](data []T) T {
	var max T
	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

// Sum sums all Numbers in an array
func Sum[T Number](data []T) T {
	var sum T
	for _, v := range data {
		sum += v
	}

	return sum
}

// Average returns the average for an array of Numbers
func Average[T Number](data []T) float64 {
	return float64(Sum(data)) / float64(len(data))
}
