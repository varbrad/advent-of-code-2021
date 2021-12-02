package utils

import (
	"strconv"
)

func ToInteger(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return value
}

func ToIntegers(stringList []string) []int {
	ints := make([]int, len(stringList))
	for i, s := range stringList {
		value := ToInteger(s)
		ints[i] = value
	}
	return ints
}
