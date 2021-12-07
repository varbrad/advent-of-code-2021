package utils

import "math"

func AverageInts(values []int) int {
	sum := SumInts(values)
	return int(math.Round(float64(sum) / float64(len(values))))
}
