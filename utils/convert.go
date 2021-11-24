package utils

import (
	"strconv"
)

func ToInteger(text string) (int, error) {
	return strconv.Atoi(text)
}

func ToIntegers(stringList []string) ([]int, error) {
	ints := make([]int, len(stringList))
	for i, s := range stringList {
		value, err := ToInteger(s)
		if err != nil {
			return nil, err
		}
		ints[i] = value
	}
	return ints, nil
}
