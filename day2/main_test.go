package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}

		expected := 150
		actual := Day2Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay2Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}

		expected := 900
		actual := Day2Part2(input)

		assert.Equal(t, expected, actual)
	})
}
