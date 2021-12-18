package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func OnePartTwo() {
	file, err := os.Open("./input-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	increments := 0
	var lastValues [4]int
	lastValues[0] = math.MaxInt32
	lastValues[1] = math.MaxInt32
	lastValues[2] = math.MaxInt32
	lastValues[3] = math.MaxInt32
	windowIndex := 0

	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Failed to parse int")
			os.Exit(1)
		}

		lastValues[windowIndex%4] = measurement
		if hasIncrement(lastValues, windowIndex) {
			increments = increments + 1
		}
		windowIndex = (windowIndex + 1) % 4
	}

	fmt.Printf("Found %d increments\n", increments)
}

func prevIndex(index int) int {
	return (index - 1 + 4) % 4
}

func hasIncrement(lastValues [4]int, index int) bool {
	currentWindow := lastValues[index] + lastValues[prevIndex(index)] + lastValues[prevIndex(prevIndex(index))]
	previousWindow := lastValues[prevIndex(index)] + lastValues[prevIndex(prevIndex(index))] + lastValues[prevIndex(prevIndex(prevIndex(index)))]
	return currentWindow > previousWindow

}

// 199  A
// 200  A B
// 208  A B C
// 210    B C D
// 200  E   C D
// 207  E F   D
// 240  E F G
// 269    F G H
// 260  I   G H
// 263  I J   H
// 268  I J K
