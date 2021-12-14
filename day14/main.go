package main

import (
	"math"
	"strings"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day14/input")
	utils.Day(14, "Extended Polymerization").Input(input).Part1(Day14Part1).Part2(Day14Part2)
}

func Day14Part1(input []string) int {
	return solve(input, 10)
}

func Day14Part2(input []string) int {
	return solve(input, 40)
}

func solve(input []string, steps int) int {
	p := makePolymer(input)
	for i := 0; i < steps; i++ {
		p.step()
	}
	return p.diff()
}

type polymer struct {
	pairs  map[string]int
	lookup map[string]string
	last   byte
}

func (p *polymer) step() {
	newPairs := map[string]int{}
	for k, v := range p.pairs {
		newPairs[string(k[0])+p.lookup[k]] += v
		newPairs[p.lookup[k]+string(k[1])] += v
	}
	p.pairs = newPairs
}

func (p *polymer) diff() int {
	sums := map[byte]int{}
	for k, v := range p.pairs {
		sums[k[0]] += v
	}
	sums[p.last] += 1
	min, max := math.MaxInt, math.MinInt
	for _, v := range sums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func makePolymer(input []string) *polymer {
	p := polymer{lookup: make(map[string]string)}
	for i, line := range input {
		if len(line) == 0 {
			continue
		}
		if i == 0 {
			p.pairs = make(map[string]int)
			p.last = line[len(line)-1]
			for i = 0; i < len(line)-1; i++ {
				p.pairs[string(line[i])+string(line[i+1])]++
			}
			continue
		}
		parts := strings.Split(line, " -> ")
		p.lookup[parts[0]] = parts[1]
	}
	return &p
}
