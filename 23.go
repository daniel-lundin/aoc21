package main

import (
	"fmt"
	"math"
	"os"
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
	// ampiphodes := map[Coord2D]string{
	// 	{2, 1}: "A",
	// 	{2, 2}: "D",
	// 	{2, 3}: "D",
	// 	{2, 4}: "D",
	// 	{4, 1}: "C",
	// 	{4, 2}: "C",
	// 	{4, 3}: "B",
	// 	{4, 4}: "D",
	// 	{6, 1}: "B",
	// 	{6, 2}: "B",
	// 	{6, 3}: "A",
	// 	{6, 4}: "B",
	// 	{8, 1}: "A",
	// 	{8, 2}: "A",
	// 	{8, 3}: "C",
	// 	{8, 4}: "C",
	// }

	moves := possibleMoves(ampiphodes)

	lowestCost := math.MaxInt
	makeMove(moves, ampiphodes, 0, &lowestCost)

	fmt.Printf("lowest cost %d\n", lowestCost)
}

func makeMove(moves [][]Coord2D, ampis map[Coord2D]string, cost int, lowestCost *int) {

	for _, move := range moves {
		clonedAmpis := cloneAmpis(ampis)
		//fmt.Println()
		//printAmbis(ampis)
		//fmt.Printf("moving %v\n", move)
		ampi := clonedAmpis[move[0]]
		clonedAmpis[move[0]] = ""
		clonedAmpis[move[1]] = ampi
		//fmt.Printf("after move:\n")
		//printAmbis(clonedAmpis)
		//fmt.Println()
		//time.Sleep(500 * time.Millisecond)
		newMoves := possibleMoves(clonedAmpis)

		moveCost := (intAbs(move[0].x-move[1].x) + move[0].y + move[1].y) * getMoveCost(ampis[move[0]])

		if moveCost+cost > *lowestCost {
			continue
		}
		if len(newMoves) > 0 {
			makeMove(newMoves, clonedAmpis, cost+moveCost, lowestCost)
		} else {
			// fmt.Printf("no more moves\n ")
			// printAmbis(clonedAmpis)
			if allInPlace(clonedAmpis) {
				*lowestCost = minInt(cost+moveCost, *lowestCost)
				//fmt.Printf("good posidition %d", cost+moveCost)
				// printAmbis(clonedAmpis)
			}
		}
	}
}

func allInPlace(ampis map[Coord2D]string) bool {
	for coord, ampi := range ampis {
		if ampi == "" {
			continue
		}
		if coord.x != sideroomX(ampi) {
			return false
		}

	}
	return true
}

func intAbs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func getMoveCost(ampi string) int {
	if ampi == "A" {
		return 1
	}
	if ampi == "B" {
		return 10
	}
	if ampi == "C" {
		return 100
	}
	if ampi == "D" {
		return 1000
	}
	fmt.Printf("Unknown ampi %s\n", ampi)
	os.Exit(1)
	return -1
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
		if ampi == "" {
			continue
		}
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
				if freePassage(coord.x, desinationX, ampis) {
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
			// Correct place
			if coord.x == sideroomX(ampi) && ampis[Coord2D{coord.x, 2}] == ampi {
				continue
			}
			// Find hallways spots and final destinations
			if possibleDesination != -1 {
				if freePassage(coord.x, desinationX, ampis) {
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
				if freePassage(coord.x, desinationX, ampis) {
					// fmt.Printf("found move %v\n", []Coord2D{coord, {desinationX, possibleDesination}})
					moves = append(moves, []Coord2D{coord, {desinationX, possibleDesination}})
					continue
				}
			}
		}
	}
	return moves
}

func freePassage(fromX int, toX int, ampis map[Coord2D]string) bool {
	minX := minInt(fromX, toX)
	maxX := maxInt(fromX, toX)
	for x := minX; x < maxX; x++ {
		if x == fromX {
			continue
		}
		if ampis[Coord2D{x, 0}] != "" {
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
	destinationX := sideroomX(ampi)
	if sideroomEmpty(destinationX, ampis) {
		return 2
	}
	if ampis[Coord2D{destinationX, 2}] == ampi && ampis[Coord2D{destinationX, 1}] == "" {
		return 1
	}
	return -1
}

func hallwaySpots(startX int, ampis map[Coord2D]string) []int {
	possibleSpots := make([]int, 0)
	// left
	for x := startX - 1; x >= 0; x-- {
		if ampis[Coord2D{x, 0}] == "" {
			// Not outside doors
			if x != 2 && x != 4 && x != 6 && x != 8 {
				possibleSpots = append(possibleSpots, x)
			}
		} else {
			break
		}
	}

	// right
	for x := startX + 1; x <= 10; x++ {
		if ampis[Coord2D{x, 0}] == "" {
			// Not outside doors
			if x != 2 && x != 4 && x != 6 && x != 8 {
				possibleSpots = append(possibleSpots, x)
			}
		} else {
			break
		}
	}
	return possibleSpots
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
	for x := 2; x <= 8; x++ {
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
	for x := 2; x <= 8; x++ {
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
