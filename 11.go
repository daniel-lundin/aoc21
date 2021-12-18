package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Eleven() {
	file, err := os.Open("./input-11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	octoEnergies := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		octoEnergies = append(octoEnergies, StringListToNumbers(strings.Split(line, "")))
	}

	totalFlashes := 0
	for i := 0; i < 1000; i++ {
		for row, octoRow := range octoEnergies {
			for col, _ := range octoRow {
				octoEnergies[row][col]++

			}

		}

		flash := true
		for flash {
			flashes := 0

			for row, octoRow := range octoEnergies {
				for col, octo := range octoRow {
					if octo > 9 {
						flashes += 1
						octoEnergies[row][col] = 0
						increaseAdjacentEnergies(&octoEnergies, row, col)
					}
				}
			}
			if flashes == 0 {
				flash = false
			} else {
				totalFlashes += flashes
			}
		}
		allZeros := true
		for _, octoRow := range octoEnergies {
			for _, octo := range octoRow {
				if octo != 0 {
					allZeros = false
					break
				}
			}
		}
		if allZeros {
			fmt.Printf("All zeroes at step %d\n", i+1)
			break
		}
		printOctoEnergy(octoEnergies)
	}
	fmt.Println(totalFlashes)
}

func increaseAdjacentEnergies(octoEnergies *[][]int, row int, col int) {
	fromRow := maxInt(0, row-1)
	toRow := minInt(len(*octoEnergies)-1, row+1)
	fromCol := maxInt(0, col-1)
	toCol := minInt(len((*octoEnergies)[0])-1, col+1)
	for i := fromRow; i <= toRow; i++ {
		for j := fromCol; j <= toCol; j++ {
			if (*octoEnergies)[i][j] > 0 {
				(*octoEnergies)[i][j]++
			}
		}
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func printOctoEnergy(octos [][]int) {
	for _, octoRow := range octos {
		for _, octo := range octoRow {
			fmt.Printf("%d", octo)
		}
		fmt.Println()

	}
}
