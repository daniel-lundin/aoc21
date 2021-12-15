package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Fourteen() {
	file, err := os.Open("./input-14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()

	insertionRules := make(map[string]byte, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		insertionRules[parts[0]] = parts[1][0]
	}

	counts := make(map[byte]int, 0)
	for i := 0; i < len(polymer); i++ {
		counts[polymer[i]]++
	}
	pairs := make([]string, 0)
	for i := 0; i < len(polymer)-1; i++ {
		pairs = append(pairs, polymer[i:i+2])
	}

	for _, pair := range pairs {
		pairInsertion(pair, insertionRules, counts, 0, 10)
	}

	leastCommon := math.MaxInt
	mostCommon := 0
	for _, value := range counts {
		leastCommon = minInt(value, leastCommon)
		mostCommon = maxInt(value, mostCommon)
	}
	fmt.Printf("least common %d\n", leastCommon)
	fmt.Printf("most common %d\n", mostCommon)
	fmt.Println(mostCommon - leastCommon)
}

func sumPolymer(counts map[byte]int) int {
	count := 0
	for _, value := range counts {
		count += value
	}
	return count
}

func pairInsertion(pair string, insertionRules map[string]byte, counts map[byte]int, step int, steps int) {
	if step >= steps {
		return
	}
	insertion := insertionRules[pair]
	counts[insertion]++
	pairInsertion(fmt.Sprintf("%c%c", pair[0], insertion), insertionRules, counts, step+1, steps)
	pairInsertion(fmt.Sprintf("%c%c", insertion, pair[1]), insertionRules, counts, step+1, steps)
}
