package main

import (
	"fmt"
	"math"
)

type game struct {
	rests    string
	holes    [4]string
	holeSize int
	cost     int
}

type location struct {
	mode      string
	index     int
	holeIndex int
}

type move struct {
	from location
	to   location
	cost int
}

func makeInitialGame(holeSize int, a, b, c, d []rune) game {
	return game{
		rests:    "...........",
		holes:    [4]string{string(a), string(b), string(c), string(d)},
		holeSize: holeSize,
	}
}

var restMap = map[int]bool{
	0:  true,
	1:  true,
	2:  false,
	3:  true,
	4:  false,
	5:  true,
	6:  false,
	7:  true,
	8:  false,
	9:  true,
	10: true,
}

func (g game) executeMove(m move) game {

	value := '.'
	if m.from.mode == "rest" {
		value = rune(g.rests[m.from.index])
		g.rests = replaceAtIndex(g.rests, '.', m.from.index)
	} else {
		value = rune(g.holes[m.from.index][m.from.holeIndex])
		g.holes[m.from.index] = g.holes[m.from.index][0:m.from.holeIndex]
	}

	if value == '.' {
		panic("A value of '.' was encountered!")
	}

	if m.to.mode == "rest" {
		g.rests = replaceAtIndex(g.rests, value, m.to.index)
	} else {
		g.holes[m.to.index] += string(value)
	}

	g.cost = g.cost + m.cost

	return g
}

func (g game) getScore() int {
	score := 0
	for hx, hole := range g.holes {
		allRight := true
	inner:
		for _, item := range hole {
			if !rightHoleType(item, hx) {
				allRight = false
				break inner
			}
		}
		if allRight {
			score += g.holeSize
		}
	}
	return score
}

