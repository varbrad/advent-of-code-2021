package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSnailfishNumber(t *testing.T) {
	inputs := []string{
		"[1,2]",
		"[1,1]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}

	for _, input := range inputs {
		actual := parseSnailfishNumber(input).String()
		assert.Equal(t, input, actual)
	}
}

func TestAddSnailfishNumber(t *testing.T) {
	a := parseSnailfishNumber("[1,2]")
	b := parseSnailfishNumber("[3,[4,5]]")

	result := a.Add(b)

	assert.Equal(t, result.left, a)
	assert.Equal(t, result.right, b)
	assert.Equal(t, result.String(), "[[1,2],[3,[4,5]]]")
}

func TestReduceSnailfishNumber(t *testing.T) {
	t.Run("Should work for a non-applicable example", func(t *testing.T) {
		input := "[1,2]"

		expected := "[1,2]"
		actual := parseSnailfishNumber(input).Reduce().String()

		assert.Equal(t, expected, actual)
	})

	t.Run("Should work for a simple example #1", func(t *testing.T) {
		input := "[[[[[9,8],1],2],3],4]"

		expected := "[[[[0,9],2],3],4]"
		actual := parseSnailfishNumber(input).Reduce().String()

		assert.Equal(t, expected, actual)
	})

	t.Run("Should work for a simple example #2", func(t *testing.T) {
		input := "[[6,[5,[4,[3,2]]]],1]"

		expected := "[[6,[5,[7,0]]],3]"
		actual := parseSnailfishNumber(input).Reduce().String()

		assert.Equal(t, expected, actual)
	})

	t.Run("Should work for a simple example #3", func(t *testing.T) {
		input := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"

		expected := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
		actual := parseSnailfishNumber(input).Reduce().String()

		assert.Equal(t, expected, actual)
	})
}

func TestSnailfishMagnitude(t *testing.T) {
	t.Run("Should work for a simple example #1", func(t *testing.T) {
		input := "[[1,2],[[3,4],5]]"

		expected := 143
		actual := parseSnailfishNumber(input).Magnitude()

		assert.Equal(t, expected, actual)
	})

	t.Run("Should work for a simple example #2", func(t *testing.T) {
		input := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"

		expected := 3488
		actual := parseSnailfishNumber(input).Magnitude()

		assert.Equal(t, expected, actual)
	})
}
