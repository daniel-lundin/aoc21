package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Five() {
	file, err := os.Open("./input-5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pairs := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")

		fromCoord := parseCoord(coords[0])
		toCoord := parseCoord(coords[1])

		pair := []int{fromCoord[0], fromCoord[1], toCoord[0], toCoord[1]}
		pairs = append(pairs, pair)
	}
	maxX := getMaxX(pairs) + 1
	maxY := getMaxY(pairs) + 1

	diagram := make([]int, maxX*maxY)

	for _, pair := range pairs {

		if isStraigtLine(pair) {
			startX := min(pair[0], pair[2])
			endX := max(pair[0], pair[2])
			startY := min(pair[1], pair[3])
			endY := max(pair[1], pair[3])
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					index := x + y*maxX
					diagram[index] = diagram[index] + 1
				}
			}
		} else {
			// fmt.Println(pair)
			xInc := pair[2] - pair[0]
			yInc := pair[3] - pair[1]
			currX := pair[0]
			currY := pair[1]
			// fmt.Printf("x inc %d y inc %d curX %d currY %d\n", xInc, yInc, currX, currY)
			for i := 0; i <= abs(xInc); i++ {
				// fmt.Printf("plotting at %d %d\n", currX, currY)
				index := currX + currY*maxX
				diagram[index] = diagram[index] + 1
				currX += sign(xInc)
				currY += sign(yInc)
			}

		}
	}

	// plotDiagram(diagram, maxX)
	overlaps := countOverlaps(diagram)
	fmt.Printf("overlaps %d \n", overlaps)

}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
func sign(number int) int {
	if number < 0 {
		return -1
	}
	return 1
}

func isStraigtLine(pair []int) bool {
	if pair[0] != pair[2] && pair[1] != pair[3] {
		return false
	}
	return true
}

func countOverlaps(diagram []int) int {
	overlaps := 0
	for i := 0; i < len(diagram); i++ {
		if diagram[i] >= 2 {
			overlaps += 1
		}
	}
	return overlaps
}

func plotDiagram(diagram []int, width int) {
	for i := 0; i < len(diagram); i++ {
		if i%width == 0 {
			fmt.Println()
		}
		fmt.Printf("%d", diagram[i])
	}
	fmt.Println()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getMaxX(pairs [][]int) int {
	max := 0
	for _, pair := range pairs {
		if pair[0] > max {
			max = pair[0]
		}
		if pair[2] > max {
			max = pair[2]
		}
	}
	return max
}

func getMaxY(pairs [][]int) int {
	max := 0
	for _, pair := range pairs {
		if pair[1] > max {
			max = pair[1]
		}
		if pair[3] > max {
			max = pair[3]
		}
	}
	return max
}

func parseCoord(coord string) []int {
	coords := strings.Split(coord, ",")
	x, errX := strconv.Atoi(coords[0])
	y, errY := strconv.Atoi(coords[1])
	if errX != nil {
		fmt.Println("error parsing int")
		os.Exit(1)
	}
	if errY != nil {
		fmt.Println("error parsing int")
		os.Exit(1)
	}
	return []int{x, y}

}
