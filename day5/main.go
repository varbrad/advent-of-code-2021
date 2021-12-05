package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day5/input")
	utils.Day(5).Input(input).Part1(Day5Part1).Part2(Day5Part2)
}

func Day5Part1(input []string) int {
	lines := parseInput(input)

	cells := map[string]int{}

	for _, line := range lines {
		if line.a.x != line.b.x && line.a.y != line.b.y {
			continue
		}
		if line.a.x == line.b.x {
			minY := int(math.Min(float64(line.a.y), float64(line.b.y)))
			maxY := int(math.Max(float64(line.a.y), float64(line.b.y)))

			for y := minY; y <= maxY; y++ {
				cells[fmt.Sprintf("%d,%d", line.a.x, y)]++
			}
		} else {
			minX := int(math.Min(float64(line.a.x), float64(line.b.x)))
			maxX := int(math.Max(float64(line.a.x), float64(line.b.x)))

			for x := minX; x <= maxX; x++ {
				cells[fmt.Sprintf("%d,%d", x, line.a.y)]++
			}
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

func Day5Part2(input []string) int {
	lines := parseInput(input)

	cells := map[string]int{}

	for _, line := range lines {
		x, y := line.a.x, line.a.y
		max := 999
		for {
			cells[fmt.Sprintf("%d,%d", x, y)]++
			max--
			if x == line.b.x && y == line.b.y {
				break
			}
			if line.a.x < line.b.x {
				x++
			} else if line.a.x > line.b.x {
				x--
			}
			if line.a.y < line.b.y {
				y++
			} else if line.a.y > line.b.y {
				y--
			}
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
