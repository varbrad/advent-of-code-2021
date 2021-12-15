package main

import (
	"math"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToIntGrid("day15/input")
	utils.Day(15, "Chiton").Input(input).Part1(Day15Part1).Part2(Day15Part2)
}

func Day15Part1(input [][]int) int {
	return solve(input)
}

func Day15Part2(input [][]int) int {
	large := make([][]int, len(input)*5)
	for y, row := range input {
		large[y] = make([]int, len(row)*5)
		for i := 1; i < 5; i++ {
			large[y+len(row)*i] = make([]int, len(row)*5)
		}
		for x, val := range row {
			large[y][x] = val
			for i := 0; i < 5; i++ {
				for d := 0; d < 5; d++ {
					v := val + i + d
					for v > 9 {
						v = v - 9
					}
					large[y+len(row)*i][x+len(row)*d] = v
				}
			}
		}
	}

	return solve(large)
}

func solve(input [][]int) int {
	grid := makeGrid(input)

	open := make(map[point]*node)
	closed := make(map[point]*node)

	start := point{0, 0}
	end := point{len(grid.grid[0]) - 1, len(grid.grid) - 1}
	open[start] = &node{id: start, value: grid.get(start), f: grid.heuristic(start), g: grid.heuristic(start)}

	current := lowestFScore(open)
	for len(open) > 0 {
		current = lowestFScore(open)

		closed[current.id] = current
		delete(open, current.id)

		if current.id == end {
			break
		}

		neighbours := grid.neighbours(current.id)
		for _, p := range neighbours {
			g := current.g + grid.get(p)
			_, isOpen := open[p]
			_, isClosed := closed[p]
			if !isOpen && !isClosed {
				open[p] = &node{id: p, value: grid.get(p), f: g + grid.heuristic(p), g: g, previous: current}
			} else if isOpen {
				if open[p].g > g {
					open[p].g = g
					open[p].f = g + grid.heuristic(p)
					open[p].previous = current
				}
			}
		}
	}

	sum := 0
	for {
		if current.id == start {
			break
		}
		sum += current.value
		current = current.previous
	}

	return sum
}

func lowestFScore(open map[point]*node) *node {
	lowest := math.MaxInt64
	var lowestNode *node
	for _, n := range open {
		if n.f < lowest {
			lowest = n.f
			lowestNode = n
		}
	}
	return lowestNode
}

type grid struct {
	grid [][]int
}

func makeGrid(input [][]int) grid {
	return grid{grid: input}
}

func (g *grid) get(p point) int {
	return g.grid[p.y][p.x]
}

func (g *grid) heuristic(p point) int {
	maxX, maxY := len(g.grid[0]), len(g.grid)
	return int(math.Abs(float64(p.x-maxX))) + int(math.Abs(float64(p.y-maxY)))
}

func (g *grid) neighbours(p point) []point {
	neighbours := []point{}
	if p.x > 0 {
		neighbours = append(neighbours, point{p.x - 1, p.y})
	}
	if p.x < len(g.grid[0])-1 {
		neighbours = append(neighbours, point{p.x + 1, p.y})
	}
	if p.y > 0 {
		neighbours = append(neighbours, point{p.x, p.y - 1})
	}
	if p.y < len(g.grid)-1 {
		neighbours = append(neighbours, point{p.x, p.y + 1})
	}
	return neighbours
}

type point struct {
	x int
	y int
}

type node struct {
	id       point
	value    int
	f        int
	g        int
	previous *node
}
