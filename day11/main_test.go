package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay11Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := [][]int{
			{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
			{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
			{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
			{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
			{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
			{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
			{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
			{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
			{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
			{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
		}

		expected := 1656
		actual := Day11Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToIntGrid("day11/input")

		expected := 1588
		actual := Day11Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay11Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := [][]int{
			{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
			{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
			{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
			{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
			{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
			{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
			{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
			{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
			{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
			{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
		}

		expected := 195
		actual := Day11Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToIntGrid("day11/input")

		expected := 517
		actual := Day11Part2(input)

		assert.Equal(t, expected, actual)
	})
}
