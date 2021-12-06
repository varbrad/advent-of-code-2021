package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay6Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{3, 4, 3, 1, 2}

		expected := 5934
		actual := Day6Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadCommaInputToIntegerList("day6/input")

		expected := 387413
		actual := Day6Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay6Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{3, 4, 3, 1, 2}

		expected := 26984457539
		actual := Day6Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadCommaInputToIntegerList("day6/input")

		expected := 1738377086345
		actual := Day6Part2(input)

		assert.Equal(t, expected, actual)
	})
}
