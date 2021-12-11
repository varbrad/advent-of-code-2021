package main

import (
	"fmt"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToIntGrid("day11/input")
	utils.Day(11).Input(input).Part1(Day11Part1).Part2(Day11Part2)
}

func Day11Part1(input [][]int) int {
	og := makeOctoGrid(input)
	return utils.SumIter(func(ix int) int {
		return og.step()
	}, 100)
}

func Day11Part2(input [][]int) int {
	og := makeOctoGrid(input)
	for n := 0; n < 10_000; n++ {
		og.step()
		if og.allZeroes() {
			return n + 1
		}
	}
	return -1
}

type octoGrid struct {
	grid [][]int
}

type point struct {
	x int
	y int
}

func makeOctoGrid(grid [][]int) octoGrid {
	duplicate := make([][]int, len(grid))
	for i := range grid {
		duplicate[i] = make([]int, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return octoGrid{duplicate}
}

func (og *octoGrid) step() int {
	initalFlashed := og.initialFlashed()
	totalFlashed := og.calculateFlashed(initalFlashed)
	og.resetFlashed()
	return totalFlashed
}

func (og *octoGrid) initialFlashed() []point {
	flashing := []point{}
	for a := 0; a < len(og.grid)*len(og.grid[0]); a++ {
		x, y := a%len(og.grid), a/len(og.grid)
		og.grid[y][x]++
		if og.grid[y][x] == 10 {
			flashing = append(flashing, point{x, y})
		}
	}
	return flashing
}

func (og *octoGrid) calculateFlashed(flashing []point) int {
	for ix := 0; ix < len(flashing); ix++ {
		p := flashing[ix]
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x == 0 && y == 0 {
					continue
				}
				if og.incrementCell(p.x+x, p.y+y) == 10 {
					flashing = append(flashing, point{p.x + x, p.y + y})
				}
			}
		}
	}
	return len(flashing)
}

func (og *octoGrid) resetFlashed() {
	for a := 0; a < len(og.grid)*len(og.grid[0]); a++ {
		x, y := a%len(og.grid), a/len(og.grid)
		if og.grid[y][x] > 9 {
			og.grid[y][x] = 0
		}
	}
}

func (og *octoGrid) incrementCell(x, y int) int {
	if x < 0 || y < 0 || y > len(og.grid)-1 || x > len(og.grid[y])-1 {
		return -1
	}
	og.grid[y][x]++
	return og.grid[y][x]
}

func (og *octoGrid) allZeroes() bool {
	for _, row := range og.grid {
		for _, v := range row {
			if v != 0 {
				return false
			}
		}
	}
	return true
}

func (og octoGrid) String() string {
	str := ""
	for _, row := range og.grid {
		for _, val := range row {
			str += fmt.Sprintf("%d ", val)
		}
		str += "\n"
	}
	return str
}
