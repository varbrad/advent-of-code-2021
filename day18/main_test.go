package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay18Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
			"[[[5,[2,8]],4],[5,[[9,9],0]]]",
			"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
			"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
			"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
			"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
			"[[[[5,4],[7,7]],8],[[8,3],8]]",
			"[[9,3],[[9,9],[6,[4,9]]]]",
			"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
			"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
		}

		expected := 4140
		actual := Day18Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day18/input")

		expected := 4145
		actual := Day18Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay18Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
			"[[[5,[2,8]],4],[5,[[9,9],0]]]",
			"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
			"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
			"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
			"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
			"[[[[5,4],[7,7]],8],[[8,3],8]]",
			"[[9,3],[[9,9],[6,[4,9]]]]",
			"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
			"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
		}

		expected := 3993
		actual := Day18Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day18/input")

		expected := 4855
		actual := Day18Part2(input)

		assert.Equal(t, expected, actual)
	})
}
