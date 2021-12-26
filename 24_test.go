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
	t.Run("has negate program", func(t *testing.T) {
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
