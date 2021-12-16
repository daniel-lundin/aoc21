package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Pair struct {
	a byte
	b byte
}

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

	insertionRules := make(map[Pair]byte, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		insertionRules[Pair{parts[0][0], parts[0][1]}] = parts[1][0]
	}

	counts := make(map[byte]int, 0)
	for i := 0; i < len(polymer); i++ {
		counts[polymer[i]]++
	}
	pairs := make([]Pair, 0)
	for i := 0; i < len(polymer)-1; i++ {
		pairs = append(pairs, Pair{polymer[i], polymer[i+1]})
	}

	localCounts := make(map[Pair]map[byte]int, 0)
	emptyCounts := make(map[Pair]map[byte]int, 0)
	for key, _ := range insertionRules {
		localCounts[key] = make(map[byte]int, 0)
		pairInsertion(key, insertionRules, localCounts[key], emptyCounts, 0, 20)
	}

	for _, pair := range pairs {
		pairInsertion(pair, insertionRules, counts, localCounts, 0, 20)
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

func pairInsertion(pair Pair, insertionRules map[Pair]byte, counts map[byte]int, localCounts map[Pair]map[byte]int, step int, steps int) {
	if step >= steps {
		for key, count := range localCounts[pair] {
			counts[key] += count
		}
		return
	}
	insertion := insertionRules[pair]
	counts[insertion]++
	pairInsertion(Pair{pair.a, insertion}, insertionRules, counts, localCounts, step+1, steps)
	pairInsertion(Pair{insertion, pair.b}, insertionRules, counts, localCounts, step+1, steps)
}
