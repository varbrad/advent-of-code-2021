package main

import (
	"strconv"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day3/input")
	utils.Day(3, "Binary Diagnostic").Input(input).Part1(Day3Part1).Part2(Day3Part2)
}

type Sums struct {
	zero int
	one  int
}

func Day3Part1(input []string) int {
	commons := getCommons(input)
	gamma := getString("common", commons)
	epsilon := getString("uncommon", commons)
	return calculate(gamma, epsilon)
}

func Day3Part2(input []string) int {
	oxygen := recursiveSearcher("common", input, 0)
	co2 := recursiveSearcher("uncommon", input, 0)
	return calculate(oxygen, co2)
}

func calculate(a string, b string) int {
	gamma, _ := strconv.ParseInt(a, 2, 64)
	epsilon, _ := strconv.ParseInt(b, 2, 64)

	return int(gamma * epsilon)
}

func getCommons(input []string) []Sums {
	commons := make([]Sums, len(input[0]))
	for _, s := range input {
		for ix, rune := range s {
			char := string(rune)
			if char == "0" {
				commons[ix].zero += 1
			} else {
				commons[ix].one += 1
			}
		}
	}
	return commons
}

func recursiveSearcher(mode string, input []string, index int) string {
	commons := getCommons(input)
	v := getString(mode, commons)

	left := []string{}
	for _, s := range input {
		if s[index] == v[index] {
			left = append(left, s)
		}
	}
	if len(left) == 1 {
		return left[0]
	}
	return recursiveSearcher(mode, left, index+1)
}

func getString(mode string, commons []Sums) string {
	s := ""
	for _, v := range commons {
		if mode == "common" {
			if v.one >= v.zero {
				s += "1"
			} else {
				s += "0"
			}
		} else {
			if v.zero > v.one {
				s += "1"
			} else {
				s += "0"
			}
		}
	}
	return s
}
