package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay12Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"start-A",
			"start-b",
			"A-c",
			"A-b",
			"b-d",
			"A-end",
			"b-end",
		}

		expected := 10
		actual := Day12Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve the example #2", func(t *testing.T) {
		input := []string{
			"dc-end",
			"HN-start",
			"start-kj",
			"dc-start",
			"dc-HN",
			"LN-dc",
			"HN-end",
			"kj-sa",
			"kj-HN",
			"kj-dc",
		}

		expected := 19
		actual := Day12Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve the example #3", func(t *testing.T) {
		input := []string{
			"fs-end",
			"he-DX",
			"fs-he",
			"start-DX",
			"pj-DX",
			"end-zg",
			"zg-sl",
			"zg-pj",
			"pj-he",
			"RW-he",
			"fs-DX",
			"pj-RW",
			"zg-RW",
			"start-pj",
			"he-WI",
			"zg-he",
			"pj-fs",
			"start-RW",
		}

		expected := 226
		actual := Day12Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day12/input")

		expected := 5457
		actual := Day12Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay12Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"start-A",
			"start-b",
			"A-c",
			"A-b",
			"b-d",
			"A-end",
			"b-end",
		}

		expected := 36
		actual := Day12Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve the example #2", func(t *testing.T) {
		input := []string{
			"dc-end",
			"HN-start",
			"start-kj",
			"dc-start",
			"dc-HN",
			"LN-dc",
			"HN-end",
			"kj-sa",
			"kj-HN",
			"kj-dc",
		}

		expected := 103
		actual := Day12Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve the example #3", func(t *testing.T) {
		input := []string{
			"fs-end",
			"he-DX",
			"fs-he",
			"start-DX",
			"pj-DX",
			"end-zg",
			"zg-sl",
			"zg-pj",
			"pj-he",
			"RW-he",
			"fs-DX",
			"pj-RW",
			"zg-RW",
			"start-pj",
			"he-WI",
			"zg-he",
			"pj-fs",
			"start-RW",
		}

		expected := 3509
		actual := Day12Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day12/input")

		expected := 128506
		actual := Day12Part2(input)

		assert.Equal(t, expected, actual)
	})
}
