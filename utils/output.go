package utils

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func genericCall(f interface{}, arg interface{}) interface{} {
	switch f.(type) {
	case func([]string) int:
		return f.(func([]string) int)(arg.([]string))
	case func(string) int:
		return f.(func(string) int)(arg.(string))
	case func([]string) string:
		return f.(func([]string) string)(arg.([]string))
	case func(string) string:
		return f.(func(string) string)(arg.(string))
	case func([]int) int:
		return f.(func([]int) int)(arg.([]int))
	case func(int) int:
		return f.(func(int) int)(arg.(int))
	case func([]int) string:
		return f.(func([]int) string)(arg.([]int))
	case func(int) string:
		return f.(func(int) string)(arg.(int))
	default:
		return nil
	}
}

func timeRun(part string, fn interface{}, a interface{}) {
	start := time.Now()
	result := genericCall(fn, a)
	elapsed := time.Since(start)
	log(part, result, "("+elapsed.String()+")")
}

type AoC struct {
	day   int
	input interface{}
}

func Day(day int) *AoC {
	color.Set(color.FgMagenta)

	fmt.Println("❄️  🎄 AoC 2021 - Day", day, "🎄 ❄️")

	color.Unset()
	return &AoC{day, nil}
}

func (aoc *AoC) Input(input interface{}) *AoC {
	aoc.input = input
	return aoc
}

func (aoc *AoC) Part1(fn interface{}) *AoC {
	timeRun("Part 1", fn, aoc.input)
	return aoc
}

func (aoc *AoC) Part2(fn interface{}) *AoC {
	timeRun("Part 2", fn, aoc.input)
	return aoc
}

func log(prefix string, result interface{}, timing string) {
	color.Set(color.FgBlue)
	fmt.Print(prefix, " ")
	color.Set(color.FgRed)
	fmt.Printf("%v ", result)
	color.Set(color.FgYellow)
	fmt.Print(timing, "\n")
	color.Unset()
}
