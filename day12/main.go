package main

import (
	"regexp"
	"strings"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day12/input")
	utils.Day(12, "Passage Pathing").Input(input).Part1(Day12Part1).Part2(Day12Part2)
}

func Day12Part1(input []string) int {
	return solve(input, 1)
}

func Day12Part2(input []string) int {
	return solve(input, 2)
}

func solve(input []string, maxVisits int) int {
	cs := createCaveSystem(input)
	stack := [][]*cave{{cs.lookup["start"]}}
	routes := [][]*cave{}

	for len(stack) > 0 {
		currentPath := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, cnx := range getConnections(currentPath, maxVisits) {
			newPath := append(clone(currentPath), cnx)
			if cnx.name == "end" {
				routes = append(routes, newPath)
			} else {
				stack = append(stack, newPath)
			}
		}

	}

	return len(routes)
}

func createCaveSystem(input []string) *caveSystem {
	splitRegex := regexp.MustCompile(`^(\w+)-(\w+)$`)
	lookup := make(map[string]*cave)
	for _, line := range input {
		result := splitRegex.FindStringSubmatch(line)
		caveA := upsertCave(&lookup, result[1])
		caveB := upsertCave(&lookup, result[2])
		caveA.connections = append(caveA.connections, caveB)
		caveB.connections = append(caveB.connections, caveA)
	}
	return &caveSystem{lookup}
}

type cave struct {
	name        string
	size        string
	connections []*cave
}

type caveSystem struct {
	lookup map[string]*cave
}

func getCaveSize(origin string) string {
	if strings.ToUpper(origin) == origin {
		return "large"
	}
	return "small"
}

func upsertCave(lookup *map[string]*cave, name string) *cave {
	v := (*lookup)[name]
	if v == nil {
		(*lookup)[name] = &cave{name: name, size: getCaveSize(name), connections: []*cave{}}
	}
	return (*lookup)[name]
}

func getConnections(currentPath []*cave, maxVisits int) []*cave {
	c := currentPath[len(currentPath)-1]
	cnxs := []*cave{}

	for _, cnx := range c.connections {
		if !canVisit(currentPath, cnx, maxVisits) {
			continue
		}
		cnxs = append(cnxs, cnx)
	}
	return cnxs
}

func canVisit(path []*cave, next *cave, maxVisits int) bool {
	if next.size == "large" {
		return true
	}
	if next.name == "start" {
		return false
	}
	if next.name == "end" {
		return true
	}
	instances := make(map[string]int)
	alreadyHere := false
	for _, c := range path {
		if c.size == "large" || c.name == "start" || c.name == "end" {
			continue
		}
		if c == next {
			alreadyHere = true
		}
		instances[c.name]++
	}
	if alreadyHere == false {
		return true
	}
	for _, v := range instances {
		if v >= maxVisits {
			return false
		}
	}
	return true
}

func clone(slice []*cave) []*cave {
	return append([]*cave{}, slice...)
}
