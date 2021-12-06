package utils

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

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

func ReadInputToList(path string) []string {
	contents := ReadInput(path)
	return strings.Split(strings.TrimSpace(string(contents)), "\n")
}

func ReadInputToIntegerList(path string) []int {
	rows := ReadInputToList(path)
	return ToIntegers(rows)
}

func ReadCommaInputToIntegerList(path string) []int {
	data := ReadInput(path)
	return ToIntegers(strings.Split(strings.TrimSpace(string(data)), ","))
}
