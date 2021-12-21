package main

import (
	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := []int{10, 4}
	utils.Day(21, "Dirac Dice").Input(input).Part1(Day21Part1).Part2(Day21Part2)
}

func Day21Part1(input []int) int {
	players := [2]*[2]int{
		{input[0] - 1, 0},
		{input[1] - 1, 0},
	}
	diceRolls := 0
	for {
		for _, p := range players {
			for i := 0; i < 3; i++ {
				diceRolls++
				p[0] = (p[0] + diceRolls) % 10
			}
			p[1] += (p[0] + 1)
			if p[1] >= 1000 {
				return utils.MinInteger(players[0][1], players[1][1]) * diceRolls
			}
		}
	}
}

func Day21Part2(input []int) int {
	wins := makeDiceGame().simulate(input[0]-1, input[1]-1, 0, 0)
	return utils.MaxInteger(wins[0], wins[1])
}

type diceGame struct {
	cache     *map[[4]int][2]int
	diceRolls *[27]int
}

func makeDiceGame() *diceGame {
	rolls := [27]int{}
	for i := 0; i < 27; i++ {
		rolls[i] = i/9 + (i/3)%3 + i%3 + 3
	}
	return &diceGame{
		cache:     &map[[4]int][2]int{},
		diceRolls: &rolls,
	}
}

func (ddg *diceGame) readCache(p1, p2, s1, s2 int) ([2]int, bool) {
	result, ok := (*ddg.cache)[[4]int{p1, p2, s1, s2}]
	return result, ok
}

func (ddg *diceGame) writeCache(p1, p2, s1, s2 int, wins [2]int) [2]int {
	(*ddg.cache)[[4]int{p1, p2, s1, s2}] = wins
	return wins
}

func (ddg *diceGame) simulate(p1, p2, s1, s2 int) [2]int {
	if s1 >= 21 {
		return [2]int{1, 0}
	} else if s2 >= 21 {
		return [2]int{0, 1}
	}
	if cv, ch := ddg.readCache(p1, p2, s1, s2); ch {
		return cv
	}
	wins := [2]int{0, 0}
	for _, ds := range *ddg.diceRolls {
		newPos := (p1 + ds) % 10
		newScore := s1 + newPos + 1

		results := ddg.simulate(p2, newPos, s2, newScore)
		wins[0] += results[1]
		wins[1] += results[0]
	}
	return ddg.writeCache(p1, p2, s1, s2, wins)
}
