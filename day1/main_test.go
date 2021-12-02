package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay1Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{
			199,
			200,
			208,
			210,
			200,
			207,
			240,
			269,
			260,
			263,
		}

		expected := 7
		actual := Day1Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToIntegerList("day1/input")

		expected := 1713
		actual := Day1Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay1Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{
			199,
			200,
			208,
			210,
			200,
			207,
			240,
			269,
			260,
			263,
		}

		expected := 5
		actual := Day1Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToIntegerList("day1/input")

		expected := 1734
		actual := Day1Part2(input)

		assert.Equal(t, expected, actual)
	})
}
