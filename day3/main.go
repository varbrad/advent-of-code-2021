package main

import (
	"strconv"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day3/input")
	utils.Day(3).Input(input).Part1(Day3Part1)
}

type Sums struct {
	zero int
	one  int
}

func Day3Part1(input []string) int {
	common := make([]Sums, len(input[0]))
	for _, s := range input {
		for ix, c := range s {
			switch string(c) {
			case "0":
				common[ix].zero = common[ix].zero + 1
			case "1":
				common[ix].one = common[ix].one + 1
			}
		}
	}

	a := ""
	b := ""
	for _, s := range common {
		if s.zero > s.one {
			a = a + "0"
			b = b + "1"
		} else {
			a = a + "1"
			b = b + "0"
		}
	}

	gamma, _ := strconv.ParseInt(a, 2, 64)
	epsilon, _ := strconv.ParseInt(b, 2, 64)
	return int(gamma * epsilon)
}
