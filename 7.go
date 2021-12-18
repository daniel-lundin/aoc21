package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Seven() {
	file, err := os.Open("./input-7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	positions := StringListToNumbers(strings.Split(line, ","))
	positionMap := make(map[int]int, 0)

	for _, position := range positions {
		positionMap[position] += 1
	}

	minPosition := getMinPosition(positions)
	maxPosition := getMaxPosition(positions)
	minFuelCost := math.MaxInt
	bestPosition := math.MaxInt
	for i := minPosition; i <= maxPosition; i++ {
		fuelCost := 0
		for key, count := range positionMap {
			distance := aritmeticSum(abs(key - i))
			fuelCost += distance * count
		}
		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
			bestPosition = i
		}
	}

	fmt.Printf("Best position %d, fuel cost: %d\n", bestPosition, minFuelCost)
}

func aritmeticSum(number int) int {
	sum := 0
	for i := 0; i <= number; i++ {
		sum += i
	}
	return sum
}
func getMinPosition(positions []int) int {
	minimum := math.MaxInt
	for _, position := range positions {
		if position < minimum {
			minimum = position
		}
	}
	return minimum
}

func getMaxPosition(positions []int) int {
	maximum := math.MinInt
	for _, position := range positions {
		if position > maximum {
			maximum = position
		}
	}
	return maximum
}
