package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageInts(t *testing.T) {
	t.Run("Should average integers with an integer average", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}

		expected := 3
		actual := AverageInts(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should average integers with a float average", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}

		expected := 4
		actual := AverageInts(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should panic 0 if an empty list", func(t *testing.T) {
		input := []int{}

		assert.PanicsWithError(t, "cannot average empty array", func() { AverageInts(input) })
	})
}
