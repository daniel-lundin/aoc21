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
		riskMap = append(riskMap, StringListToNumbers(strings.Split(line, "")))
	}

	shortestPath := 830
	lowestCosts := make([][]int, len(riskMap))
	for rowIndex, row := range riskMap {
		lowestCosts[rowIndex] = make([]int, len(row))
		for colIndex := range row {
			lowestCosts[rowIndex][colIndex] = math.MaxInt
		}
	}
	walkCave(riskMap, 0, 0, -riskMap[0][0], &lowestCosts, &shortestPath)
	fmt.Println(shortestPath)
}

func walkCave(riskMap [][]int, x int, y int, currentCost int, lowestCosts *[][]int, shortestPath *int) {
	if x > len(riskMap)-1 || y > len(riskMap[0])-1 {
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
		// fmt.Printf(".")
		return
	}

	if x+1 < len(riskMap)-1 && y < len(riskMap[0])-1 {
		if riskMap[x][y+1] < riskMap[x+1][y] {

			walkCave(riskMap, x, y+1, currentCost, lowestCosts, shortestPath)
			walkCave(riskMap, x+1, y, currentCost, lowestCosts, shortestPath)
		} else {
			walkCave(riskMap, x+1, y, currentCost, lowestCosts, shortestPath)
			walkCave(riskMap, x, y+1, currentCost, lowestCosts, shortestPath)
		}
	} else {
		walkCave(riskMap, x, y+1, currentCost, lowestCosts, shortestPath)
		walkCave(riskMap, x+1, y, currentCost, lowestCosts, shortestPath)
	}
}
