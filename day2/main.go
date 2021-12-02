package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input, err := utils.ReadInputToList("day2/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Day(2).Part1(Day2Part1, input).Part2(Day2Part2, input)
}

func Day2Part1(input []string) int {
	x, y, instructions := 0, 0, processInput(input)
	for _, ins := range instructions {
		switch ins.operation {
		case "forward":
			x += ins.value
		case "down":
			y += ins.value
		case "up":
			y -= ins.value
		}
	}
	return x * y
}

func Day2Part2(input []string) int {
	aim, depth, horizontal, instructions := 0, 0, 0, processInput(input)
	for _, ins := range instructions {
		switch ins.operation {
		case "forward":
			horizontal += ins.value
			depth += aim * ins.value
		case "down":
			aim += ins.value
		case "up":
			aim -= ins.value
		}
	}
	return depth * horizontal
}

type Instruction struct {
	operation string
	value     int
}

func processInput(input []string) []Instruction {
	instructions := make([]Instruction, len(input))
	for _, v := range input {
		a := strings.Split(v, " ")
		val, _ := strconv.Atoi(a[1])
		instructions = append(instructions, Instruction{a[0], val})
	}
	return instructions
}
