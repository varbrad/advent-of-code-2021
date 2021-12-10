package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay7Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

		expected := 37
		actual := Day7Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadCommaInputToIntegerList("day7/input")

		expected := 348664
		actual := Day7Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay7Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

		expected := 168
		actual := Day7Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadCommaInputToIntegerList("day7/input")

		expected := 100220525
		actual := Day7Part2(input)

		assert.Equal(t, expected, actual)
	})
}
