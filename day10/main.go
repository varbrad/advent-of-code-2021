package main

import (
	"sort"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day10/input")
	utils.Day(10).Input(input).Part1(Day10Part1).Part2(Day10Part2)
}

func Day10Part1(input []string) int {
	sum := 0
	for _, line := range input {
		pl := createLine(line)
		sum += pl.corruptionScore()
	}
	return sum
}

func Day10Part2(input []string) int {
	scores := []int{}
	for _, line := range input {
		pl := createLine(line)
		if pl.status != "incomplete" {
			continue
		}
		scores = append(scores, pl.incompleteScore())
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

type parsedLine struct {
	status   string
	stack    []rune
	lastRune rune
}

func createLine(input string) parsedLine {
	return parseLine([]rune(input), []rune{})
}

func parseLine(runes []rune, stack []rune) parsedLine {
	if len(runes) == 0 {
		return parsedLine{"incomplete", stack, 0}
	}
	current := runes[0]
	switch current {
	case '(', '{', '[', '<':
		stack = append(stack, current)
		return parseLine(runes[1:], stack)
	case ')', '}', ']', '>':
		lastStack := stack[len(stack)-1]
		if doRunesClose(lastStack, current) {
			return parseLine(runes[1:], stack[:len(stack)-1])
		} else {
			return parsedLine{"corrupted", stack, current}
		}
	}
	panic("unknown character encountered")
}

func doRunesClose(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	case '<':
		return close == '>'
	}
	panic("unknown open rune")
}

func (pl parsedLine) corruptionScore() int {
	if pl.status != "corrupted" {
		return 0
	}
	switch pl.lastRune {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	panic("unknown final rune")
}

func (pl parsedLine) incompleteScore() int {
	if pl.status != "incomplete" {
		return 0
	}
	sum := 0
	for i := len(pl.stack) - 1; i >= 0; i-- {
		sum *= 5
		switch pl.stack[i] {
		case '(':
			sum += 1
		case '[':
			sum += 2
		case '{':
			sum += 3
		case '<':
			sum += 4
		}
	}
	return sum
}
