package main

import (
	"fmt"
	"regexp"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day5/input")
	utils.Day(5, "Hydrothermal Venture").Input(input).Part1(Day5Part1).Part2(Day5Part2)
}

func Day5Part1(input []string) int {
	return solver(input, false)
}

func Day5Part2(input []string) int {
	return solver(input, true)
}

func solver(input []string, includeDiagonals bool) int {
	lines := parseInput(input)

	cells := map[string]int{}

	for _, line := range lines {
		if !includeDiagonals && line.a.x != line.b.x && line.a.y != line.b.y {
			continue
		}
		x, y := line.a.x, line.a.y
		dx := getSign(line.a.x, line.b.x)
		dy := getSign(line.a.y, line.b.y)
		for {
			cells[fmt.Sprintf("%d,%d", x, y)]++
			if x == line.b.x && y == line.b.y {
				break
			}
			x += dx
			y += dy
		}
	}

	sum := 0
	for _, v := range cells {
		if v > 1 {
			sum++
		}
	}
	return sum
}

func getSign(a int, b int) int {
	switch {
	case a < b:
		return 1
	case b < a:
		return -1
	default:
		return 0
	}
}

var parseRegex = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

type Point struct {
	x int
	y int
}

type Line struct {
	a Point
	b Point
}

func parseInput(input []string) []Line {
	lines := make([]Line, len(input))
	for i, line := range input {
		matches := parseRegex.FindStringSubmatch(line)
		a := Point{utils.ToInteger(matches[1]), utils.ToInteger(matches[2])}
		b := Point{utils.ToInteger(matches[3]), utils.ToInteger(matches[4])}
		lines[i] = Line{a, b}
	}
	return lines
}
