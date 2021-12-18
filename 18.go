package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type SnailfishNumber struct {
	left     *SnailfishNumber
	right    *SnailfishNumber
	parent   *SnailfishNumber
	value    int
	isNumber bool
}

func Eighteen() {
	file, err := os.Open("./input-18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numberList := make([]*SnailfishNumber, 0)
	for scanner.Scan() {
		line := scanner.Text()
		numberList = append(numberList, parseSnailfish(line, nil))
	}

	// Part 1
	// for len(numberList) > 1 {
	// 	first := numberList[0]
	// 	second := numberList[1]
	// 	result := addSnailFishes(first, second)
	// 	numberList[1] = result
	// 	numberList = numberList[1:]
	// }
	// fmt.Printf("Final number%v\n", printSnailFish(*numberList[0]))
	// fmt.Printf("Magnitude %d\n", magnitude(numberList[0]))
	maxMagnitude := 0
	for _, number1 := range numberList {
		for _, number2 := range numberList {

			result := addSnailFishes(cloneSnailFish(number1, nil), cloneSnailFish(number2, nil))
			maxMagnitude = maxInt(maxMagnitude, magnitude(result))
			fmt.Printf("magnitude %d\n", magnitude(result))
		}
	}
	fmt.Printf("max  %d\n", maxMagnitude)
}

func cloneSnailFish(fish, parent *SnailfishNumber) *SnailfishNumber {
	if fish == nil {
		return nil
	}

	clone := new(SnailfishNumber)
	clone.left = cloneSnailFish(fish.left, clone)
	clone.right = cloneSnailFish(fish.right, clone)
	clone.isNumber = fish.isNumber
	clone.parent = parent
	clone.value = fish.value
	return clone
}
func magnitude(snailfish *SnailfishNumber) int {
	if snailfish.isNumber {
		return snailfish.value
	}

	return magnitude(snailfish.left)*3 + 2*magnitude(snailfish.right)
}

func addSnailFishes(first, second *SnailfishNumber) *SnailfishNumber {
	newFish := new(SnailfishNumber)
	newFish.left = first
	newFish.right = second
	newFish.isNumber = false
	first.parent = newFish
	second.parent = newFish
	reduceNumber(newFish)
	return newFish
}

func reduceNumber(snailfish *SnailfishNumber) {
	for true {
		for explodeNumber(snailfish, 0) {
		}

		if splitNumber(snailfish) {
			continue
		}
		break
	}
}

func makeSnailfishLiteral(value int, parent *SnailfishNumber) *SnailfishNumber {
	snailfish := new(SnailfishNumber)
	snailfish.isNumber = true
	snailfish.value = value
	snailfish.parent = parent
	return snailfish
}

func splitNumber(snailfish *SnailfishNumber) bool {
	if snailfish.isNumber {
		if snailfish.value >= 10 {
			snailfish.isNumber = false
			snailfish.left = makeSnailfishLiteral(snailfish.value/2, snailfish)
			snailfish.right = makeSnailfishLiteral(int(math.Ceil(float64(snailfish.value)/2.0)), snailfish)
			return true
		}
		return false
	}

	if splitNumber(snailfish.left) {
		return true
	}
	if splitNumber(snailfish.right) {
		return true
	}
	return false

}
func explodeNumber(snailfish *SnailfishNumber, level int) bool {
	if level > 4 || snailfish == nil {
		return false
	}
	if level == 4 && !snailfish.isNumber {
		explodePair(snailfish)
		return true
	}
	if !explodeNumber(snailfish.left, level+1) {
		if !explodeNumber(snailfish.right, level+1) {
			return false
		}
	}
	return true
}

func explodePair(fish *SnailfishNumber) {
	if fish.isNumber {
		return
	}
	addToNeighborLeft(fish, fish.left.value)
	addToNeighborRight(fish, fish.right.value)
	fish.value = 0
	fish.isNumber = true
}

func addToNeighborLeft(fish *SnailfishNumber, value int) {
	if fish.parent != nil {
		if fish == fish.parent.left {
			addToNeighborLeft(fish.parent, value)
		}
		if fish == fish.parent.right {
			addToRightMostChild(fish.parent.left, value)
		}
	}

}
func addToNeighborRight(fish *SnailfishNumber, value int) {
	if fish.parent != nil {
		if fish == fish.parent.right {
			addToNeighborRight(fish.parent, value)
		}
		if fish == fish.parent.left {
			addToLeftMostChild(fish.parent.right, value)
		}
	}

}

func addToRightMostChild(fish *SnailfishNumber, value int) {
	if fish.isNumber {
		fish.value += value
		return
	}
	addToRightMostChild(fish.right, value)

}
func addToLeftMostChild(fish *SnailfishNumber, value int) {
	if fish.isNumber {
		fish.value += value
		return
	}
	addToLeftMostChild(fish.left, value)

}

func printSnailFish(snailfish SnailfishNumber) string {
	if snailfish.isNumber {
		return fmt.Sprintf("%d", snailfish.value)
	}
	return "[" + printSnailFish(*snailfish.left) + "," + printSnailFish(*snailfish.right) + "]"

}

func parseSnailfish(line string, parent *SnailfishNumber) *SnailfishNumber {
	token := line[0]
	if token == '[' {
		level := 0
		endBracket := -1
		commaIndex := -1
		for i := 1; i < len(line); i++ {
			char := line[i]
			if char == '[' {
				level++
			}
			if char == ',' && level == 0 {
				commaIndex = i
			}
			if char == ']' {
				level--
				if level == -1 {
					endBracket = i
					break
				}
			}
		}
		snailfish := new(SnailfishNumber)
		snailfish.left = parseSnailfish(line[1:commaIndex], snailfish)
		snailfish.right = parseSnailfish(line[commaIndex+1:endBracket], snailfish)
		snailfish.parent = parent
		snailfish.isNumber = false
		return snailfish
	}
	number, err := strconv.Atoi(string(token))
	if err != nil {
		fmt.Printf("Failed to parse number at 0 %s\n", line)
		os.Exit(1)
	}
	snailfish := new(SnailfishNumber)
	snailfish.value = number
	snailfish.isNumber = true
	snailfish.parent = parent
	return snailfish

}
