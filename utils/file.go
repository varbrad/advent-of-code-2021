package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadInput(path string) []byte {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fullPath := filepath.Join(cwd, path)
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
