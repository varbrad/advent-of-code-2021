package main

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day18/input")
	utils.Day(18, "Snailfish").Input(input).Part1(Day18Part1).Part2(Day18Part2)
}

func Day18Part1(input []string) int {
	current := parseSnailfishNumber(input[0])
	for _, r := range input[1:] {
		next := parseSnailfishNumber(r)
		current = current.Add(next)
		current.Reduce()
	}

	return current.Magnitude()
}

func Day18Part2(input []string) int {
	max := -1
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			max = utils.MaxInteger(max, maximumMagnitude(input[i], input[j]))
		}
	}

	return max
}

func maximumMagnitude(a string, b string) int {
	return utils.MaxInteger(
		parseSnailfishNumber(a).Add(parseSnailfishNumber(b)).Reduce().Magnitude(),
		parseSnailfishNumber(b).Add(parseSnailfishNumber(a)).Reduce().Magnitude(),
	)
}
