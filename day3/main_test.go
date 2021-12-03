package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}

		expected := 198
		actual := Day3Part1(input)

		assert.Equal(t, expected, actual)
	})
}
