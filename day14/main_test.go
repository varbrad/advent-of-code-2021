package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay14Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"NNCB",
			"",
			"CH -> B",
			"HH -> N",
			"CB -> H",
			"NH -> C",
			"HB -> C",
			"HC -> B",
			"HN -> C",
			"NN -> C",
			"BH -> H",
			"NC -> B",
			"NB -> B",
			"BN -> B",
			"BB -> N",
			"BC -> B",
			"CC -> N",
			"CN -> C",
		}

		expected := 1588
		actual := Day14Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day14/input")

		expected := 2937
		actual := Day14Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay14Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"NNCB",
			"",
			"CH -> B",
			"HH -> N",
			"CB -> H",
			"NH -> C",
			"HB -> C",
			"HC -> B",
			"HN -> C",
			"NN -> C",
			"BH -> H",
			"NC -> B",
			"NB -> B",
			"BN -> B",
			"BB -> N",
			"BC -> B",
			"CC -> N",
			"CN -> C",
		}

		expected := 2188189693529
		actual := Day14Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day14/input")

		expected := 3390034818249
		actual := Day14Part2(input)

		assert.Equal(t, expected, actual)
	})
}
