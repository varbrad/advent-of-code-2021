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
	case func([][]int) int:
		return f.(func([][]int) int)(arg.([][]int))
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

type AoC struct {
	day   int
	title string
	input interface{}
}

func Day(day int, title string) *AoC {
	return &AoC{day, title, nil}
}

func (aoc *AoC) Input(input interface{}) *AoC {
	aoc.input = input
	return aoc
}

func (aoc *AoC) Part1(fn interface{}) *AoC {
	aoc.timeRun("Part 1", fn, aoc.input)
	return aoc
}

func (aoc *AoC) Part2(fn interface{}) *AoC {
	aoc.timeRun("Part 2", fn, aoc.input)
	return aoc
}

func (aoc *AoC) timeRun(part string, fn interface{}, a interface{}) {
	start := time.Now()
	result := genericCall(fn, a)
	elapsed := time.Since(start)
	aoc.log(part, result, "("+fmt.Sprint(float64(elapsed.Microseconds())/1000.0)+"ms)")
}

func (aoc *AoC) log(prefix string, result interface{}, timing string) {
	color.Set(color.FgMagenta)
	fmt.Printf("Day " + fmt.Sprintf("%2v", aoc.day) + ": ")
	fmt.Printf("%-26v", aoc.title)
	color.Set(color.FgBlue)
	fmt.Print(prefix, " ")
	color.Set(color.FgRed)
	fmt.Printf("%20v ", result)
	color.Set(color.FgYellow)
	fmt.Printf("%14v", timing+"\n")
	color.Unset()

}
