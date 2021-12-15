package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Thirteen() {
	file, err := os.Open("./input-13-example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dots := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ",")
		dots = append(dots, StringListToNumbers(parts))
	}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		foldCoords := strings.Split(parts[2], "=")
		fmt.Printf("%s %s\n", foldCoords[0], foldCoords[1])
	}

	fmt.Println(dots)

}
