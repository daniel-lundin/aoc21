package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dot struct {
	x int
	y int
}

type Fold struct {
	axis  string
	value int
}

func Thirteen() {
	file, err := os.Open("./input-13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dots := make(map[Dot]bool, 0)
	width := 0
	height := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := StringListToNumbers(strings.Split(line, ","))
		dots[Dot{parts[0], parts[1]}] = true
		width = maxInt(parts[0], width)
		height = maxInt(parts[1], height)
	}

	folds := make([]Fold, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		foldParts := strings.Split(parts[2], "=")
		number, err := strconv.Atoi(foldParts[1])
		if err != nil {
			fmt.Println("Error parsing int")
			os.Exit(1)
		}
		folds = append(folds, Fold{foldParts[0], number})
	}

	for _, fold := range folds {
		updatedDots := make(map[Dot]bool, 0)
		if fold.axis == "x" {
			for dot, _ := range dots {
				if dot.x < fold.value {
					updatedDots[dot] = true
				} else {
					updatedDot := Dot{fold.value - (dot.x - fold.value), dot.y}
					updatedDots[updatedDot] = true
				}
			}
			width = width - (width - fold.value) - 1
		} else {
			for dot, _ := range dots {
				if dot.y < fold.value {
					updatedDots[dot] = true
				} else {
					updatedDot := Dot{dot.x, fold.value - (dot.y - fold.value)}
					updatedDots[updatedDot] = true
				}
			}
			height = height - (height - fold.value) - 1
		}

		dots = updatedDots
	}
	fmt.Printf("Completed %d %d\n", width, height)
	plotDots(dots, width, height)
	fmt.Println(len(dots))
}

func plotDots(dots map[Dot]bool, width int, height int) {
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if dots[Dot{x, y}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
