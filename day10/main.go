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
		sum += parseSymbols([]rune(line), []rune{})
	}
	return sum
}

func Day10Part2(input []string) int {
	scores := []int{}
	for _, line := range input {
		incomplete, stack := getIncompleteStack([]rune(line), []rune{})
		if !incomplete {
			continue
		}
		scores = append(scores, calculateScore(stack, 0))
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func calculateScore(stack []rune, sum int) int {
	if len(stack) == 0 {
		return sum
	}
	sum *= 5
	switch stack[len(stack)-1] {
	case '(':
		sum += 1
	case '[':
		sum += 2
	case '{':
		sum += 3
	case '<':
		sum += 4
	}
	return calculateScore(stack[:len(stack)-1], sum)
}

func parseSymbols(input []rune, stack []rune) int {
	if len(input) == 0 {
		return 0
	}
	current := input[0]
	switch current {
	case '(', '{', '[', '<':
		stack = append(stack, current)
		return parseSymbols(input[1:], stack)
	case ')', '}', ']', '>':
		lastStack := stack[len(stack)-1]
		if lastStack == '(' && current == ')' ||
			lastStack == '{' && current == '}' ||
			lastStack == '[' && current == ']' ||
			lastStack == '<' && current == '>' {
			return parseSymbols(input[1:], stack[:len(stack)-1])
		} else {
			switch current {
			case ')':
				return 3
			case ']':
				return 57
			case '}':
				return 1197
			case '>':
				return 25137
			}
		}
	}
	return parseSymbols(input[1:], stack)
}

func getIncompleteStack(input []rune, stack []rune) (incomplete bool, finalStack []rune) {
	if len(input) == 0 {
		return true, stack
	}
	current := input[0]
	switch current {
	case '(', '{', '[', '<':
		return getIncompleteStack(input[1:], append(stack, current))
	case ')', '}', ']', '>':
		lastStack := stack[len(stack)-1]
		if lastStack == '(' && current == ')' ||
			lastStack == '{' && current == '}' ||
			lastStack == '[' && current == ']' ||
			lastStack == '<' && current == '>' {
			return getIncompleteStack(input[1:], stack[:len(stack)-1])
		} else {
			return false, nil
		}
	}
	return false, nil
}
