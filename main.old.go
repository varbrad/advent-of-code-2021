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
	cs := makeCaveSystem(input)

	stack := [][]string{{"start"}}
	routes := [][]string{}

	for len(stack) > 0 {
		currentPath := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, cnx := range getConnections(currentPath, cs, maxVisits) {
			newPath := append(clone(currentPath), cnx)
			if cnx == "end" {
				routes = append(routes, newPath)
			} else {
				stack = append(stack, newPath)
			}
		}
	}

	return len(routes)
}

type cave struct {
	connections []string
	size        string
}

type caveSystem struct {
	lookup map[string]*cave
}

var splitRegex = regexp.MustCompile(`^(\w+)-(\w+)$`)

func makeCaveSystem(input []string) caveSystem {
	lookup := make(map[string]*cave)
	for _, line := range input {
		result := splitRegex.FindStringSubmatch(line)
		a, b := result[1], result[2]
		valA := lookup[a]
		valB := lookup[b]
		if valA == nil {
			lookup[a] = &cave{[]string{}, getCaveSize(a)}
		}
		if valB == nil {
			lookup[b] = &cave{[]string{}, getCaveSize(b)}
		}
		lookup[a].connections = append(lookup[a].connections, b)
		lookup[b].connections = append(lookup[b].connections, a)
	}
	return caveSystem{lookup}
}

func getCaveSize(origin string) string {
	if strings.ToUpper(origin) == origin {
		return "large"
	}
	return "small"
}

func getConnections(path []string, caveSystem caveSystem, limit int) []string {
	currentCave := path[len(path)-1]
	cnxs := []string{}
	cave := caveSystem.lookup[currentCave]
	if cave == nil {
		return []string{}
	}
	for _, cnx := range cave.connections {
		if !canVisit(path, cnx, limit) {
			continue
		}
		cnxs = append(cnxs, cnx)
	}
	return cnxs
}

func canVisit(list []string, value string, max int) bool {
	if value == "start" {
		return false
	}
	if value == "end" {
		return true
	}
	if strings.ToUpper(value) == value {
		return true
	}
	instances := make(map[string]int)
	alreadyHere := false
	for _, v := range list {
		if strings.ToUpper(v) == v || v == "start" || v == "end" {
			continue
		}
		if v == value {
			alreadyHere = true
		}
		instances[v]++
	}
	if alreadyHere == false {
		return true
	}
	for _, v := range instances {
		if v >= max {
			return false
		}
	}
	return true
}

func clone(slice []string) []string {
	return append([]string{}, slice...)
}
