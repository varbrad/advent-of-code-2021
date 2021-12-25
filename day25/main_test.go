package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay25Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"v...>>.vv>",
			".vv>>.vv..",
			">>.>v>...v",
			">>v>>.>.v.",
			"v>v.vv.v..",
			">.>>..v...",
			".vv..>.>v.",
			"v.v..>>v.v",
			"....v..v.>",
		}

		expected := 58
		actual := Day25Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my actual puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day25/input")

		expected := 432
		actual := Day25Part1(input)

		assert.Equal(t, expected, actual)
	})
}
