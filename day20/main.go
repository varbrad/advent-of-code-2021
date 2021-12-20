package main

import (
	"math"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day20/input")
	utils.Day(20, "Trench Map").Input(input).Part1(Day20Part1).Part2(Day20Part2)
}

func Day20Part1(input []string) int {
	return solve(input, 2)
}

func Day20Part2(input []string) int {
	return solve(input, 50)
}

func solve(input []string, loops int) int {
	algo := input[0]
	pixels := make(map[point]bool)
	for x, line := range input[2:] {
		for y, v := range line {
			pixels[point{x, y}] = v == '#'
		}
	}
	for i := 0; i < loops; i++ {
		newPixels := make(map[point]bool)
		b := getBounds(pixels)
		for x := b.minX - 1; x <= b.maxX+1; x++ {
			for y := b.minY - 1; y <= b.maxY+1; y++ {
				binary := 0
				for _, dx := range []int{-1, 0, 1} {
					for _, dy := range []int{-1, 0, 1} {
						binary = binary << 1
						if v, ok := pixels[point{x + dx, y + dy}]; ok {
							if v {
								binary |= 1
							}
						} else if algo[0] == '#' && i%2 != 0 {
							binary |= 1
						}
					}
				}
				newPixels[point{x, y}] = algo[binary] == '#'
			}
		}
		pixels = newPixels
	}
	return countLitPixels(pixels)
}

func getBounds(pixels map[point]bool) bounds {
	bounds := bounds{math.MaxInt, math.MinInt, math.MaxInt, math.MinInt}
	for p := range pixels {
		bounds.minX = utils.MinInteger(bounds.minX, p.x)
		bounds.maxX = utils.MaxInteger(bounds.maxX, p.x)
		bounds.minY = utils.MinInteger(bounds.minY, p.y)
		bounds.maxY = utils.MaxInteger(bounds.maxY, p.y)
	}
	return bounds
}

func countLitPixels(pixels map[point]bool) int {
	sum := 0
	for _, v := range pixels {
		if v {
			sum++
		}
	}
	return sum
}

type point struct{ x, y int }
type bounds struct{ minX, maxX, minY, maxY int }
