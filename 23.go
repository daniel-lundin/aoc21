package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Ampiphode struct {
	kind string
	x    int
	y    int
}

func TwentyThree() {
	ampiphodes := map[Coord2D]string{
		{2, 1}: "A",
		{2, 2}: "D",
		{2, 3}: "D",
		{2, 4}: "D",
		{4, 1}: "C",
		{4, 2}: "C",
		{4, 3}: "B",
		{4, 4}: "D",
		{6, 1}: "B",
		{6, 2}: "B",
		{6, 3}: "A",
		{6, 4}: "B",
		{8, 1}: "A",
		{8, 2}: "A",
		{8, 3}: "C",
		{8, 4}: "C",
	}

	moves := possibleMoves(ampiphodes)

	lowestCost := math.MaxInt
	makeMove(moves, ampiphodes, 0, &lowestCost)

	fmt.Printf("lowest cost %d\n", lowestCost)
}

func makeMove(moves [][]Coord2D, ampis map[Coord2D]string, cost int, lowestCost *int) {

	for _, move := range moves {
		clonedAmpis := cloneAmpis(ampis)
		ampi := clonedAmpis[move[0]]
		clonedAmpis[move[0]] = ""
		clonedAmpis[move[1]] = ampi
		newMoves := possibleMoves(clonedAmpis)

		moveCost := (intAbs(move[0].x-move[1].x) + move[0].y + move[1].y) * getMoveCost(ampis[move[0]])

		newCost := moveCost + cost
		if newCost >= *lowestCost {
			continue
		}
		if len(newMoves) > 0 {
			makeMove(newMoves, clonedAmpis, newCost, lowestCost)
		} else {
			if allInPlace(clonedAmpis) {
				if newCost < *lowestCost {
					fmt.Printf("New lowest cost %d\n", newCost)
					*lowestCost = newCost
				}
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
	movesAndCosts := make([][]int, 0)
	for coord, ampi := range ampis {
		if ampi == "" {
			continue
		}
		freeXSpots := hallwaySpots(coord.x, ampis)
		desinationX := sideroomX(ampi)
		possibleDesination := allowedHome(ampi, ampis)
		// Bottom ampi
		if coord.y == 4 {
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
					movesAndCosts = append(movesAndCosts, []int{coord.x, coord.y, desinationX, possibleDesination, getMoveCost(ampi)})
					continue
				}
			}

			for _, x := range freeXSpots {
				// fmt.Printf("found move %v\n", []Coord2D{coord, {x, 0}})
				movesAndCosts = append(movesAndCosts, []int{coord.x, coord.y, x, 0, getMoveCost(ampi)})
			}

		}

		if coord.y > 0 {
			// Correct place
			if coord.x == sideroomX(ampi) && correctBelow(ampi, coord.x, coord.y, ampis) {
				continue
			}

			if !emptyAbove(coord.x, coord.y, ampis) {
				continue
			}
			// Find hallways spots and final destinations
			if possibleDesination != -1 {
				if freePassage(coord.x, desinationX, ampis) {
					movesAndCosts = append(movesAndCosts, []int{coord.x, coord.y, desinationX, possibleDesination, getMoveCost(ampi)})
					continue
				}
			}

			for _, freeXSpot := range freeXSpots {
				movesAndCosts = append(movesAndCosts, []int{coord.x, coord.y, freeXSpot, 0, getMoveCost(ampi)})
			}
		}

		if coord.y == 0 {
			if possibleDesination != -1 {
				if freePassage(coord.x, desinationX, ampis) {
					movesAndCosts = append(movesAndCosts, []int{coord.x, coord.y, desinationX, possibleDesination, getMoveCost(ampi)})
					continue
				}
			}
		}
	}

	sort.Sort(ByCost(movesAndCosts))
	moves := make([][]Coord2D, 0)
	for _, m := range movesAndCosts {
		moves = append(moves, []Coord2D{{m[0], m[1]}, {m[2], m[3]}})
	}
	return moves
}

type ByCost [][]int

func (a ByCost) Len() int           { return len(a) }
func (a ByCost) Less(i, j int) bool { return a[i][4] < a[j][4] }
func (a ByCost) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func correctBelow(ampi string, x int, y int, ampis map[Coord2D]string) bool {
	for yy := y + 1; yy <= 4; yy++ {
		value := ampis[Coord2D{x, yy}]
		if value != ampi && value != "" {
			return false
		}
	}
	return true
}

func emptyAbove(x, y int, ampis map[Coord2D]string) bool {
	for yy := y - 1; yy >= 0; yy-- {
		if ampis[Coord2D{x, yy}] != "" {
			return false
		}
	}
	return true
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
	for y := 1; y <= 4; y++ {
		if ampis[Coord2D{x, y}] != "" {
			return false
		}
	}
	return true
}

func allowedHome(ampi string, ampis map[Coord2D]string) int {
	destinationX := sideroomX(ampi)
	if sideroomEmpty(destinationX, ampis) {
		return 4
	}
	y1 := ampis[Coord2D{destinationX, 1}]
	y2 := ampis[Coord2D{destinationX, 2}]
	y3 := ampis[Coord2D{destinationX, 3}]
	y4 := ampis[Coord2D{destinationX, 4}]
	if y4 == ampi && y3 == "" {
		return 3
	}
	if y4 == ampi && y3 == ampi && y2 == "" {
		return 2
	}
	if y4 == ampi && y3 == ampi && y2 == ampi && y1 == "" {
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

	for y := 2; y <= 4; y++ {
		fmt.Printf("  #")
		for x := 2; x <= 8; x++ {
			ambi := ampis[Coord2D{x, y}]
			if ambi != "" {
				fmt.Printf(ambi)
			} else if x%2 == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("#\n")
	}
	fmt.Println("  #########  ")
}
