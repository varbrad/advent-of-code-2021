package main

import (
	"fmt"
	"regexp"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day13/input")
	utils.Day(13, "Transparent Origami").Input(input).Part1(Day13Part1).Part2(Day13Part2)
}

func Day13Part1(input []string) int {
	paper := toPaper(input)
	paper.fold(0)
	return len(paper.points)
}

func Day13Part2(input []string) string {
	paper := toPaper(input)
	for i := range paper.folds {
		paper.fold(i)
	}
	fmt.Println(paper.Print())
	return "See above!"
}

type point struct {
	x int
	y int
}

type fold struct {
	dir   string
	value int
}

type paper struct {
	points map[point]int
	folds  []fold
}

func toPaper(input []string) paper {
	points := map[point]int{}
	folds := []fold{}

	pointR := regexp.MustCompile(`(\d+),(\d+)`)
	foldR := regexp.MustCompile(`fold along (x|y)=(\d+)`)

	for _, line := range input {
		if pointMatches := pointR.FindStringSubmatch(line); pointMatches != nil {
			x, y := utils.ToInteger(pointMatches[1]), utils.ToInteger(pointMatches[2])
			points[point{x, y}]++
		} else if foldMatches := foldR.FindStringSubmatch(line); foldMatches != nil {
			folds = append(
				folds,
				fold{dir: foldMatches[1], value: utils.ToInteger(foldMatches[2])},
			)
		}
	}

	return paper{points, folds}
}

func (p *paper) fold(index int) {
	fold := p.folds[index]
	for cp := range p.points {
		if (fold.dir == "x" && cp.x <= fold.value) || (fold.dir == "y" && cp.y <= fold.value) {
			continue
		}
		newPoint := point{cp.x, cp.y}
		if fold.dir == "x" {
			newPoint.x -= (cp.x - fold.value) * 2
		} else {
			newPoint.y -= (cp.y - fold.value) * 2
		}
		delete(p.points, cp)
		p.points[newPoint]++
	}
}

func (p *paper) Print() string {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for cp := range p.points {
		if cp.x < minX {
			minX = cp.x
		} else if cp.x > maxX {
			maxX = cp.x
		}
		if cp.y < minY {
			minY = cp.y
		} else if cp.y > maxY {
			maxY = cp.y
		}
	}
	out := ""
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if p.points[point{x, y}] == 0 {
				out += "  "
			} else {
				out += "██"
			}
		}
		if y < maxY {
			out += "\n"
		}
	}
	return out
}
