package main

import (
	"log"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input, err := utils.ReadInputToIntegerList("day1/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Day(1).Part1(Day1Part1, input).Part2(Day1Part2, input)
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
