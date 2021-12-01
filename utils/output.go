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

func Part1(fn interface{}, a interface{}) {
	timeRun("Part 1", fn, a)
}

func Part2(fn interface{}, a interface{}) {
	timeRun("Part 2", fn, a)
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
