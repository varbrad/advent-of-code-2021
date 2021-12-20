package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varbrad/advent-of-code-2021/utils"
)

func TestDay17Part1(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
			"",
			"#..#.",
			"#....",
			"##..#",
			"..#..",
			"..###",
		}

		expected := 35
		actual := Day20Part1(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day20/input")

		expected := 4917
		actual := Day20Part1(input)

		assert.Equal(t, expected, actual)
	})
}

func TestDay17Part2(t *testing.T) {
	t.Run("Should solve the example", func(t *testing.T) {
		input := []string{
			"..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
			"",
			"#..#.",
			"#....",
			"##..#",
			"..#..",
			"..###",
		}

		expected := 3351
		actual := Day20Part2(input)

		assert.Equal(t, expected, actual)
	})

	t.Run("Should solve my puzzle input", func(t *testing.T) {
		input := utils.ReadInputToList("day20/input")

		expected := 16389
		actual := Day20Part2(input)

		assert.Equal(t, expected, actual)
	})
}
