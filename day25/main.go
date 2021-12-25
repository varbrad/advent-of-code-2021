package main

// ...>>>>>...

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day25/input")
	utils.Day(25, "Sea Cucumber").Input(input).Part1(Day25Part1)
}

func Day25Part1(input []string) int {
	cucumbers := parseCucumbers(input)

	n := 0
	for {
		n++
		moves := cucumbers.step()
		if moves == 0 {
			break
		}
	}

	return n
}

type point struct {
	x int
	y int
}

type cucumbers struct {
	east   map[point]bool
	south  map[point]bool
	width  int
	height int
}

func parseCucumbers(input []string) *cucumbers {
	c := cucumbers{
		east:   make(map[point]bool),
		south:  make(map[point]bool),
		height: len(input),
		width:  len(input[0]),
	}

	for y, row := range input {
		for x, rune := range row {
			if rune == '>' {
				c.east[point{x, y}] = true
			} else if rune == 'v' {
				c.south[point{x, y}] = true
			}
		}
	}

	return &c
}

func (c *cucumbers) step() int {
	eastMovers := []point{}
	for p := range c.east {
		if c.isFree(p.x+1, p.y) {
			eastMovers = append(eastMovers, p)
		}
	}

	for _, p := range eastMovers {
		delete(c.east, p)
		c.east[c.normalisedPoint(p.x+1, p.y)] = true
	}

	southMovers := []point{}
	for p := range c.south {
		if c.isFree(p.x, p.y+1) {
			southMovers = append(southMovers, p)
		}
	}

	for _, p := range southMovers {
		delete(c.south, p)
		c.south[c.normalisedPoint(p.x, p.y+1)] = true
	}

	return len(eastMovers) + len(southMovers)
}

func (c *cucumbers) isFree(x, y int) bool {
	p := c.normalisedPoint(x, y)
	if _, takenEast := c.east[p]; takenEast {
		return false
	}
	if _, takenSouth := c.south[p]; takenSouth {
		return false
	}
	return true
}

func (c *cucumbers) normalisedPoint(x, y int) point {
	return point{x: x % c.width, y: y % c.height}
}

func (c *cucumbers) getXY(x, y int) string {
	p := c.normalisedPoint(x, y)
	if _, takenEast := c.east[p]; takenEast {
		return ">"
	}
	if _, takenSouth := c.south[p]; takenSouth {
		return "v"
	}
	return "."
}

func (c *cucumbers) String() string {
	s := ""
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			s += c.getXY(x, y)
		}
		s += "\n"
	}
	return s
}
