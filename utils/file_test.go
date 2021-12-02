package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputCanParseAFile(t *testing.T) {
	expected := "1\n2\n3\n"
	actual := ReadInput("utils/__tests__/list")

	assert.Equal(t, expected, string(actual))
}

func TestReadInputHandlesNonExistantFile(t *testing.T) {
	assert.Panics(t, func() { ReadInput("utils/__tests__/non-existant-file") })
}

func TestReadInputToListCanParseAList(t *testing.T) {
	expected := []string{"1", "2", "3"}
	actual := ReadInputToList("utils/__tests__/list")

	assert.Equal(t, expected, actual)
}

func TestReadInputToListHandlesNonExistantFile(t *testing.T) {
	assert.Panics(t, func() { ReadInputToList("utils/__tests__/non-existant-file") })
}

func TestReadInputToIntegerListCanParseAList(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := ReadInputToIntegerList("utils/__tests__/list")

	assert.Equal(t, expected, actual)
}

func TestReadInputToIntegerListHandlesNonExistantFile(t *testing.T) {
	assert.Panics(t, func() { ReadInputToIntegerList("utils/__tests__/non-existant-file") })
}

func TestReadInputToIntegerListHandlesInvalidList(t *testing.T) {
	assert.Panics(t, func() { ReadInputToIntegerList("utils/__tests__/string-list") })
}