func (g game) getPossibleMoves() []move {
	moves := []move{}

	for ix, item := range g.rests {
		ok := restMap[ix]
		if !ok || item == '.' {
			continue
		}

		from := location{mode: "rest", index: ix}
		leftBound := maxLeft(g, ix)
		rightBound := maxRight(g, ix)

		for hx, hole := range g.holes {
			// Skip this hole if
			// 1. it's full
			// 2. it's the wrong type
			// 3. its too far left
			// 4. it's too far right
			// 5. It has wrong elements in it
			holeX := 2 + (hx * 2)
			if len(hole) == g.holeSize || !rightHoleType(item, hx) || holeX < leftBound || holeX > rightBound {
				continue
			}
			ok := true
			for _, holeItem := range hole {
				if !rightHoleType(holeItem, hx) {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}

			to := location{mode: "hole", index: hx, holeIndex: len(hole)}
			moves = append(moves, g.createMove(from, to))
		}
	}

	for ix, hole := range g.holes {
		itemsInHole := len(hole)
		if itemsInHole == 0 {
			continue
		}

		if itemsInHole == 1 {
			if rightHoleType(rune(hole[0]), ix) {
				continue
			}
		}

		if itemsInHole == 2 {
			if rightHoleType(rune(hole[0]), ix) && rightHoleType(rune(hole[1]), ix) {
				continue
			}
		}

		if itemsInHole == 3 {
			if rightHoleType(rune(hole[0]), ix) && rightHoleType(rune(hole[1]), ix) && rightHoleType(rune(hole[2]), ix) {
				continue
			}
		}

		if itemsInHole == 4 {
			if rightHoleType(rune(hole[0]), ix) && rightHoleType(rune(hole[1]), ix) && rightHoleType(rune(hole[2]), ix) && rightHoleType(rune(hole[3]), ix) {
				continue
			}
		}

		topItemIndex := itemsInHole - 1

		from := location{
			mode:      "hole",
			index:     ix,
			holeIndex: topItemIndex,
		}

		holeX := 2 + (ix * 2)

		leftBound := maxLeft(g, holeX)
		rightBound := maxRight(g, holeX)

		for i := leftBound; i <= rightBound; i++ {
			okRest := restMap[i]
			if !okRest {
				continue
			}
			to := location{mode: "rest", index: i}
			moves = append(moves, g.createMove(from, to))
		}

	}

	return moves
}

func (g game) isDone() bool {
	for hx, hole := range g.holes {
		if len(hole) != g.holeSize {
			return false
		}
		for _, item := range hole {
			if !rightHoleType(item, hx) {
				return false
			}
		}
	}
	return true
}

func (g game) String() string {
	str := "#############\n#"
	for _, r := range g.rests {
		str += string(r)
	}
	for i := 0; i < g.holeSize; i++ {
		if i == 0 {
			str += "#\n###"
		} else {
			str += "  #"
		}

		for _, hole := range g.holes {
			r := '.'
			if len(hole) >= g.holeSize-i {
				r = rune(hole[g.holeSize-i-1])
			}
			str += string(r) + "#"
		}

		if i == 0 {
			str += "##\n"
		} else {
			str += "  \n"
		}
	}
	str += "  #########  \nTotal cost: " + fmt.Sprint(g.cost) + "\nTotal score: " + fmt.Sprint(g.getScore()) + "\n"

	return str
}

func maxLeft(g game, x int) int {
	for _x := x - 1; _x >= 0; _x-- {
		if g.rests[_x] != '.' {
			return _x + 1
		}
	}
	return 0
}

func maxRight(g game, x int) int {
	for _x := x + 1; _x < len(g.rests); _x++ {
		if g.rests[_x] != '.' {
			return _x - 1
		}
	}
	return len(g.rests) - 1
}

func (g game) createMove(from, to location) move {
	hole, rest := from, to
	if hole.mode == "rest" {
		hole, rest = rest, hole
	}

	holeX := 2 + (hole.index * 2)
	restX := rest.index
	holeY := g.holeSize - hole.holeIndex

	cost := int(math.Abs(float64(holeX-restX))) + holeY

	value := g.getValue(from)
	if value == '.' {
		panic("Value of '.' encountered!")
	}
	switch value {
	case 'B':
		cost *= 10
	case 'C':
		cost *= 100
	case 'D':
		cost *= 1000
	}

	return move{from, to, cost}
}

func (g game) getValue(l location) rune {
	if l.mode == "rest" {
		return rune(g.rests[l.index])
	}
	return rune(g.holes[l.index][l.holeIndex])
}

func rightHoleType(item rune, x int) bool {
	switch item {
	case 'A':
		return x == 0
	case 'B':
		return x == 1
	case 'C':
		return x == 2
	case 'D':
		return x == 3
	}
	return false
}

func runeToInt(r rune) int {
	switch r {
	case 'A':
		return 0
	case 'B':
		return 1
	case 'C':
		return 2
	case 'D':
		return 3
	}
	return -1
}

func parseToGame(input []string) game {
	rests := "..........."
	for i := 1; i <= 11; i++ {
		rests = replaceAtIndex(rests, rune(input[1][i]), i-1)
	}

	holes := [4]string{}

	for i := 0; i < 4; i++ {
		stack := []rune{}
		for j := len(input) - 2; j >= 2; j-- {
			v := rune(input[j][3+i*2])
			if v != '.' {
				stack = append(stack, v)
			} else {
				break
			}
		}
		holes[i] = string(stack)
	}

	return game{
		rests:    rests,
		holes:    holes,
		holeSize: len(input) - 3,
		cost:     0,
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func (g game) heuristic() int {
	total := 0

	for hx, hole := range g.holes {
		holeX := 2 + (2 * hx)
		for _, item := range hole {
			idealX := 2 + (2 * runeToInt(item))
			diff := 1 + int(math.Abs(float64(holeX-idealX)))
			switch item {
			case 'B':
				diff *= 10
			case 'C':
				diff *= 100
			case 'D':
				diff *= 1000
			}
			total += diff
		}
	}

	for ix, item := range g.rests {
		if item == '.' {
			continue
		}
		holeX := 2 + (2 * runeToInt(item))
		diff := int(math.Abs(float64(holeX - ix)))
		switch item {
		case 'B':
			diff *= 10
		case 'C':
			diff *= 100
		case 'D':
			diff *= 1000
		}
		total += diff
	}

	return total / 2
}
