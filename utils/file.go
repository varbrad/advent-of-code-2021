package utils

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

// ReadInput  Low-level API to read bytes from a filepath
func ReadInput(path string) []byte {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "..")
	fullPath := filepath.Join(basepath, path)
	bytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}
	return bytes
}

// ReadInputToString Reads a file and returns a trimmed string of the contents
func ReadInputToString(path string) string {
	contents := ReadInput(path)
	return strings.TrimSpace(string(contents))
}

// ReadInputToList Reads a file and returns an array of trimmed strings, split by newlines
func ReadInputToList(path string) []string {
	contents := ReadInput(path)
	return strings.Split(strings.TrimSpace(string(contents)), "\n")
}

// ReadInputToIntegerList Similar to ReadInputToList, but casts all values to integers
func ReadInputToIntegerList(path string) []int {
	rows := ReadInputToList(path)
	return ToIntegers(rows)
}

// ReadCommaInputToIntegerList Reads a file and returns an array of integers, delimited by commas
func ReadCommaInputToIntegerList(path string) []int {
	data := ReadInput(path)
	return ToIntegers(strings.Split(strings.TrimSpace(string(data)), ","))
}

// ReadInputToIntGrid Reads a file and returns a 2D array of digits, split by newlines
func ReadInputToIntGrid(path string) [][]int {
	rows := ReadInputToList(path)
	ints := make([][]int, len(rows))
	for ix, row := range rows {
		ints[ix] = ToIntegers(strings.Split(strings.TrimSpace(row), ""))
	}
	return ints
}
