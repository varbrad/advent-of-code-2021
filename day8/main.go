package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day8/input")
	utils.Day(8, "Seven Segment Search").Input(input).Part1(Day8Part1).Part2(Day8Part2)
}

func Day8Part1(input []string) int {
	displays := parseInput(input)
	sum := 0
	for _, display := range displays {
		for _, digit := range display.digits {
			switch len(digit) {
			case 2, 3, 4, 7:
				sum++
			}
		}
	}
	return sum
}

func Day8Part2(input []string) int {
	displays := parseInput(input)
	sum := 0
	for _, display := range displays {
		sum += display.Solve()
	}
	return sum
}

type Display struct {
	signals      []string
	digits       []string
	solvedDigits map[int]string
}

// 1. Determine 1 [2-seg], 4 [4-seg], 7 [3-seg], 8 [7-seg]
// 2. Determine 3 (only 5-segment number that contains both of 1's segments)
// 3. Determine 6 (only 6-segment number that doesnt contain both of 1's segments)
// 4. Determine 9 (only 6-segment number that contains all of 4's segments)
// 5. 0 is only 6-segment number left
// 6. Determine 5 by finding remaining digit with only 1 diff from 6
// 7. 2 Is the last digit remaining
func (d *Display) Solve() int {
	d.solvedDigits = map[int]string{}
	d.Solve1478()
	d.Solve369()
	d.Solve05()

	str := ""
	for _, v := range d.digits {
		str += fmt.Sprint(d.SignalToInt(v))
	}

	return utils.ToInteger(str)
}

func (d *Display) Solve1478() {
	for i := len(d.signals) - 1; i >= 0; i-- {
		signal := d.signals[i]
		switch len(signal) {
		case 2:
			d.solvedDigits[1] = signal
			d.RemoveSignal(i)
		case 3:
			d.solvedDigits[7] = signal
			d.RemoveSignal(i)
		case 4:
			d.solvedDigits[4] = signal
			d.RemoveSignal(i)
		case 7:
			d.solvedDigits[8] = signal
			d.RemoveSignal(i)
		}
	}
}

func (d *Display) Solve369() {
	for i := len(d.signals) - 1; i >= 0; i-- {
		signal := d.signals[i]
		if len(signal) == 5 && utils.StringContainsLetters(signal, d.solvedDigits[1]) {
			d.solvedDigits[3] = signal
			d.RemoveSignal(i)
		} else if len(signal) == 6 && !utils.StringContainsLetters(signal, d.solvedDigits[1]) {
			d.solvedDigits[6] = signal
			d.RemoveSignal(i)
		} else if len(signal) == 6 && utils.StringContainsLetters(signal, d.solvedDigits[4]) {
			d.solvedDigits[9] = signal
			d.RemoveSignal(i)
		}
	}
}

func (d *Display) Solve05() {
	for i := len(d.signals) - 1; i >= 0; i-- {
		signal := d.signals[i]
		if len(signal) == 6 {
			d.solvedDigits[0] = signal
			d.RemoveSignal(i)
		} else if len(signal) == 5 && utils.CountStringDiff(signal, d.solvedDigits[6]) == 1 {
			d.solvedDigits[5] = signal
			d.RemoveSignal(i)
		} else {
			d.solvedDigits[2] = signal
			d.RemoveSignal(i)
		}
	}
}

func (d *Display) RemoveSignal(signalIndex int) {
	d.signals = append(d.signals[:signalIndex], d.signals[signalIndex+1:]...)
}

func (d *Display) SignalToInt(signal string) int {
	for key, v := range d.solvedDigits {
		if v == signal {
			return key
		}
	}
	return -1
}

func (d *Display) Sort() {
	for ix, digit := range d.digits {
		d.digits[ix] = utils.SortString(digit)
	}
	for ix, signal := range d.signals {
		d.signals[ix] = utils.SortString(signal)
	}
}

var inputRegex = regexp.MustCompile(`^((?:\w+\s?)+) \| ((?:\w+\s?)+)$`)

func parseInput(input []string) []Display {
	displays := []Display{}
	for _, line := range input {
		matches := inputRegex.FindStringSubmatch(line)
		display := Display{
			signals: strings.Split(matches[1], " "),
			digits:  strings.Split(matches[2], " "),
		}
		display.Sort()
		displays = append(displays, display)
	}
	return displays
}
