package main

import (
	"regexp"
	"strings"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day4/input")
	utils.Day(4).Input(input).Part1(Day4Part1).Part2(Day4Part2)
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

var spaceRegex = regexp.MustCompile(`(\s+)`)

func parseInput(input []string) ([]int, []*BingoBoard) {
	numbers := []int{}
	for _, v := range strings.Split(input[0], ",") {
		numbers = append(numbers, utils.ToInteger(v))
	}

	boards := []*BingoBoard{}
	for i := 2; i < len(input); i++ {
		board := &BingoBoard{}
		for j := 0; j < 5; j++ {
			row := &BingoRow{}
			rowValues := spaceRegex.Split(strings.TrimSpace(input[i+j]), -1)
			for k := 0; k < 5; k++ {
				cell := &BingoCell{value: utils.ToInteger(rowValues[k])}
				row[k] = cell
			}
			board[j] = row
		}
		boards = append(boards, board)

		i += 5
	}
	return numbers, boards
}

type BingoCell struct {
	value  int
	marked bool
}
type BingoRow [5]*BingoCell
type BingoBoard [5]*BingoRow

func (bb *BingoBoard) Mark(value int) {
	cell := bb.FindCellByValue(value)
	if cell == nil {
		return
	}
	cell.marked = true
}

func (bb *BingoBoard) IsBingo() bool {
	for i := 0; i < 5; i++ {
		xCheck := bb[i][0].marked && bb[i][1].marked && bb[i][2].marked && bb[i][3].marked && bb[i][4].marked
		if xCheck {
			return true
		}
		yCheck := bb[0][i].marked && bb[1][i].marked && bb[2][i].marked && bb[3][i].marked && bb[4][i].marked
		if yCheck {
			return true
		}
	}
	return false
}

func (bb *BingoBoard) FindCellByValue(value int) *BingoCell {
	for y := 0; y < len(bb); y++ {
		row := bb[y]
		for x := 0; x < len(row); x++ {
			cell := row[x]
			if cell.value == value {
				return cell
			}
		}
	}
	return nil
}

func (bb *BingoBoard) SumOfUnmarked() int {
	sum := 0
	for y := 0; y < len(bb); y++ {
		row := bb[y]
		for x := 0; x < len(row); x++ {
			cell := row[x]
			if !cell.marked {
				sum += cell.value
			}
		}
	}
	return sum
}
