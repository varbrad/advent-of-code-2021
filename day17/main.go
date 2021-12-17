package main

import (
	"math"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := []int{138, 184, -125, -71}
	utils.Day(17, "Trick Shot").Input(input).Part1(Day17Part1).Part2(Day17Part2)
}

func Day17Part1(input []int) int {
	return findBestTrickShot(input)
}

func Day17Part2(input []int) int {
	xs := getValidXHits(input)
	ys := getValidYHits(input)
	total := 0
	for _, y := range ys {
		for _, x := range xs {
			ok := fullTrickShot(x, y, input)
			if ok {
				total++
			}
		}
	}
	return total
}

func getValidXHits(input []int) []int {
	validX := []int{}
	s := input[0:2]
	for i := -250; i < 250; i++ {
		t := xShot(i, s)
		if t == "hit" {
			validX = append(validX, i)
		}
	}
	return validX
}

func getValidYHits(input []int) []int {
	validY := []int{}
	s := input[2:4]
	for i := -250; i < 250; i++ {
		t, _ := trickShot(i, s)
		if t == "hit" {
			validY = append(validY, i)
		}
	}
	return validY
}

func findBestTrickShot(input []int) int {
	bestHit := -1
	for i := 0; i < 250; i++ {
		t, r := trickShot(i, input[2:4])
		if t == "hit" {
			if r > bestHit {
				bestHit = r
			}
		}
	}
	return bestHit
}

func trickShot(dy int, target []int) (string, int) {
	y := 0
	maxY := math.MinInt
	for {
		y += dy
		if y > maxY {
			maxY = y
		}
		dy -= 1
		if y <= target[1] && y >= target[0] {
			return "hit", maxY
		} else if y < target[0] {
			return "miss", maxY
		}
	}
}

func xShot(dx int, target []int) string {
	x := 0
	for {
		x += dx
		if dx > 0 {
			dx--
		} else if dx < 0 {
			dx++
		}
		if x <= target[1] && x >= target[0] {
			return "hit"
		}
		if dx == 0 {
			return "miss"
		}
	}
}

func fullTrickShot(dx int, dy int, input []int) bool {
	x := 0
	y := 0
	for {
		x += dx
		y += dy
		if dx > 0 {
			dx--
		} else if dx < 0 {
			dx++
		}
		dy -= 1
		if x <= input[1] && x >= input[0] && y <= input[3] && y >= input[2] {
			return true
		}
		if y < input[2] || (x > input[1] && x < input[0] && dx == 0) {
			return false
		}
	}
}
