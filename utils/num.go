package utils

import (
	"errors"
	"math"
)

func AverageInts(values []int) int {
	if len(values) == 0 {
		panic(errors.New("cannot average empty array"))
	}
	sum := SumInts(values)
	return RoundToInt(float64(sum) / float64(len(values)))
}

func RoundToInt(value float64) int {
	return int(math.Round(value))
}
