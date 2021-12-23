package main

import (
	"fmt"
	"testing"
)

func TestHasIntersection(t *testing.T) {
	first := Cube{Coord3D{0, 0, 0}, 10, 10, 10, "on"}
	second := Cube{Coord3D{5, 5, 5}, 10, 10, 10, "on"}
	third := Cube{Coord3D{15, 15, 15}, 10, 10, 10, "on"}

	t.Run("has cube intersection", func(t *testing.T) {
		if !hasCubeIntersection(first, second) {
			t.Errorf("Should have intersection")
		}
		if hasCubeIntersection(first, third) {
			t.Errorf("Should not have intersection")
		}
	})
}

func TestSplitCube(t *testing.T) {

	// t.Run("Generated overlaps", func(t *testing.T) {
	// 	cubes := make([]Cube, 0)
	// 	for x := -2; x < 5; x++ {
	// 		for y := -2; y < 5; y++ {
	// 			for z := -2; z < 5; z++ {
	// 				for xLength := 1; xLength < 4; xLength++ {
	// 					for yLength := 1; yLength < 4; yLength++ {
	// 						for zLength := 1; zLength < 4; zLength++ {
	// 							cubes = append(cubes, Cube{Coord3D{x, y, z}, xLength, yLength, zLength, "on"})

	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 		fmt.Println("x")
	// 	}

	// 	t.Errorf("cubes %d\n", len(cubes))

	// 	cube := Cube{Coord3D{0, 0, 0}, 3, 3, 3, "on"}
	// 	for _, cube1 := range cubes {
	// 		if !hasCubeIntersection(cube, cube1) {
	// 			continue
	// 		}
	// 		splittedCubes := splitCube(cube, cube1)
	// 		area := sumCubes(splittedCubes)
	// 		cubeMap := make(map[Coord3D]bool)
	// 		manualCubeFill(cube, cubeMap)
	// 		manualCubeFill(cube1, cubeMap)
	// 		if area != cubeCountOn(cubeMap) {
	// 			t.Errorf("%v %v, expected %d, got %d\n", cube, cube1, cubeCountOn(cubeMap), area)
	// 		}
	// 	}
	// })

	t.Run("Simple left(offset X only)", func(t *testing.T) {
		first := Cube{Coord3D{0, 0, 0}, 2, 2, 2, "on"}
		second := Cube{Coord3D{1, 0, 0}, 2, 2, 2, "on"}

		splittedCubes := splitCube(first, second)

		area := sumCubes(splittedCubes)
		if area != 12 {
			t.Errorf("Expected cube sum to be 12, got %v\n", area)
		}

	})
	t.Run("Simple rright(offset X only)", func(t *testing.T) {
		first := Cube{Coord3D{1, 0, 0}, 2, 2, 2, "on"}
		second := Cube{Coord3D{0, 0, 0}, 2, 2, 2, "on"}

		splittedCubes := splitCube(first, second)

		area := sumCubes(splittedCubes)
		if area != 12 {
			t.Errorf("Expected cube sum to be 12, got %v\n", area)
		}

	})

	t.Run("Offset X/Y", func(t *testing.T) {
		first := Cube{Coord3D{0, 0, 0}, 2, 2, 1, "on"}
		second := Cube{Coord3D{1, 1, 0}, 2, 2, 1, "on"}

		splittedCubes := splitCube(first, second)
		t.Errorf("XXX splitted cubes%v\n", splittedCubes)

		area := sumCubes(splittedCubes)
		if area != 7 {
			t.Errorf("Expected cube sum to be 7, got %v\n", area)
		}

	})

	t.Run("Offset X/Y/Z", func(t *testing.T) {
		first := Cube{Coord3D{10, 10, 10}, 3, 3, 3, "on"}
		second := Cube{Coord3D{11, 11, 11}, 3, 3, 3, "on"}

		splittedCubes := splitCube(first, second)

		area := sumCubes(splittedCubes)
		if area != 46 {
			t.Errorf("Expected cube sum to be 46, got %v\n", area)
		}
	})

	t.Run("Simple off overlap", func(t *testing.T) {
		first := Cube{Coord3D{0, 0, 0}, 3, 3, 3, "on"}
		second := Cube{Coord3D{0, 0, 0}, 1, 3, 3, "off"}

		splittedCubes := splitCube(first, second)
		area := sumCubes(splittedCubes)

		if area != 18 {
			t.Errorf("Expected cube size 18, got %d\n", area)
		}
	})

	t.Run("Full overlap", func(t *testing.T) {
		first := Cube{Coord3D{0, 0, 0}, 3, 3, 3, "on"}
		second := Cube{Coord3D{0, 0, 0}, 4, 4, 4, "off"}

		splittedCubes := splitCube(first, second)
		area := sumCubes(splittedCubes)

		if area != 0 {
			t.Errorf("Expected cube size 0, got %d\n", area)
		}
	})

	t.Run("partial overlap 1", func(t *testing.T) {
		first := Cube{Coord3D{10, 10, 10}, 1, 3, 3, "on"}
		second := Cube{Coord3D{9, 9, 9}, 3, 3, 3, "off"}
		splittedCubes := splitCube(first, second)
		area := sumCubes(splittedCubes)
		cubeMap := make(map[Coord3D]bool)
		manualCubeFill(first, cubeMap)
		manualCubeFill(second, cubeMap)
		fmt.Printf("cords on%v\n", cubesOn(cubeMap))

		if area != cubeCountOn(cubeMap) {
			t.Errorf("Expected cube size 5, got %d\n", area)
		}
	})
	t.Run("partial overlap 1", func(t *testing.T) {
		first := Cube{Coord3D{11, 10, 10}, 2, 3, 3, "on"}
		second := Cube{Coord3D{9, 9, 9}, 3, 3, 3, "off"}
		splittedCubes := splitCube(first, second)
		area := sumCubes(splittedCubes)
		cubeMap := make(map[Coord3D]bool)
		manualCubeFill(first, cubeMap)
		manualCubeFill(second, cubeMap)
		fmt.Printf("cords on%v\n", cubesOn(cubeMap))

		if area != cubeCountOn(cubeMap) {
			t.Errorf("Expected cube size 5, got %d\n", area)
		}
	})
}
