package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay13Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"6,10",
			"0,14",
			"9,10",
			"0,3",
			"10,4",
			"4,11",
			"6,0",
			"6,12",
			"4,1",
			"0,13",
			"10,12",
			"3,4",
			"3,0",
			"8,4",
			"1,10",
			"2,14",
			"8,10",
			"9,0",
			"",
			"fold along y=7",
			"fold along x=5",
		}

		expected := 17
		actual := Day13Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day13/input")

		expected := 775
		actual := Day13Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay13Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"6,10",
			"0,14",
			"9,10",
			"0,3",
			"10,4",
			"4,11",
			"6,0",
			"6,12",
			"4,1",
			"0,13",
			"10,12",
			"3,4",
			"3,0",
			"8,4",
			"1,10",
			"2,14",
			"8,10",
			"9,0",
			"",
			"fold along y=7",
			"fold along x=5",
		}

		expected := "See above!"
		actual := Day13Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day13/input")

		expected := "See above!"
		actual := Day13Part2(input)

		assert.Equal(t, expected, actual)
	})
}
