package main

import (
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

	t.Run("Add cubes", func(t *testing.T) {
		first := Cube{Coord3D{10, 10, 10}, 3, 3, 3, "on"}
		second := Cube{Coord3D{11, 11, 11}, 3, 3, 3, "on"}

		splittedCubes := addCubes(first, second)
		area := sumCubes(splittedCubes)
		// t.Errorf("splitted cubes %v\n", splittedCubes)
		cubeMap := make(map[Coord3D]bool)
		manualCubeFill(first, cubeMap)
		manualCubeFill(second, cubeMap)

		if area != cubeCountOn(cubeMap) {
			t.Errorf("Expected cube sum to be %d, got %d\n", cubeCountOn(cubeMap), area)
		}
	})

	t.Run("Add cubes 2", func(t *testing.T) {
		first := Cube{Coord3D{-22, -29, -38}, 51, 53, 55, "on"}
		second := Cube{Coord3D{-20, -36, -47}, 47, 15, 55, "on"}

		splittedCubes := addCubes(first, second)
		area := sumCubes(splittedCubes)
		// t.Errorf("splitted cubes %v\n", splittedCubes)
		cubeMap := make(map[Coord3D]bool)
		manualCubeFill(first, cubeMap)
		manualCubeFill(second, cubeMap)

		if area != cubeCountOn(cubeMap) {
			t.Errorf("Expected cube sum to be %d, got %d\n", cubeCountOn(cubeMap), area)
		}
	})
	t.Run("Add cubes 3", func(t *testing.T) {
		first := Cube{Coord3D{-22, -29, -38}, 51, 53, 55, "on"}
		second := Cube{Coord3D{-20, -21, -47}, 47, 39, 21, "on"}

		splittedCubes := addCubes(first, second)
		area := sumCubes(splittedCubes)
		// t.Errorf("splitted cubes %v\n", splittedCubes)
		cubeMap := make(map[Coord3D]bool)
		manualCubeFill(first, cubeMap)
		manualCubeFill(second, cubeMap)

		if area != cubeCountOn(cubeMap) {
			t.Errorf("Expected cube sum to be %d, got %d\n", cubeCountOn(cubeMap), area)
		}
	})
	t.Run("Add cubes 4", func(t *testing.T) {
		first := Cube{Coord3D{-22, -29, -38}, 51, 53, 55, "on"}
		second := Cube{Coord3D{-20, -21, -47}, 47, 39, 21, "on"}
		// cube intersects with {{-20 -21 -26} 54 45 55 on}

		splittedCubes := addCubes(first, second)
		area := sumCubes(splittedCubes)
		// t.Errorf("splitted cubes %v\n", splittedCubes)
		cubeMap := make(map[Coord3D]bool)
		manualCubeFill(first, cubeMap)
		manualCubeFill(second, cubeMap)

		if area != cubeCountOn(cubeMap) {
			t.Errorf("Expected cube sum to be %d, got %d\n", cubeCountOn(cubeMap), area)
		}
	})
}
