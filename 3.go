package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Three() {
	file, err := os.Open("./input-3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers [][]int

	for scanner.Scan() {
		line := scanner.Text()
		bits := make([]int, 0)

		chars := strings.Split(line, "")
		for i := 0; i < len(chars); i++ {
			number, error := strconv.Atoi(chars[i])
			if error != nil {
				os.Exit(1)
			}
			bits = append(bits, number)
		}
		numbers = append(numbers, bits)
	}

	var oneBits []int
	for i := 0; i < len(numbers[0]); i++ {
		oneBits = append(oneBits, 0)
	}

	for i := 0; i < len(numbers); i++ {
		bitString := numbers[i]
		for j := 0; j < len(bitString); j++ {
			if bitString[j] == 1 {
				oneBits[j] += 1
			}
		}
	}

	rowWidth := len(oneBits)
	var gammaBits []int
	var epsilonBits []int
	for i := 0; i < rowWidth; i++ {
		if oneBits[i] >= len(numbers)/2 {
			gammaBits = append(gammaBits, 1)
			epsilonBits = append(epsilonBits, 0)
		} else {
			gammaBits = append(gammaBits, 0)
			epsilonBits = append(epsilonBits, 1)
		}
	}

	gamma := binToHex(gammaBits)
	epsilon := binToHex(epsilonBits)

	fmt.Printf("gamma %d epsilon %d\n", gamma, epsilon)
	fmt.Printf("Power consumption: %d\n", gamma*epsilon)
	oxygenGeneratorRating := lifeSupportRating(numbers, rowWidth, mostCommonValue)
	CO2ScrubberRating := lifeSupportRating(numbers, rowWidth, leastCommonValue)
	fmt.Printf("oxygen generator rating: %d \n", oxygenGeneratorRating)
	fmt.Printf("co2 scrubber rating : %d \n", CO2ScrubberRating)
	fmt.Printf("life support rating: %d \n", oxygenGeneratorRating*CO2ScrubberRating)
}

func lifeSupportRating(numbers [][]int, rowWidth int, bitCriteria func([][]int, int) int) int {
	var matches [][]int
	for i := 0; i < len(numbers); i++ {
		matches = append(matches, numbers[i])
	}

	for i := 0; i < rowWidth; i++ {
		localMatches := make([][]int, 0)

		bit := bitCriteria(matches, i)

		for _, match := range matches {
			if match[i] == bit {
				localMatches = append(localMatches, match)
			}
		}
		matches = localMatches

		if len(matches) == 1 {
			break
		}

	}
	return binToHex(matches[0])
}

func mostCommonValue(numbers [][]int, index int) int {
	ones := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i][index] == 1 {
			ones += 1
		}
	}
	if ones >= len(numbers)-ones {
		return 1
	}
	return 0
}

func leastCommonValue(numbers [][]int, index int) int {
	ones := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i][index] == 1 {
			ones += 1
		}
	}
	if ones >= len(numbers)-ones {
		return 0
	}
	return 1
}

func binToHex(bits []int) int {
	pow := 1
	result := 0
	for i := len(bits) - 1; i >= 0; i-- {
		result += bits[i] * pow
		pow = pow << 1
	}
	return result
}
