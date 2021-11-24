package utils

import "fmt"

func Part1(a ...interface{}) {
	log("Part 1 =", a...)
}

func Part2(a ...interface{}) {
	log("Part 2 =", a...)
}

func log(prefix string, a ...interface{}) {
	args := append([]interface{}{prefix}, a...)
	fmt.Println(args...)
}
