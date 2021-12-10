package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
}
