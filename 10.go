package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Ten() {
	file, err := os.Open("./input-10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	openings := map[string]bool{
		"<": true,
		"(": true,
		"[": true,
		"{": true,
	}
	closings := map[string]bool{
		">": true,
		")": true,
		"]": true,
		"}": true,
	}
	pairs := map[string]string{
		"<": ">",
		"(": ")",
		"[": "]",
		"{": "}",
	}
	errorScores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	completionScoreMap := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	syntaxErrorScore := 0
	completionScores := make([]int, 0)
	for _, line := range lines {

		stack := make([]string, 0)
		correctLine := true
		illegalChar := ""
		for _, char := range strings.Split(line, "") {
			if openings[char] {
				stack = append(stack, char)
			}
			if closings[char] {
				top := stack[len(stack)-1]
				if pairs[top] != char {
					//fmt.Printf("%s - Expected %s, but found %s instead(at line %d, char %d)\n", line, pairs[top], char, lineNo, pos)
					illegalChar = char
					correctLine = false
					break
				} else {
					stack = stack[0 : len(stack)-1]
				}
			}
		}
		if correctLine == false {
			syntaxErrorScore += errorScores[illegalChar]
		}
		// Incomplete line
		if correctLine && len(stack) != 0 {
			completionScore := 0
			completionSequence := make([]string, 0)
			for i := len(stack) - 1; i >= 0; i-- {
				completionScore *= 5
				completionScore += completionScoreMap[pairs[stack[i]]]
				completionSequence = append(completionSequence, pairs[stack[i]])
			}
			completionScores = append(completionScores, completionScore)
			// fmt.Printf("Completing line %s with %v, score: %d\n", line, completionSequence, completionScore)
		}
	}
	sort.Ints(completionScores)
	fmt.Printf("Syntax error score: %d\n", syntaxErrorScore)
	fmt.Printf("Completion score: %d\n", completionScores[len(completionScores)/2])

}
