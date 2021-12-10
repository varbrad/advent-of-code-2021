package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay10Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"[({(<(())[]>[[{[]{<()<>>",
			"[(()[<>])]({[<{<<[]>>(",
			"{([(<{}[<>[]}>{[]{[(<()>",
			"(((({<>}<{<{<>}{[]{[]{}",
			"[[<[([]))<([[{}[[()]]]",
			"[{[{({}]{}}([{[{{{}}([]",
			"{<[[]]>}<{[{[{[]{()[[[]",
			"[<(<(<(<{}))><([]([]()",
			"<{([([[(<>()){}]>(<<{{",
			"<{([{{}}[<[[[<>{}]]]>[]]",
		}

		expected := 26397
		actual := Day10Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day10/input")

		expected := 369105
		actual := Day10Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay10Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"[({(<(())[]>[[{[]{<()<>>",
			"[(()[<>])]({[<{<<[]>>(",
			"{([(<{}[<>[]}>{[]{[(<()>",
			"(((({<>}<{<{<>}{[]{[]{}",
			"[[<[([]))<([[{}[[()]]]",
			"[{[{({}]{}}([{[{{{}}([]",
			"{<[[]]>}<{[{[{[]{()[[[]",
			"[<(<(<(<{}))><([]([]()",
			"<{([([[(<>()){}]>(<<{{",
			"<{([{{}}[<[[[<>{}]]]>[]]",
		}

		expected := 288957
		actual := Day10Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day10/input")

		expected := 3999363569
		actual := Day10Part2(input)

		assert.Equal(t, expected, actual)
	})
}
