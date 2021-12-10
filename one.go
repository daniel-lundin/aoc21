package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func One() {
	file, err := os.Open("./input-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lastMeasurement := math.MaxInt32
	increments := 0

	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Failed to parse int")
			os.Exit(1)
		}

		if measurement > lastMeasurement {
			increments = increments + 1
		}
		lastMeasurement = measurement
	}

	fmt.Printf("Found %d increments", increments)

}
