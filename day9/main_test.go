package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay9Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := [][]int{
			{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
			{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
			{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
			{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
			{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
		}

		expected := 15
		actual := Day9Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToIntGrid("day9/input")

		expected := 554
		actual := Day9Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay9Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := [][]int{
			{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
			{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
			{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
			{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
			{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
		}

		expected := 1134
		actual := Day9Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToIntGrid("day9/input")

		expected := 1017792
		actual := Day9Part2(input)

		assert.Equal(t, expected, actual)
	})
}
