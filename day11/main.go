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
	n := 0
	for {
		n++
		og.step()
		ok := true
		for _, row := range og.grid {
			for _, v := range row {
				if v != 0 {
					ok = false
					break
				}
			}
		}
		if ok {
			return n
		}
	}
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
	flashing := []point{}

	for y := 0; y < len(og.grid); y++ {
		for x := 0; x < len(og.grid[y]); x++ {
			og.grid[y][x] += 1
			if og.grid[y][x] == 10 {
				flashing = append(flashing, point{x, y})
			}
		}
	}

	totalFlashed := og.handleFlashing(flashing)

	for y := 0; y < len(og.grid); y++ {
		for x := 0; x < len(og.grid[y]); x++ {
			if og.grid[y][x] > 9 {
				og.grid[y][x] = 0
			}
		}
	}

	return totalFlashed
}

func (og *octoGrid) handleFlashing(flashing []point) int {
	totalFlashing := len(flashing)
	for {
		if len(flashing) == 0 {
			return totalFlashing
		}
		p := flashing[0]
		flashing = flashing[1:]
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if x == 0 && y == 0 {
					continue
				}
				gy := p.y + y
				gx := p.x + x
				if gx < 0 || gy < 0 || gy > len(og.grid)-1 || gx > len(og.grid[gy])-1 {
					continue
				}
				og.grid[p.y+y][p.x+x]++
				if og.grid[p.y+y][p.x+x] == 10 {
					flashing = append(flashing, point{p.x + x, p.y + y})
					totalFlashing++
				}
			}
		}
	}
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
