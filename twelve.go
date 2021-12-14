package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func Twelve() {
	file, err := os.Open("./input-12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	directions := make(map[string][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		directions[parts[0]] = append(directions[parts[0]], parts[1])
		directions[parts[1]] = append(directions[parts[1]], parts[0])
	}

	visitedCaves := make(map[string]bool, 0)
	completedPaths := make([]string, 0)
	currentPath := make([]string, 0)

	walk(directions, "start", currentPath, &completedPaths, visitedCaves, false)
	fmt.Printf("Compeleted maps %d\n", len(completedPaths))

}

func walk(directions map[string][]string, position string, currentPath []string, completedPaths *[]string, visitedCaves map[string]bool, doubleVisitUsed bool) {

	if position == "start" && len(currentPath) > 0 {
		return
	}
	if position == "end" {
		completedPath := strings.Join(append(currentPath, "end"), ",")
		(*completedPaths) = append(*completedPaths, completedPath)
		return
	}

	newVisitedMap := copyVisitedMap(visitedCaves)
	if isSmallCave(position) {
		if visitedCaves[position] {
			if doubleVisitUsed {
				return
			}
			doubleVisitUsed = true
		}
		newVisitedMap[position] = true
	}

	currentPath = append(currentPath, position)

	destinations := directions[position]

	for _, destination := range destinations {
		walk(directions, destination, currentPath, completedPaths, newVisitedMap, doubleVisitUsed)

	}
}

func copyVisitedMap(visitedMap map[string]bool) map[string]bool {
	newMap := make(map[string]bool, 0)
	for k, v := range visitedMap {
		newMap[k] = v
	}
	return newMap
}

func isSmallCave(location string) bool {
	if location == "start" || location == "end" {
		return false
	}
	for _, char := range location {
		if !unicode.IsLower(char) {
			return false
		}
	}
	return true
}
