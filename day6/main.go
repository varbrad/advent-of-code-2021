package main

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadCommaInputToIntegerList("day6/input")
	utils.Day(6).Input(input).Part1(Day6Part1).Part2(Day6Part2)
}

func Day6Part1(input []int) int {
	result := parseInput(input)

	for i := 0; i < 80; i++ {
		result = process(result)
	}
	return utils.SumInts(result[:])
}

func Day6Part2(input []int) int {
	result := parseInput(input)

	for i := 0; i < 256; i++ {
		result = process(result)
	}
	return utils.SumInts(result[:])
}

func parseInput(input []int) [9]int {
	list := [9]int{}
	for _, v := range input {
		list[v]++
	}
	return list
}

func process(input [9]int) [9]int {
	list := [9]int{}
	zeroes := input[0]

	for i := 0; i < 8; i++ {
		list[i] = input[i+1]
	}

	list[6] += zeroes
	list[8] = zeroes

	return list
}
