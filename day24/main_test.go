package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay24Part1(t *testing.T) {
	t.Run("Should validate my solution", func(t *testing.T) {
		input := utils.ReadInputToList("day24/input")

		expected := 96299896449997
		actual := Day24Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay24Part2(t *testing.T) {
	t.Run("Should validate my solution", func(t *testing.T) {
		input := utils.ReadInputToList("day24/input")

		expected := 31162141116841
		actual := Day24Part2(input)

		assert.Equal(t, expected, actual)
	})
}
