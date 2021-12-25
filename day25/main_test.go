package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay22Part1(t *testing.T) {
	t.Run("Should solve an example #1", func(t *testing.T) {
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
