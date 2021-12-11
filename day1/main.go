package main

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToIntegerList("day1/input")
	utils.Day(1, "Sonar Sweep").Input(input).Part1(Day1Part1).Part2(Day1Part2)
}

func Day1Part1(input []int) int {
	return calculateIncreases(input, 1)
}

func Day1Part2(input []int) int {
	return calculateIncreases(input, 3)
}

func calculateIncreases(numbers []int, gap int) int {
	inc := 0
	for i := 0; i < len(numbers)-gap; i++ {
		if numbers[i] < numbers[i+gap] {
			inc += 1
		}
	}
	return inc
}
