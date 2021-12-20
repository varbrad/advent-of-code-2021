package utils

import (
	"errors"
	"math"
)

// AverageInts Returns the average of the provided integers, rounded to the nearest integer
func AverageInts(values []int) int {
	if len(values) == 0 {
		panic(errors.New("cannot average empty array"))
	}
	sum := SumInts(values)
	return RoundToInt(float64(sum) / float64(len(values)))
}

// RoundToInt Rounds the provided float64 to the nearest integer
func RoundToInt(value float64) int {
	return int(math.Round(value))
}

// MaxInteger Returns the larger of the two integers provided
func MaxInteger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInteger Returns the smaller of the two integers provided
func MinInteger(a, b int) int {
	if a > b {
		return b
	}
	return a
}
