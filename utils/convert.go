package utils

import (
	"strconv"
)

// ToInteger Forces a string to an integer, else panics
func ToInteger(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return value
}

// ToIntegers Forces a slice of strings to integers, else panics
func ToIntegers(stringList []string) []int {
	ints := make([]int, len(stringList))
	for i, s := range stringList {
		value := ToInteger(s)
		ints[i] = value
	}
	return ints
}
