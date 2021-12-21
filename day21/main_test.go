package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay21Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{4, 8}

		expected := 739785
		actual := Day21Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my actual puzzle input", func(t *testing.T) {
		input := []int{10, 4}

		expected := 908091
		actual := Day21Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay21Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{4, 8}

		expected := 444356092776315
		actual := Day21Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my actual puzzle input", func(t *testing.T) {
		input := []int{10, 4}

		expected := 190897246590017
		actual := Day21Part2(input)

		assert.Equal(t, expected, actual)
	})
}
