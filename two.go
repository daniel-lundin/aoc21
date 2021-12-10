package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func down(amount int, aim *int) {
	*aim += amount
}
func up(amount int, aim *int) {
	*aim -= amount
}
func forward(amount int, horizontalPosition *int, depth *int, aim int) {
	*horizontalPosition += amount
	*depth += amount * aim
}

func Two() {
	file, err := os.Open("./input-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	depth := 0
	horizontalPosition := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		stringSlice := strings.Split(line, " ")
		command := stringSlice[0]
		argument, err := strconv.Atoi(stringSlice[1])
		if err != nil {
			fmt.Printf("Failed to parse command: %s \n", line)
			os.Exit(1)
		}
		if command == "forward" {
			forward(argument, &horizontalPosition, &depth, aim)
		}
		if command == "up" {
			up(argument, &aim)
		}
		if command == "down" {
			down(argument, &aim)
		}

	}

	fmt.Printf("Depth: %d Horizontal position: %d, result: %d\n", depth, horizontalPosition, depth*horizontalPosition)
}
