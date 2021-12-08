package utils

import (
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func CountStringDiff(a string, b string) int {
	diff := 0
	for _, char := range a {
		if !StringContainsLetters(b, string(char)) {
			diff++
		}
	}
	for _, char := range b {
		if !StringContainsLetters(a, string(char)) {
			diff++
		}
	}
	return diff
}

func StringContainsLetters(str string, letters string) bool {
	for _, letter := range letters {
		if !strings.Contains(str, string(letter)) {
			return false
		}
	}
	return true
}
