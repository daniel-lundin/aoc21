package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Cube struct {
	minPoint Coord3D
	xLength  int
	yLength  int
	zLength  int
	mode     string
}

func makeCube(minPoint Coord3D, xLength, yLength, zLength int, mode string) Cube {
	return Cube{minPoint, xLength, yLength, zLength, mode}
}

func TwentyTwo() {
	file, err := os.Open("./input-22.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cubes := make([]Cube, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		mode := parts[0]
		coords := strings.Split(parts[1], ",")

		parts = strings.Split(coords[0], "=")
		span := strings.Split(parts[1], "..")
		xFrom := parseInt(span[0])
		xTo := parseInt(span[1])

		parts = strings.Split(coords[1], "=")
		span = strings.Split(parts[1], "..")
		yFrom := parseInt(span[0])
		yTo := parseInt(span[1])

		parts = strings.Split(coords[2], "=")
		span = strings.Split(parts[1], "..")
		zFrom := parseInt(span[0])
		zTo := parseInt(span[1])

		cubes = append(cubes, Cube{Coord3D{xFrom, yFrom, zFrom}, xTo - xFrom + 1, yTo - yFrom + 1, zTo - zFrom + 1, mode})
	}

	processedCubes := make([]Cube, 0)
	for _, cube := range cubes {

		updatedCubes := make([]Cube, 0)
		for _, processedCube := range processedCubes {
			if !hasCubeIntersection(cube, processedCube) {

				updatedCubes = append(updatedCubes, processedCube)
				continue
			}
			updatedCubes = append(updatedCubes, cutOut(processedCube, cube)...)
		}

		processedCubes = updatedCubes

		if cube.mode == "on" {
			processedCubes = append(processedCubes, cube)
		}

		totalSize := 0
		for _, cube := range processedCubes {
			totalSize += cubeSize(cube)
		}
	}

	totalSize := 0
	for _, cube := range processedCubes {
		totalSize += cubeSize(cube)
	}

	fmt.Printf("total cubes on %d\n", totalSize)
}

func cutOut(first Cube, second Cube) []Cube {
	minX := maxInt(first.minPoint.x, second.minPoint.x)
	maxX := minInt(first.minPoint.x+first.xLength-1, second.minPoint.x+second.xLength-1)
	minY := maxInt(first.minPoint.y, second.minPoint.y)
	maxY := minInt(first.minPoint.y+first.yLength-1, second.minPoint.y+second.yLength-1)
	minZ := maxInt(first.minPoint.z, second.minPoint.z)
	maxZ := minInt(first.minPoint.z+first.zLength-1, second.minPoint.z+second.zLength-1)

	cubes := make([]Cube, 0)

	// X slices
	if first.minPoint.x < minX {
		cube := Cube{first.minPoint, minX - first.minPoint.x, first.yLength, first.zLength, "on"}
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}
	if first.minPoint.x+first.xLength-1 > maxX {
		cube := Cube{Coord3D{maxX + 1, first.minPoint.y, first.minPoint.z}, (first.minPoint.x + first.xLength) - maxX - 1, first.yLength, first.zLength, "on"}
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	// Y Slices
	if first.minPoint.y < minY {
		cube := Cube{Coord3D{minX, first.minPoint.y, first.minPoint.z}, maxX - minX + 1, minY - first.minPoint.y, first.zLength, "on"}
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	if first.minPoint.y+first.yLength-1 > maxY {
		cube := Cube{Coord3D{minX, maxY + 1, first.minPoint.z}, maxX - minX + 1, (first.minPoint.y + first.yLength - 1) - maxY, first.zLength, "on"}
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	// Z Slices
	if first.minPoint.z < minZ {
		cube := Cube{Coord3D{minX, minY, first.minPoint.z}, maxX - minX + 1, maxY - minY + 1, minZ - first.minPoint.z, "on"}
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	if first.minPoint.z+first.zLength-1 > maxZ {
		cube := Cube{Coord3D{minX, minY, maxZ + 1}, maxX - minX + 1, maxY - minY + 1, first.minPoint.z + first.zLength - maxZ - 1, "on"}
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	return cubes
}

func cubeSize(cube Cube) int {
	return (cube.xLength) * (cube.yLength) * (cube.zLength)
}

func hasCubeIntersection(first, second Cube) bool {
	if first.minPoint.x+first.xLength-1 < second.minPoint.x || first.minPoint.x > second.minPoint.x+second.xLength-1 {
		return false
	}
	if first.minPoint.y+first.yLength-1 < second.minPoint.y || first.minPoint.y > second.minPoint.y+second.yLength-1 {
		return false
	}
	if first.minPoint.z+first.zLength-1 < second.minPoint.z || first.minPoint.z > second.minPoint.z+second.zLength-1 {
		return false
	}
	return true
}
