package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay15Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := [][]int{
			{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
			{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
			{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
			{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
			{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
			{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
			{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
			{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
			{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
			{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
		}

		expected := 40
		actual := Day15Part1(input)

		assert.Equal(t, expected, actual)
	})

	// t.Run("Should solve my puzzle input", func(t *testing.T) {
	// 	input := utils.ReadInputToList("day14/input")

	// 	expected := 2937
	// 	actual := Day14Part1(input)

	// 	assert.Equal(t, expected, actual)
	// })
}

func TestDay15Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := [][]int{
			{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
			{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
			{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
			{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
			{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
			{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
			{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
			{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
			{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
			{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
		}

		expected := 315
		actual := Day15Part2(input)

		assert.Equal(t, expected, actual)
	})
}
