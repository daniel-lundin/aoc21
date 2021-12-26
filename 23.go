package main

import (
	"fmt"
	"strings"
)

type Ampiphode struct {
	kind string
	x    int
	y    int
}

func TwentyThree() {

	// #############
	// #...........#
	// ###B#C#B#D###
	//   #A#D#C#A#
	//   #########
	// energyUsed := 0
	ampiphodes := []Ampiphode{
		{"B", 3, 1},
		{"A", 3, 2},
		{"C", 5, 1},
		{"D", 5, 2},
		{"B", 7, 1},
		{"C", 7, 2},
		{"D", 9, 1},
		{"A", 9, 2},
	}

	fmt.Printf("Possible moves %v\n", ampiphodes)
	input := `
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`
	// grid := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
	}

}
