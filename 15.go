package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Fifteen() {
	file, err := os.Open("./input-15.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	riskMap := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := StringListToNumbers(strings.Split(line, ""))
		expandedRow := make([]int, 5*len(row))
		for i := 0; i < 5; i++ {
			for rowIndex, value := range row {
				expandedRow[i*len(row)+rowIndex] = (value+i-1)%9 + 1
			}

		}
		riskMap = append(riskMap, expandedRow)
	}

	expandedRiskMap := make([][]int, 0)
	for i := 0; i < 5; i++ {
		for _, row := range riskMap {
			updatedRow := make([]int, len(row))
			for index, value := range row {
				updatedRow[index] = (value+i-1)%9 + 1
			}
			expandedRiskMap = append(expandedRiskMap, updatedRow)
		}
	}

	shortestPath := math.MaxInt
	lowestCosts := make([][]int, len(expandedRiskMap))
	for rowIndex, row := range expandedRiskMap {
		lowestCosts[rowIndex] = make([]int, len(row))
		for colIndex := range row {
			lowestCosts[rowIndex][colIndex] = math.MaxInt
		}
	}
	walkCave(expandedRiskMap, 0, 0, -expandedRiskMap[0][0], &lowestCosts, &shortestPath)
	fmt.Println(shortestPath)
}

func walkCave(riskMap [][]int, x int, y int, currentCost int, lowestCosts *[][]int, shortestPath *int) {
	if x > len(riskMap)-1 || y > len(riskMap[0])-1 || x < 0 || y < 0 {
		return
	}
	if x == len(riskMap)-1 && y == len(riskMap)-1 {
		cost := currentCost + riskMap[x][y]

		if cost < *shortestPath {
			*shortestPath = cost
		}
	}

	currentCost += riskMap[x][y]
	if currentCost < (*lowestCosts)[x][y] {
		(*lowestCosts)[x][y] = currentCost
	} else {
		return
	}
	if currentCost >= *shortestPath {
		return
	}

	walkCave(riskMap, x, y+1, currentCost, lowestCosts, shortestPath)
	walkCave(riskMap, x+1, y, currentCost, lowestCosts, shortestPath)
	walkCave(riskMap, x-1, y, currentCost, lowestCosts, shortestPath)
	walkCave(riskMap, x, y-1, currentCost, lowestCosts, shortestPath)
}
