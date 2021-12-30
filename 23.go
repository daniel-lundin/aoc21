package main

import (
	"fmt"
	"time"
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
	ampiphodes := map[Coord2D]string{
		{2, 1}: "B",
		{2, 2}: "A",
		{4, 1}: "C",
		{4, 2}: "D",
		{6, 1}: "B",
		{6, 2}: "C",
		{8, 1}: "D",
		{8, 2}: "A",
	}

	moves := possibleMoves(ampiphodes)

	highestCost := 0
	makeMove(moves, ampiphodes, 0, &highestCost)
}

func makeMove(moves [][]Coord2D, ampis map[Coord2D]string, cost int, highestCost *int) {

	for _, move := range moves {
		clonedAmbis := cloneAmpis(ampis)
		printAmbis(ampis)
		fmt.Printf("moving %v\n", move)
		ampi := clonedAmbis[move[0]]
		clonedAmbis[move[0]] = ""
		clonedAmbis[move[1]] = ampi
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("after move:\n")
		printAmbis(clonedAmbis)
		newMoves := possibleMoves(clonedAmbis)

		if len(newMoves) > 0 {
			makeMove(newMoves, clonedAmbis, cost+1, highestCost)
		}
	}
}

func printAmbis(ampis map[Coord2D]string) {
	fmt.Println("#############")
	fmt.Printf("#")
	for x := 0; x < 11; x++ {
		ambi := ampis[Coord2D{x, 0}]
		if ambi != "" {
			fmt.Printf(ambi)
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Printf("#\n")
	fmt.Printf("###")
	for x := 2; x < 8; x++ {
		ambi := ampis[Coord2D{x, 1}]
		if ambi != "" {
			fmt.Printf(ambi)
		} else if x%2 == 0 {
			fmt.Printf(".")
		} else {
			fmt.Printf("#")
		}
	}
	fmt.Printf("###\n")

	fmt.Printf("  #")
	for x := 2; x < 8; x++ {
		ambi := ampis[Coord2D{x, 2}]
		if ambi != "" {
			fmt.Printf(ambi)
		} else if x%2 == 0 {
			fmt.Printf(".")
		} else {
			fmt.Printf("#")
		}
	}
	fmt.Printf("#\n")
	fmt.Println("  #########  ")
}

func cloneAmpis(ampis map[Coord2D]string) map[Coord2D]string {
	newMap := make(map[Coord2D]string)
	for k, v := range ampis {
		newMap[k] = v
	}
	return newMap
}

func possibleMoves(ampis map[Coord2D]string) [][]Coord2D {
	moves := make([][]Coord2D, 0)
	for coord, ampi := range ampis {
		// fmt.Printf("Checking possible moves for %s %v\n", ampi, coord)
		freeXSpots := hallwaySpots(coord.x, ampis)
		desinationX := sideroomX(ampi)
		possibleDesination := freeDestinations(ampi, ampis)
		// Bottom ampi
		if coord.y == 2 {
			// Where I'm supposed to be
			if coord.x == sideroomX(ampi) {
				continue
			}
			// Blocked
			if ampis[Coord2D{coord.x, coord.y - 1}] != "" {
				continue
			}

			// Final destination possible
			if possibleDesination != -1 {
				if freePassage(coord.x, desinationX, freeXSpots) {
					// fmt.Printf("found move %v\n", []Coord2D{coord, {desinationX, possibleDesination}})
					moves = append(moves, []Coord2D{coord, {desinationX, possibleDesination}})
					continue
				}
			}

			for _, x := range freeXSpots {
				// fmt.Printf("found move %v\n", []Coord2D{coord, {x, 0}})
				moves = append(moves, []Coord2D{coord, {x, 0}})
			}

		}

		if coord.y == 1 {
			// Find hallways spots and final destinations
			if possibleDesination != -1 {
				if freePassage(coord.x, desinationX, freeXSpots) {
					// fmt.Printf("found move %v\n", []Coord2D{coord, {desinationX, possibleDesination}})
					moves = append(moves, []Coord2D{coord, {desinationX, possibleDesination}})
					continue
				}
			}

			for _, freeXSpot := range freeXSpots {
				// fmt.Printf("found move %v\n", []Coord2D{coord, {freeXSpot, 0}})
				moves = append(moves, []Coord2D{coord, {freeXSpot, 0}})
			}
		}

		if coord.y == 0 {
			if possibleDesination != -1 {
				if freePassage(coord.x, desinationX, freeXSpots) {
					// fmt.Printf("found move %v\n", []Coord2D{coord, {desinationX, possibleDesination}})
					moves = append(moves, []Coord2D{coord, {desinationX, possibleDesination}})
					continue
				}
			}
		}
	}
	return moves
}

func freePassage(fromX int, toX int, freeXSpots []int) bool {
	minX := minInt(fromX, toX)
	maxX := maxInt(fromX, toX)
	for x := minX; x < maxX; x++ {
		free := false
		for _, freeX := range freeXSpots {
			if freeX == x {
				free = true
			}
		}
		if !free {
			return false
		}
	}
	return true
}

func sideroomX(ampi string) int {
	if ampi == "A" {
		return 2
	}
	if ampi == "B" {
		return 4
	}
	if ampi == "C" {
		return 6
	}
	if ampi == "D" {
		return 8
	}
	return -1
}

func sideroomEmpty(x int, ampis map[Coord2D]string) bool {
	return ampis[Coord2D{x, 1}] == "" && ampis[Coord2D{x, 2}] == ""
}

func freeDestinations(ampi string, ampis map[Coord2D]string) int {
	desinationX := sideroomX(ampi)
	if sideroomEmpty(desinationX, ampis) {
		return 2
	}
	if ampis[Coord2D{desinationX, 2}] == ampi && ampis[Coord2D{desinationX, 1}] == "" {
		return 1
	}
	return -1
}

func hallwaySpots(startX int, ampis map[Coord2D]string) []int {
	possibleSpots := make([]int, 0)
	// left
	for x := startX; x >= 0; x-- {
		if ampis[Coord2D{x, 0}] == "" {
			// Not outside doors
			if x != 2 && x != 4 && x != 6 && x != 8 {
				possibleSpots = append(possibleSpots, x)
			}
		}
	}

	// right
	for x := startX; x <= 10; x++ {
		if ampis[Coord2D{x, 0}] == "" {
			// Not outside doors
			if x != 2 && x != 4 && x != 6 && x != 8 {
				possibleSpots = append(possibleSpots, x)
			}
		}
	}
	return possibleSpots
}
