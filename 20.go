package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Twenty() {
	file, err := os.Open("./input-20.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	algoString := scanner.Text()
	scanner.Scan()

	image := make(map[Coord2D]bool)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		pixels := strings.Split(line, "")
		for index, pixel := range pixels {
			if pixel == "#" {
				image[Coord2D{index, row}] = true
			}
		}

		row++
	}

	plotImage(image, false)

	output := image
	infinitePixelValue := false
	for rounds := 0; rounds < 50; rounds++ {
		newOutput, newInfiniteValue := applyAlgorithm(output, infinitePixelValue, algoString)
		output = newOutput
		infinitePixelValue = newInfiniteValue
	}

	fmt.Printf("pixels %d\n", len(output))
}

func applyAlgorithm(image map[Coord2D]bool, infinitePixelValue bool, algoString string) (map[Coord2D]bool, bool) {
	output := make(map[Coord2D]bool)
	minX, maxX, minY, maxY := imgDimensions(image)

	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			bitList := make([]int, 0)
			for j := -1; j < 2; j++ {
				for i := -1; i < 2; i++ {
					if x+i < minX || x+i > maxX || y+j < minY || y+j > maxY {
						if infinitePixelValue {
							bitList = append(bitList, 1)
						} else {
							bitList = append(bitList, 0)
						}
					} else if image[Coord2D{x + i, y + j}] {
						bitList = append(bitList, 1)
					} else {
						bitList = append(bitList, 0)
					}
				}
			}
			pixel := algoString[bitsToDecimal(bitList)]
			if pixel == '#' {
				output[Coord2D{x, y}] = true
			}

		}
	}
	var value byte
	if infinitePixelValue {
		value = algoString[bitsToDecimal([]int{1, 1, 1, 1, 1, 1, 1, 1, 1})]
	} else {
		value = algoString[0]
	}

	return output, value == '#'
}

func plotImage(image map[Coord2D]bool, infinitePixelValue bool) {
	minX, maxX, minY, maxY := imgDimensions(image)

	fmt.Printf("Width height: %d %d, image:\n", maxX-minX, maxY-minY)

	padding := 5
	for y := minY - padding; y <= maxY+padding; y++ {
		for x := minX - padding; x <= maxX+padding; x++ {
			if x < minX || x > maxX || y < minY || y > maxY {
				if infinitePixelValue {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
				continue
			}
			if image[Coord2D{x, y}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func imgDimensions(image map[Coord2D]bool) (int, int, int, int) {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for coord, _ := range image {
		minX = minInt(minX, coord.x)
		maxX = maxInt(maxX, coord.x)
		minY = minInt(minY, coord.y)
		maxY = maxInt(maxY, coord.y)
	}
	return minX, maxX, minY, maxY
}
