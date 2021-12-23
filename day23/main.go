package main

import (
	"fmt"
	"sort"
)

func main() {

	games := []game{{
		rests: [7]int{-1, -1, -1, -1, -1, -1, -1},
		holes: [4][]int{
			{0, 1},
			{3, 2},
			{2, 1},
			{0, 3},
		},
		totalCost: 0,
	}}

	finished := []game{}

	for len(games) > 0 {
		g := games[0]
		games = games[1:]

		moves := g.getPossibleMoves()
		if len(moves) == 0 {
			continue
		}
		for _, move := range moves {
			nextGame := g.executeMove(move)
			if nextGame.isFinished() {
				finished = append(finished, nextGame)
			} else if nextGame.totalCost < 20000 {
				games = append(games, g.executeMove(move))
			} else {
				fmt.Println("Ignoring crazy game")
			}
		}
		sort.SliceStable(games, func(i, j int) bool {
			scoreA := games[i].score
			scoreB := games[j].score

			costA := games[i].totalCost
			costB := games[j].totalCost

			if scoreA == scoreB {
				return costA < costB
			}

			return scoreB < scoreA
		})
	}

	if len(finished) > 0 {
		fmt.Println("Finished", len(finished), "games")
		fmt.Println(len(games), "games left")

		sort.SliceStable(finished, func(i, j int) bool {
			return finished[i].totalCost < finished[j].totalCost
		})

		for _, g := range finished {
			fmt.Println(g.totalCost)
		}
		return
	}

	fmt.Println("Not finished - Games =", len(games))
	fmt.Println("Lowest score game so far is", games[0].totalCost)
	fmt.Println(games[0])
}

type game struct {
	rests     [7]int
	holes     [4][]int
	totalCost int
	score     int
}

type location struct {
	variant string
	index   int
	value   int
	holePos int
}

type move struct {
	from location
	to   location
	cost int
}

func (g game) isFinished() bool {
	for i, v := range g.holes {
		if len(v) != 2 {
			return false
		}
		if v[0] != i || v[1] != i {
			return false
		}
	}
	return true
}

func (g game) getScore() int {
	score := 0
	for i, v := range g.holes {
		if len(v) > 0 && v[0] == i {
			score++
		}
		if len(v) > 1 && v[1] == i {
			score++
		}
	}
	return score
}

func (g game) executeMove(m move) game {
	if m.from.variant == "rest" {
		g.rests[m.from.index] = -1
	} else {
		g.holes[m.from.index] = g.holes[m.from.index][0 : 1-m.from.holePos]
	}

	if m.to.variant == "rest" {
		g.rests[m.to.index] = m.from.value
	} else {
		g.holes[m.to.index] = append(g.holes[m.to.index], m.from.value)
	}

	g.totalCost += m.cost
	g.score = g.getScore()

	return g
}

func (g game) getPossibleMoves() []move {
	moves := []move{}

	for i, v := range g.rests {
		if v == -1 {
			continue
		}
		hole := g.holes[v]
		if len(hole) == 0 || (len(hole) == 1 && hole[0] == v) {
			from := location{variant: "rest", index: i, value: v}
			to := location{variant: "hole", index: v, holePos: 1 - len(hole)}
			moves = append(moves, createMove(from, to))
		}
	}

	for i, hole := range g.holes {
		if len(hole) == 0 {
			continue
		}

		if len(hole) == 1 && hole[0] == i {
			continue
		}

		if len(hole) == 2 && hole[0] == i && hole[1] == i {
			continue
		}

		leftIndex := i + 1
		rightIndex := i + 2

		from := location{variant: "hole", index: i, value: hole[len(hole)-1], holePos: 2 - len(hole)}

		for leftIndex >= 0 {
			val := g.rests[leftIndex]
			if val != -1 {
				break
			}
			to := location{variant: "rest", index: leftIndex}
			moves = append(moves, createMove(from, to))
			leftIndex--
		}

		for rightIndex < len(g.rests) {
			val := g.rests[rightIndex]
			if val != -1 {
				break
			}
			to := location{variant: "rest", index: rightIndex}
			moves = append(moves, createMove(from, to))
			rightIndex++
		}
	}

	sort.SliceStable(moves, func(i, j int) bool {
		a := moves[i].cost
		b := moves[j].cost
		return a < b
	})

	return moves
}

// #############
// #...B.......#
// ###B#C#.#D###
//   #A#D#C#A#
//   #########
func (g game) String() string {
	str := "#############\n#"
	for i, v := range g.rests {
		char := intToChar(v)
		str += char
		if i >= 1 && i <= 4 {
			str += "."
		}
	}
	str += "#\n###"
	for _, v := range g.holes {
		char := "."
		if len(v) >= 2 {
			char = intToChar(v[1])
		}
		str += char
		str += "#"
	}
	str += "##\n  #"
	for _, v := range g.holes {
		char := "."
		if len(v) >= 1 {
			char = intToChar(v[0])
		}
		str += char
		str += "#"
	}
	str += "  \n  #########  "
	return str
}

func intToChar(i int) string {
	switch i {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "D"
	default:
		return "."
	}
}

func createMove(from location, to location) move {
	holeItem, restItem := from, to
	if holeItem.variant == "rest" {
		holeItem, restItem = restItem, holeItem
	}

	cost := getXDiff(holeItem, restItem) + holeItem.holePos + 1 // To get to the corridor from the hole

	switch from.value {
	case 1:
		cost *= 10
	case 2:
		cost *= 100
	case 3:
		cost *= 1000
	}

	return move{from, to, cost}
}

func getXDiff(holeLocation location, rest location) int {
	holeX := (holeLocation.index * 2) + 2
	restX := rest.index
	if rest.index >= 2 {
		restX += (restX - 1)
	}
	if rest.index == 6 {
		restX--
	}

	if holeX > restX {
		return holeX - restX
	} else {
		return restX - holeX
	}
}
