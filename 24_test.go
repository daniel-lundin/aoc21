package main

import (
	"fmt"
	"testing"
)

func TestNegateProgram(t *testing.T) {
	instructions := []Instruction{
		parseInstruction("inp x"),
		parseInstruction("mul x -1"),
	}
	t.Run("negate program", func(t *testing.T) {
		input := []int{2}
		registers, err := runProgram(instructions, &input)
		if err != nil {
			t.Errorf("Should not throw error")
			return
		}

		if registers["x"] != 2 {
			fmt.Printf("should be 2 got %d\n", registers["x"])
		}
	})
}
func TestBitProgram(t *testing.T) {
	instructions := []Instruction{
		parseInstruction("inp w"),
		parseInstruction("add z w"),
		parseInstruction("mod z 2"),
		parseInstruction("div w 2"),
		parseInstruction("add y w"),
		parseInstruction("mod y 2"),
		parseInstruction("div w 2"),
		parseInstruction("add x w"),
		parseInstruction("mod x 2"),
		parseInstruction("div w 2"),
		parseInstruction("mod w 2"),
	}

	t.Run("bit program program", func(t *testing.T) {
		input := []int{10}
		registers, err := runProgram(instructions, &input)
		if err != nil {
			t.Errorf("Should not throw error")
			return
		}

		if registers["w"] != 1 {
			fmt.Printf("w be 1 got %d\n", registers["w"])
		}
		if registers["x"] != 0 {
			fmt.Printf("x be 0 got %d\n", registers["x"])
		}
		if registers["y"] != 1 {
			fmt.Printf("y be 1 got %d\n", registers["y"])
		}
		if registers["z"] != 1 {
			fmt.Printf("z be 0 got %d\n", registers["z"])
		}
	})
}
