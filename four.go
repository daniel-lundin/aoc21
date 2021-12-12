package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Four() {
	file, err := os.Open("./input-4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	bingoRow := scanner.Text()
	bingoNumbers := StringListToNumbers(strings.Split(bingoRow, ","))

	boards := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		boards = append(boards, StringListToNumbers(strings.Split(line, " "))...)
	}

	slowestBingo := 0
	slowestBingoIndex := -1
	bingoNumber := -1
	winningNumbers := make([]int, 25)

	for i := 0; i < len(boards)/25; i++ {
		hits := make([]int, 25)
		for numbersDrawn, number := range bingoNumbers {
			board := boards[i*25 : i*25+25]
			numberIndex := findIndex(board, number)

			if numberIndex != -1 {
				hits[numberIndex] = 1
				if isBingo(hits) {
					fmt.Printf("Bingo on board %d after %d hits, numbers drawn %d:", i, countOnes(hits), numbersDrawn)
					plotHits(board, hits)
					fmt.Println()
					if numbersDrawn > slowestBingo {
						slowestBingo = numbersDrawn
						slowestBingoIndex = i
						bingoNumber = number
						winningNumbers = hits
					}
					break

				}
			}
		}
	}

	fmt.Printf("Slowest bingo index %d, hits %d bingo number %d\n", slowestBingoIndex, slowestBingo, bingoNumber)

	unmarkedSum := sumUnmarkedNumbers(boards[slowestBingoIndex*25:slowestBingoIndex*25+25], winningNumbers)
	fmt.Printf("unmarkedSum %d, Final score:%d \n", unmarkedSum, unmarkedSum*bingoNumber)
}

func plotHits(board []int, hits []int) {
	for i := 0; i < len(hits); i++ {
		if i%5 == 0 {
			fmt.Println("")
		}
		if hits[i] == 1 {
			fmt.Printf(" x ")
		} else {
			if board[i] < 10 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", board[i])
			fmt.Printf(" ")
		}
	}
	fmt.Println("")

}

func sumUnmarkedNumbers(bingoBoard []int, markedNumbers []int) int {
	sum := 0
	for i := 0; i < len(bingoBoard); i++ {
		if markedNumbers[i] == 0 {
			sum += bingoBoard[i]
		}
	}
	return sum
}

func countOnes(hits []int) int {
	ones := 0
	for i := 0; i < len(hits); i++ {
		if hits[i] == 1 {
			ones += 1
		}

	}
	return ones
}

func isBingo(hits []int) bool {

	for i := 0; i < 20; i++ {
		// If first column, check whole row forward
		if i%5 == 0 {
			bingo := true
			for j := i; j < i+5; j++ {
				if hits[j] != 1 {
					bingo = false
				}
			}
			if bingo {
				return true
			}
		}
		// If first row, check column
		if i < 5 {
			bingo := true
			for j := 0; j < 5; j++ {
				if hits[i+j*5] != 1 {
					bingo = false
				}
			}
			if bingo {
				return true
			}
		}
	}
	return false
}

func findIndex(numbers []int, number int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == number {
			return i
		}
	}
	return -1
}
