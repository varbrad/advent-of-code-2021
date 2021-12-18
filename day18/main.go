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
	maximumSum := -1
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			a := parseSnailfishNumber(input[i]).Add(parseSnailfishNumber(input[j])).Reduce().Magnitude()
			b := parseSnailfishNumber(input[j]).Add(parseSnailfishNumber(input[i])).Reduce().Magnitude()

			maximumSum = utils.MaxInteger(utils.MaxInteger(a, b), maximumSum)
		}
	}

	return maximumSum
}
