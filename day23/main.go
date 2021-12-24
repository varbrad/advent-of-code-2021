package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []string{
		"#############",
		"#...........#",
		"###D#A#C#D###",
		"  #D#C#B#A#  ",
		"  #D#B#A#C#  ",
		"	 #C#A#B#B#  ",
		"	 #########  ",
	}

	// input := []string{
	// 	"#############",
	// 	"#...........#",
	// 	"###B#C#B#D###",
	// 	"  #D#C#B#A#",
	// 	"  #D#B#A#C#",
	// 	"	 #A#D#C#A#  ",
	// 	"	 #########  ",
	// }

	games := []game{parseToGame(input)}
	seenBefore := map[string]int{}
	finished := game{cost: 90000}

	for i := 0; i < 5_000_000; i++ {
		if i%10000 == 0 {
			fmt.Println(i, len(games), finished.cost)
		}
		if len(games) == 0 {
			fmt.Println("No games left... after", i, "loops")
			break
		}
		g := games[0]
		games = games[1:]

		str := g.String()

		bestCost, seen := seenBefore[str]
		if seen && bestCost < g.cost {
			continue
		}

		seenBefore[str] = g.cost

		if g.cost+g.heuristic() > finished.cost {
			continue
		}
		moves := g.getPossibleMoves()
		for _, move := range moves {
			next := g.executeMove(move)
			if next.isDone() {
				if next.cost < finished.cost {
					finished = next
				}
			} else {
				if next.cost+next.heuristic() > finished.cost {
					continue
				}
				games = append(games, next)
			}
		}

		sort.SliceStable(games, func(i, j int) bool {
			if games[i].getScore() > games[j].getScore() {
				return true
			} else if games[i].getScore() < games[j].getScore() {
				return false
			}
			return games[i].cost < games[j].cost
		})
	}

	fmt.Println(games[0])

	fmt.Println("Games: ", len(games))
	fmt.Println("Finished\n" + finished.String())
}
