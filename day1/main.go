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

	utils.Part1(Day1Part1(input))
	utils.Part2(Day1Part2(input))
}

func Day1Part1(input []int) int {
	return calculateIncreases(input)
}

func Day1Part2(input []int) int {
	l := len(input) - 2
	reducedList := make([]int, l)
	for i := 0; i < l; i++ {
		reducedList[i] = utils.SumInts(input[i : i+2])
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
