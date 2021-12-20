package main

import (
	"fmt"
	"os"
	"strconv"
)

func StringListToNumbers(args []string) []int {
	numbers := make([]int, 0)
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			continue
		}
		parsedNumber, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println("Failed parsing integer")
			fmt.Println(args[i])
			os.Exit(1)
		}

		numbers = append(numbers, parsedNumber)
	}
	return numbers
}

func bitsToDecimal(bits []int) int {
	sum := 0
	pow := 2 << (len(bits) - 2)
	for _, bit := range bits {
		sum += bit * pow
		pow = pow >> 1
	}
	return sum
}
