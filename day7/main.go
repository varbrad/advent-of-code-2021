package main

import (
	"math"
	"sort"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadCommaInputToIntegerList("day7/input")
	utils.Day(7).Input(input).Part1(Day7Part1).Part2(Day7Part2)
}

func Day7Part1(input []int) int {
	sort.Ints(input)
	median := input[len(input)/2]
	return calculateFuel(input, median, "add")
}

func Day7Part2(input []int) int {
	average := utils.AverageInts(input)
	return utils.FindMinInt(func(i int) int {
		return calculateFuel(input, average+i, "triangle")
	}, utils.MakeRange(-1, 1))
}

func calculateFuel(values []int, start int, mode string) int {
	fuel := 0
	for _, v := range values {
		diff := int(math.Abs(float64(v - start)))
		if mode == "add" {
			fuel += diff
		} else if mode == "triangle" {
			fuel += diff * (diff + 1) / 2
		}
	}
	return fuel
}
