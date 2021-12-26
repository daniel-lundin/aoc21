package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op string
	a  string
	b  string
}

func TwentyFour() {
	file, err := os.Open("./input-24.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		instruction := parseInstruction(line)
		instructions = append(instructions, instruction)
	}

	// registers := make(map[string]int)

	permutations := make([][]int, 0, 14*9)
	faultyNumbers := make(map[Coord2D]bool)
	makePermutations([]int{}, &permutations, 14, faultyNumbers, func(number []int) {

		// time.Sleep(1 * time.Second)
		startNumber := number
		registers, err := runProgram(instructions, &number)
		if err != nil {
			lastIndex := len(startNumber) - len(number) - 1
			faultyNumber := startNumber[lastIndex]
			fmt.Printf("start len %d, current len %d. Number %d\n", len(startNumber), len(number), faultyNumber)
			fmt.Printf("Divide by zero at %d \n", lastIndex)

			faultyNumbers[Coord2D{lastIndex, faultyNumber}] = true
			return
		}
		if registers["z"] == 0 {
			fmt.Printf("found valid number %v\n", startNumber)
		}
	})
}

func makePermutations(combination []int, output *[][]int, length int, faultyNumbers map[Coord2D]bool, callback func(number []int)) {
	for j := 1; j <= 9; j++ {
		skipNumber := false
		for faulty := range faultyNumbers {
			if len(combination) > faulty.x && combination[faulty.x] == faulty.y {
				skipNumber = true
			}
		}
		if skipNumber {
			// fmt.Printf("bailing out for %v\n", combination)
			continue
		}
		if len(combination) == length-1 {
			// *output = append(*output, append(combination, j))
			callback(append(combination, j))
		} else {
			makePermutations(append(combination, j), output, length, faultyNumbers, callback)
		}
	}
}

// inp a - Read an input value and write it to variable a.
// add a b - Add the value of a to the value of b, then store the result in variable a.
// mul a b - Multiply the value of a by the value of b, then store the result in variable a.
// div a b - Divide the value of a by the value of b, truncate the result to an integer, then store the result in variable a. (Here, "truncate" means to round the value toward zero.)
// mod a b - Divide the value of a by the value of b, then store the remainder in variable a. (This is also called the modulo operation.)
// eql a b - If the value of a and b are equal, then store the value 1 in variable a. Otherwise, store the value 0 in variable a.

func runInstruction(ins Instruction, input *[]int, registers map[string]int) error {
	if ins.op == "inp" {
		registers[ins.a] = (*input)[0]
		*input = (*input)[1:]
	}
	if ins.op == "add" {
		registers[ins.a] = registers[ins.a] + bOperand(ins, registers)
	}
	if ins.op == "mul" {
		registers[ins.a] = registers[ins.a] * bOperand(ins, registers)
	}
	if ins.op == "div" {
		b := bOperand(ins, registers)
		if b == 0 {
			return errors.New("Divide by zero")
		}
		registers[ins.a] = registers[ins.a] / b
	}
	if ins.op == "mod" {
		b := bOperand(ins, registers)
		if b == 0 || registers[ins.a] == 0 {
			return errors.New("Mod by zero")
		}
		registers[ins.a] = registers[ins.a] % b
	}
	if ins.op == "eql" {
		if registers[ins.a] == bOperand(ins, registers) {
			registers[ins.a] = 1
		} else {
			registers[ins.a] = 0
		}
	}
	return nil
}

func bOperand(ins Instruction, registers map[string]int) int {
	number, err := strconv.Atoi(ins.b)
	if err != nil {
		return registers[ins.b]
	} else {
		return number
	}
}

func parseInstruction(line string) Instruction {
	parts := strings.Split(line, " ")
	if len(parts) == 3 {
		return Instruction{
			parts[0],
			parts[1],
			parts[2],
		}
	} else {
		return Instruction{
			parts[0],
			parts[1],
			"",
		}
	}
}

func runProgram(instructions []Instruction, input *[]int) (map[string]int, error) {
	registers := make(map[string]int)
	for _, ins := range instructions {
		err := runInstruction(ins, input, registers)

		if err != nil {
			return registers, err
		}
	}
	return registers, nil
}
