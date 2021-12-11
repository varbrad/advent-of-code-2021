package main

import (
	"sort"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToIntGrid("day9/input")
	utils.Day(9, "Smoke Basin").Input(input).Part1(Day9Part1).Part2(Day9Part2)
}

func Day9Part1(input [][]int) int {
	grid := MakeGrid(input)
	mins := grid.GetMinimums()

	return utils.SumIter(func(ix int) int {
		min := mins[ix]
		return grid.Value(min.x, min.y)
	}, len(mins)) + len(mins)
}

func Day9Part2(input [][]int) int {
	grid := MakeGrid(input)
	mins := grid.GetMinimums()

	basins := []int{}
	for _, min := range mins {
		basins = append(basins, grid.CalculateBasin(min.x, min.y))
	}
	sort.Ints(basins)

	return utils.ProductInts(basins[len(basins)-3:])
}
