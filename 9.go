package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Coord struct {
	row int
	col int
}

func Nine() {
	file, err := os.Open("./input-9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numberChars := strings.Split(line, "")
		grid = append(grid, StringListToNumbers(numberChars))
	}

	width := len(grid[0])
	riskLevel := 0
	lowPoints := make([]Coord, 0)

	for rowIndex, row := range grid {
		for colIndex, value := range row {
			isLowPoint := true
			if rowIndex > 0 {
				isLowPoint = isLowPoint && value < grid[rowIndex-1][colIndex]
			}
			if rowIndex < len(grid)-1 {
				isLowPoint = isLowPoint && value < grid[rowIndex+1][colIndex]
			}
			if colIndex > 0 {
				isLowPoint = isLowPoint && value < grid[rowIndex][colIndex-1]
			}
			if colIndex < width-1 {
				isLowPoint = isLowPoint && value < grid[rowIndex][colIndex+1]
			}

			if isLowPoint {
				riskLevel += value + 1
				lowPoints = append(lowPoints, Coord{rowIndex, colIndex})
			}
		}
	}

	basins := make([]int, 0)
	for _, lowPoint := range lowPoints {
		basinSize := 0
		visitedNodes := make([]Coord, 0)
		walkBasin(grid, lowPoint.row, lowPoint.col, &visitedNodes)
		basinSize = len(visitedNodes)
		basins = append(basins, basinSize)

	}

	sort.Ints(basins)

	largestBasins := basins[len(basins)-3:]
	basinProduct := 1
	for _, basin := range largestBasins {
		basinProduct *= basin
	}

	fmt.Printf("Basins %v\n", basins)
	fmt.Printf("largest Basins %v %d\n", largestBasins, basinProduct)
	fmt.Printf("Risk level: %d\n", riskLevel)
}

func walkBasin(grid [][]int, row int, col int, visited *[]Coord) {
	if hasCoord(*visited, row, col) {
		return
	}

	if row < 0 || row > len(grid)-1 {
		return
	}
	if col < 0 || col > len(grid[0])-1 {
		return
	}
	if grid[row][col] < 9 {
		*visited = append(*visited, Coord{row, col})
		walkBasin(grid, row+1, col, visited)
		walkBasin(grid, row, col+1, visited)
		walkBasin(grid, row-1, col, visited)
		walkBasin(grid, row, col-1, visited)
	}

}

func hasCoord(coords []Coord, row int, col int) bool {
	for _, coord := range coords {
		if coord.row == row && coord.col == col {
			return true
		}
	}
	return false
}
