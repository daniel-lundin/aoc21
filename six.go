package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Six() {
	file, err := os.Open("./input-6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	numbers := StringListToNumbers(strings.Split(line, ","))
	lanterns := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}
	for _, number := range numbers {
		lanterns[number] += 1
	}
	fmt.Printf("inital, total lanterns %d\n", countLanterns(lanterns))

	for i := 0; i < 256; i++ {
		// fmt.Printf("day %d\n", i)
		updatedLanterns := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
			8: 0,
		}
		for j := 0; j <= 8; j++ {
			if j == 0 {
				updatedLanterns[6] += lanterns[j]
				updatedLanterns[8] += lanterns[j]
			} else {
				updatedLanterns[j-1] += lanterns[j]
			}
		}

		lanterns = updatedLanterns
	}

	fmt.Printf("final, total lanterns %d\n", countLanterns(lanterns))
	fmt.Println(lanterns)
}

func countLanterns(lanternMap map[int]int) int {
	count := 0
	for i := 0; i <= 8; i++ {
		count += lanternMap[i]
	}
	return count
}
