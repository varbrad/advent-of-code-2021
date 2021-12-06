package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{3, 4, 3, 1, 2}

		expected := 5934
		actual := Day6Part1(input)

		assert.Equal(t, expected, actual)
	})
}
