package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op string
	a  string
	b  string
}

type InputPair struct {
	z     int
	input int
}

func TwentyFour() {
	file, err := os.Open("./input-24.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := make([]Instruction, 0)
	subPrograms := make([][]Instruction, 0)
	subProgram := make([]Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		instruction := parseInstruction(line)
		if instruction.op == "inp" {
			if len(subProgram) > 0 {
				subPrograms = append([][]Instruction{subProgram}, subPrograms...)
			}
			subProgram = make([]Instruction, 0)
		}
		instructions = append(instructions, instruction)
		subProgram = append(subProgram, instruction)
	}
	subPrograms = append([][]Instruction{subProgram}, subPrograms...)

	ioMaps := make([]map[int][]InputPair, 0)
	for _, program := range subPrograms[0 : len(subPrograms)-1] {
		ioMap := make(map[int][]InputPair)
		for z := -100000; z < 100000; z++ {
			for i := 1; i < 10; i++ {
				registers := map[string]int{
					"z": z,
				}
				input := []int{i}
				registers, _ = runProgram(program, registers, &input)
				ioMap[registers["z"]] = append(ioMap[registers["z"]], InputPair{z, i})

			}
		}
		ioMaps = append(ioMaps, ioMap)
	}
	ioMap := make(map[int][]InputPair)
	for i := 1; i < 10; i++ {
		registers := map[string]int{}
		input := []int{i}
		registers, _ = runProgram(subPrograms[len(subPrograms)-1], registers, &input)
		ioMap[registers["z"]] = append(ioMap[registers["z"]], InputPair{0, i})

	}
	ioMaps = append(ioMaps, ioMap)

	completedSequences := make([]int, 0)
	findInputs(ioMaps, 0, 0, &completedSequences, 0)

	smallestNumber := math.MaxInt
	for _, number := range completedSequences {
		smallestNumber = minInt(number, smallestNumber)
	}
	fmt.Println(smallestNumber)
}

func findInputs(ioMaps []map[int][]InputPair, currentNumber int, validOutput int, completedNumbers *[]int, depth int) {
	if len(ioMaps) == 0 {
		*completedNumbers = append(*completedNumbers, currentNumber)
		return
	}
	ioMap := ioMaps[0]
	usedNumbers := make(map[int]bool)
	for _, inputPair := range ioMap[validOutput] {
		if len(ioMaps) == 2 {
			if usedNumbers[inputPair.input] {
				continue
			}
			usedNumbers[inputPair.input] = true
		}
		number := currentNumber + inputPair.input*int(math.Pow10(depth))
		findInputs(ioMaps[1:], number, inputPair.z, completedNumbers, depth+1)
	}
}
func makePermutations(combination []int, length int, callback func(number []int)) {
	for j := 9; j >= 1; j-- {
		if len(combination) == length-1 {
			callback(append(combination, j))
		} else {
			makePermutations(append(combination, j), length, callback)
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
		if len(*input) == 0 {
			return errors.New("Missing input")
		}
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
		if b <= 0 || registers[ins.a] < 0 {
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

func runProgram(instructions []Instruction, registers map[string]int, input *[]int) (map[string]int, error) {
	for _, ins := range instructions {
		err := runInstruction(ins, input, registers)

		if err != nil {
			return registers, err
		}
	}
	return registers, nil
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

func programPart(z, input int) int {

	x := (z % 26) + 14
	if x == input {
		x = 0
	} else {
		x = 1
	}
	y := 25*x + 1 // 26 or 1
	z = z * y

	z += (input + 12) * x
	return z
}
