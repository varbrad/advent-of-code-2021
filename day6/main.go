package main

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadCommaInputToIntegerList("day6/input")
	utils.Day(6, "Lanternfish").Input(input).Part1(Day6Part1).Part2(Day6Part2)
}

func Day6Part1(input []int) int {
	return calculate(input, 80)
}

func Day6Part2(input []int) int {
	return calculate(input, 256)
}

func calculate(input []int, loops int) int {
	result := parseInput(input)
	for i := 0; i < loops; i++ {
		result = process(result)
	}
	return utils.SumInts(result)
}

func parseInput(input []int) []int {
	list := make([]int, 9)
	for _, v := range input {
		list[v]++
	}
	return list
}

func process(input []int) []int {
	list := append(input[1:], input[0])
	list[6] += input[0]
	return list
}
