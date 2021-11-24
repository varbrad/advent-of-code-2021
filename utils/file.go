package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadInput(path string) ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fullPath := filepath.Join(cwd, path)
	return ioutil.ReadFile(fullPath)
}

func ReadInputToList(path string) ([]string, error) {
	contents, err := ReadInput(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(contents)), "\n"), nil
}

func ReadInputToIntegerList(path string) ([]int, error) {
	rows, err := ReadInputToList(path)
	if err != nil {
		return nil, err
	}
	return ToIntegers(rows)
}
