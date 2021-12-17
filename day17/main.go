package main

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := []int{138, 184, -125, -71}
	utils.Day(17, "Trick Shot").Input(input).Part1(Day17Part1).Part2(Day17Part2)
}

func Day17Part1(input []int) int {
	_, maxY := calculateY(0, input[2]*-1-1, 0, input[2:4])
	return maxY
}

func Day17Part2(input []int) int {
	sum := 0
	maxX := input[1] + 1
	maxY := input[2] * -1
	for dy := -maxY; dy < maxY; dy++ {
		for dx := 0; dx < maxX; dx++ {
			if calculate2d(0, 0, dx, dy, input) {
				sum++
			}
		}
	}
	return sum
}

func calculateY(y, dy, maxY int, input []int) (bool, int) {
	switch {
	case y < input[0]:
		return false, 0
	case y <= input[1]:
		return true, maxY
	case y+dy > maxY:
		maxY = y + dy
	}
	return calculateY(y+dy, dy-1, maxY, input)
}

func calculate2d(x, y, dx, dy int, input []int) bool {
	switch {
	case x > input[1] || y < input[2]:
		return false
	case x >= input[0] && y <= input[3]:
		return true
	}
	return calculate2d(x+dx, y+dy, newDx(dx), dy-1, input)
}

func newDx(dx int) int {
	if dx == 0 {
		return 0
	}
	return dx - 1
}
