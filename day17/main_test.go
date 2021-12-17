package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{20, 30, -10, -5}

		expected := 45
		actual := Day17Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := []int{138, 184, -125, -71}

		expected := 7750
		actual := Day17Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay17Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []int{20, 30, -10, -5}

		expected := 112
		actual := Day17Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := []int{138, 184, -125, -71}

		expected := 4120
		actual := Day17Part2(input)

		assert.Equal(t, expected, actual)
	})
}
