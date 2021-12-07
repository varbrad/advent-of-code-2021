package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumInts(t *testing.T) {
	t.Run("Can sum integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}

		expected := 15
		actual := SumInts(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Returns 0 if an empty list", func(t *testing.T) {
		input := []int{}

		expected := 0
		actual := SumInts(input)

		assert.Equal(t, expected, actual)
	})
}
