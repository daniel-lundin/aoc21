package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//    0000
//   1    2
//   1    2
//    3333
//   4    5
//   4    5
//    6666
type digit struct {
	value        int
	segmentCount int
	positions    []int
}

func Eight() {
	digits := []digit{
		{0, 6, []int{0, 1, 2, 4, 5, 6}},
		{1, 2, []int{2, 5}},
		{2, 5, []int{0, 2, 3, 4, 6}},
		{3, 5, []int{0, 2, 3, 5, 6}},
		{4, 4, []int{1, 2, 3, 5}},
		{5, 5, []int{0, 1, 3, 5, 6}},
		{6, 6, []int{0, 1, 3, 4, 5, 6}},
		{7, 3, []int{0, 2, 5}},
		{8, 7, []int{0, 1, 2, 3, 4, 5, 6}},
		{9, 6, []int{0, 1, 2, 3, 5, 6}},
	}
	file, err := os.Open("./input-8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := make([][]string, 0)
	outputs := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// line := "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"
		//	line := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
		splits := strings.Split(line, " | ")
		inputs = append(inputs, strings.Split(splits[0], " "))
		outputs = append(outputs, strings.Split(splits[1], " "))
	}

	decodedOutputValues := 0

	for entryIndex, input := range inputs {
		segmentCandidates := [][]string{
			{"a", "b", "c", "d", "e", "f", "g"},
			{"a", "b", "c", "d", "e", "f", "g"},
			{"a", "b", "c", "d", "e", "f", "g"},
			{"a", "b", "c", "d", "e", "f", "g"},
			{"a", "b", "c", "d", "e", "f", "g"},
			{"a", "b", "c", "d", "e", "f", "g"},
			{"a", "b", "c", "d", "e", "f", "g"},
		}

		for _, outputDigit := range input {
			hits := digitsWithSegmentCount(digits, len(outputDigit))
			characters := strings.Split(outputDigit, "")
			if len(hits) == 1 {

				for index, segmentCandidate := range segmentCandidates {
					if intListIncludes(hits[0].positions, index) == false {
						// remove from candidate
						segmentCandidates[index] = characterSubtraction(segmentCandidate, characters)
					}
				}

				for _, position := range hits[0].positions {
					segmentCandidates[position] = characterIntersection(characters, segmentCandidates[position])

				}
			}
		}
		for _, outputDigit := range input {
			characters := strings.Split(outputDigit, "")
			hits := digitsWithSegmentCount(digits, len(outputDigit))
			if len(hits) != 1 {
				characters = characterSubtraction(characters, segmentCandidates[0])
				var validHit digit
				for _, hit := range hits {
					candidatesForFit := make([][]string, 7)
					for _, position := range hit.positions {
						if position == 0 {
							continue
						}
						candidatesForFit[position] = characterIntersection(characters, segmentCandidates[position])

					}

					// is valid fit?
					singleCandidateMap := make(map[string]int, 0)
					candidateSegments := 0
					for _, candidate := range candidatesForFit {
						if len(candidate) == 1 {
							singleCandidateMap[candidate[0]]++
						}
						if len(candidate) > 0 {
							candidateSegments++
						}
					}
					doubleSingles := false
					for _, value := range singleCandidateMap {
						if value > 1 {
							doubleSingles = true

						}
					}

					if doubleSingles == false && candidateSegments == len(characters) {
						validHit = hit
					}

				}
				for index, segmentCandidate := range segmentCandidates {
					if intListIncludes(validHit.positions, index) == false {
						// remove from candidate
						segmentCandidates[index] = characterSubtraction(segmentCandidate, characters)
					}
				}
				for _, position := range validHit.positions {
					if position == 0 {
						continue
					}
					segmentCandidates[position] = characterIntersection(characters, segmentCandidates[position])

				}

			}
		}
		solution := make([]string, 0)
		for _, candidates := range segmentCandidates {
			solution = append(solution, candidates[0])
		}

		outputValue := 1000*decodeNumber(digits, solution, outputs[entryIndex][0]) + 100*decodeNumber(digits, solution, outputs[entryIndex][1]) + 10*decodeNumber(digits, solution, outputs[entryIndex][2]) +
			decodeNumber(digits, solution, outputs[entryIndex][3])

		decodedOutputValues += outputValue
	}

	fmt.Printf("summed output values %d\n", decodedOutputValues)
}

func decodeNumber(digits []digit, wireMapping []string, characters string) int {

	segmentsLit := make([]int, 0)
	for _, char := range strings.Split(characters, "") {
		index := findIndexStr(wireMapping, char)
		segmentsLit = append(segmentsLit, index)
	}

	for _, digit := range digits {
		if digit.segmentCount != len(characters) {
			continue
		}
		allMatch := true
		for _, segment := range segmentsLit {
			if findIndexInt(digit.positions, segment) == -1 {
				allMatch = false
				break
			}
		}
		if allMatch {
			return digit.value
		}
	}
	return -1
}

func intListIncludes(list []int, search int) bool {
	for _, number := range list {
		if number == search {
			return true
		}
	}
	return false
}

func characterSubtraction(a []string, b []string) []string {
	keepers := make([]string, 0)
	for _, char := range a {
		if findIndexStr(b, char) == -1 {
			keepers = append(keepers, char)
		}
	}
	return keepers
}

func characterIntersection(a []string, b []string) []string {
	intersection := make([]string, 0)
	for _, char := range a {
		if findIndexStr(b, char) != -1 {
			intersection = append(intersection, char)
		}
	}
	return intersection
}

func findIndexInt(a []int, x int) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
func findIndexStr(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func digitsWithSegmentCount(digits []digit, segmentCount int) []digit {
	hits := make([]digit, 0)
	for _, digit := range digits {
		if digit.segmentCount == segmentCount {
			hits = append(hits, digit)
		}
	}
	return hits
}

// [[d] [c g] [e b] [c g] [f a] [e b] [f a]]
// [[d] [e f] [a b] [e f] [c g] [a b] [c g]]
//    dddd
//   e    a
//   f    b
//    efef
//   c    a
//   g    b
//    cgcg
//
// [{2 5 [0 2 3 4 6]} {3 5 [0 2 3 5 6]} {5 5 [0 1 3 5 6]}]
// [c d f b e]
// 2, 3, 5
// Remove D:
// [c f b e]
//
// __a_b 7   8       ecagb 1  abcdeg 4    cafbge fdbac fegbdc | fgae cfgab fg bagce
// gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
//     :
//    cccc
//   e    f
//   e    f
//    aaaa
//   d    g
//   d    g
//    bbbb

//     1:
//    eeee
//   ,    c/a
//   ,    c/a
//    ....
//   .    c/a
//   .    c/a
//    ....
