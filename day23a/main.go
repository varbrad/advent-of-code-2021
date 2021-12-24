package main

import "fmt"

func main() {
	input := []string{
		"#############",
		"#...........#",
		"###D#A#C#D###",
		"	 #C#A#B#B#  ",
		"	 #########  ",
	}

	parseToGame(input)
}

type game struct {
	hallway [11]string
}

func parseToGame(input []string) game {
	hallway := [11]string{}
	for i := 0; i < 11; i++ {
		bytes := []byte{input[1][1]}
		hallway[i] = string(bytes)
	}

	fmt.Println(hallway)

	return game{}
}
