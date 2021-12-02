package main

import (
	"log"
	"strings"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input, err := utils.ReadInputToList("day2/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Day(2).Input(input).Part1(Day2Part1).Part2(Day2Part2)
}

func Day2Part1(input []string) int {
	return calculate(input, false)
}

func Day2Part2(input []string) int {
	return calculate(input, true)
}

func calculate(input []string, useAim bool) int {
	aim, depth, horizontal := 0, 0, 0
	for _, line := range input {
		parts := strings.Split(line, " ")
		operation, value := parts[0], utils.ToInt(parts[1])
		switch operation {
		case "forward":
			horizontal += value
			if useAim {
				depth += aim * value
			}
		case "down", "up":
			if operation == "up" {
				value *= -1
			}
			if useAim {
				aim += value
			} else {
				depth += value
			}
		}
	}
	return depth * horizontal
}
