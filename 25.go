package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func TwentyFive() {
	file, err := os.Open("./input-25.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cucumbers := map[Coord2D]string{}
	y := 0
	maxX := 0
	maxY := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "")
		for x, token := range tokens {
			cucumbers[Coord2D{x, y}] = token
			maxX = maxInt(x, maxX)
		}
		y++
		maxY = maxInt(y, maxY)
	}

	fmt.Println(cucumbers)
	steps := 1
	printCucumbers(cucumbers, maxX, maxY)
	eastMoves := moveEast(cucumbers, maxX, maxY)
	southMoves := moveSouth(cucumbers, maxX, maxY)
	fmt.Printf("After %d steps\n", steps)
	printCucumbers(cucumbers, maxX, maxY)
	for len(eastMoves) > 0 || len(southMoves) > 0 {
		eastMoves = moveEast(cucumbers, maxX, maxY)
		southMoves = moveSouth(cucumbers, maxX, maxY)
		steps++
		fmt.Printf("After %d steps\n", steps)
		// printCucumbers(cucumbers, maxX, maxY)

	}

	fmt.Printf("Stopped moving afters %d steps\n", steps)

}

func printCucumbers(cucumbers map[Coord2D]string, maxX, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			fmt.Printf(cucumbers[Coord2D{x, y}])
		}
		fmt.Println()
	}
	fmt.Println()
}

func moveEast(cucumbers map[Coord2D]string, maxX, maxY int) []Coord2D {
	moves := []Coord2D{}
	for y := 0; y < maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if cucumbers[Coord2D{x, y}] == ">" {
				nextX := (x + 1) % (maxX + 1)
				if cucumbers[Coord2D{nextX, y}] == "." {
					moves = append(moves, Coord2D{x, y})
				}
			}
		}
	}
	for _, move := range moves {
		nextX := (move.x + 1) % (maxX + 1)
		cucumbers[Coord2D{move.x, move.y}] = "."
		cucumbers[Coord2D{nextX, move.y}] = ">"
	}
	return moves
}

func moveSouth(cucumbers map[Coord2D]string, maxX, maxY int) []Coord2D {
	moves := []Coord2D{}
	for y := 0; y < maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if cucumbers[Coord2D{x, y}] == "v" {
				nextY := (y + 1) % maxY
				if cucumbers[Coord2D{x, nextY}] == "." {
					moves = append(moves, Coord2D{x, y})
				}
			}
		}
	}
	for _, move := range moves {
		nextY := (move.y + 1) % maxY
		cucumbers[Coord2D{move.x, move.y}] = "."
		cucumbers[Coord2D{move.x, nextY}] = "v"
	}
	return moves
}
