package main

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day4/input")
	utils.Day(4, "Giant Squid").Input(input).Part1(Day4Part1).Part2(Day4Part2)
}

func Day4Part1(input []string) int {
	values, boards := parseInput(input)

	for _, v := range values {
		for _, b := range boards {
			b.Mark(v)
			if b.IsBingo() {
				return b.SumOfUnmarked() * v
			}
		}
	}

	return -1
}

func Day4Part2(input []string) int {
	values, boards := parseInput(input)

	for _, v := range values {
		for j := len(boards) - 1; j >= 0; j-- {
			b := boards[j]
			b.Mark(v)
			if b.IsBingo() {
				if len(boards) == 1 {
					return b.SumOfUnmarked() * v
				}
				boards = append(boards[:j], boards[j+1:]...)
			}
		}
	}

	return -1
}
