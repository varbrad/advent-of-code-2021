package main

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToIntegerList("day1/input")
	utils.Day(1).Part1(Day1Part1, input).Part2(Day1Part2, input)
}

func Day1Part1(input []int) int {
	return calculateIncreases(input)
}

func Day1Part2(input []int) int {
	l := len(input) - 2
	reducedList := make([]int, l)
	for i := 0; i < l; i++ {
		reducedList[i] = utils.SumInts(input[i : i+3])
	}
	return calculateIncreases(reducedList)
}

func calculateIncreases(numbers []int) int {
	inc := 0

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] < numbers[i+1] {
			inc += 1
		}
	}

	return inc
}
