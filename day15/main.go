package main

import (
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
	start := point{0, 0}
	end := point{len(grid.grid[0]) - 1, len(grid.grid) - 1}

	open := openList{}
	closed := make(map[point]*node)

	open.addNode(&node{id: start, value: grid.get(start), cost: 0})

	var current *node
	for open.len() > 0 {
		current = open.next()

		closed[current.id] = current
		open.remove(current)

		if current.id == end {
			break
		}

		neighbours := grid.neighbours(current.id)
		for _, p := range neighbours {
			cost := current.cost + grid.get(p)
			isOpen := open.isOpen(p)
			_, isClosed := closed[p]
			if !isOpen && !isClosed {
				open.addNode(&node{id: p, value: grid.get(p), cost: cost, previous: current})
			} else if isOpen {
				open.compare(p, cost, current)
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

type openList struct {
	list []*node
}

func (ol *openList) addNode(n *node) {
	l := len(ol.list)
	for i := 0; i < l; i++ {
		if ol.list[i].cost > n.cost {
			ol.list = append(ol.list[:i], append([]*node{n}, ol.list[i:]...)...)
			return
		}
	}
	ol.list = append(ol.list, n)
}

func (ol *openList) len() int {
	return len(ol.list)
}

func (ol *openList) next() *node {
	return ol.list[0]
}

func (ol *openList) compare(p point, cost int, current *node) {
	for _, n := range ol.list {
		if n.id != p {
			continue
		}
		if n.cost > cost {
			n.cost = cost
			n.previous = current
			ol.remove(n)
			ol.addNode(n)
		}
	}
}

func (ol *openList) remove(node *node) {
	for i, n := range ol.list {
		if n == node {
			ol.list = append(ol.list[:i], ol.list[i+1:]...)
			return
		}
	}
}

func (ol *openList) isOpen(p point) bool {
	for _, n2 := range ol.list {
		if n2.id == p {
			return true
		}
	}
	return false
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
	cost     int
	previous *node
}
