package main

import "fmt"

type Grid struct {
	grid    [][]int
	visited map[string]bool
}

type Point struct {
	x int
	y int
}

func MakeGrid(data [][]int) Grid {
	return Grid{data, make(map[string]bool)}
}

func (g *Grid) CalculateBasin(x int, y int) int {
	me := g.Value(x, y)
	g.Visit(x, y)
	return 1 +
		g.CalculateSubBasin(x-1, y, me) +
		g.CalculateSubBasin(x+1, y, me) +
		g.CalculateSubBasin(x, y-1, me) +
		g.CalculateSubBasin(x, y+1, me)
}

func (g *Grid) CalculateSubBasin(x, y, curr int) int {
	newVal := g.Value(x, y)
	if newVal == -1 || curr >= newVal || newVal == 9 {
		return 0
	}
	return g.CalculateBasin(x, y)
}

func (g *Grid) GetMinimums() []Point {
	points := []Point{}
	for y := 0; y < len(g.grid); y++ {
		for x := 0; x < len(g.grid[y]); x++ {
			if g.IsMinimum(x, y) {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func (g *Grid) Visit(x, y int) {
	g.visited[fmt.Sprint(x)+","+fmt.Sprint(y)] = true
}

func (g *Grid) IsMinimum(x int, y int) bool {
	me := g.Value(x, y)

	left := g.Value(x-1, y)
	right := g.Value(x+1, y)
	up := g.Value(x, y-1)
	down := g.Value(x, y+1)

	return (left == -1 || me < left) &&
		(right == -1 || me < right) &&
		(up == -1 || me < up) &&
		(down == -1 || me < down)
}

func (g *Grid) Value(x int, y int) int {
	if g.visited[fmt.Sprint(x)+","+fmt.Sprint(y)] {
		return -1
	}
	if y < 0 || y >= len(g.grid) || x < 0 || x >= len(g.grid[y]) {
		return -1
	}
	return g.grid[y][x]
}
