package main

import (
	"strconv"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day3/input")
	utils.Day(3).Input(input).Part1(Day3Part1).Part2(Day3Part2)
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

func Day3Part2(input []string) int {
	oxygen := getMatching("common", input, 0)
	co2 := getMatching("uncommon", input, 0)

	gamma, _ := strconv.ParseInt(oxygen, 2, 64)
	epsilon, _ := strconv.ParseInt(co2, 2, 64)

	return int(gamma * epsilon)
}

func getMatching(mode string, input []string, index int) string {
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
	return getMatching(mode, left, index+1)
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
