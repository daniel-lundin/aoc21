package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Failed to parse %s into int\n", str)
		os.Exit(1)
	}
	return number
}

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
	file, err := os.Open("./input-22-small-example.txt")
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

	// newCubes := splitCube(cubes[0], cubes[1], "on")
	// fmt.Printf("1,2 split %v\n", newCubes)
	// for _, newCube := range newCubes {
	// 	fmt.Println(newCube)
	// }

	processedCubes := make([]Cube, 0)
	for index, cube := range cubes {

		processed := false
		updatedCubes := make([]Cube, 0)
		// fmt.Printf(">> running command with cube %v\n", cube)
		fmt.Printf("cubes count: %d\n", len(processedCubes))
		for _, processedCube := range processedCubes {
			if !hasCubeIntersection(cube, processedCube) {
				updatedCubes = append(updatedCubes, processedCube)
				continue
			}
			splitted := splitCube(processedCube, cube)
			// fmt.Printf("splitted cubes %v\n", splitted)
			updatedCubes = append(updatedCubes, splitted...)
			processed = true
		}
		// fmt.Printf("update cubes %v\n", updatedCubes)
		processedCubes = updatedCubes

		if !processed && cube.mode == "on" {
			processedCubes = append(processedCubes, cube)
		}

		// totalSize := 0
		// for _, cube := range processedCubes {
		// 	totalSize += cubeSize(cube)
		// }
		// fmt.Println("===")
		//fmt.Printf(">> total size after step %d: %d\n", index, totalSize)
		fmt.Printf(">> step %d complete \n", index)
		// fmt.Println("===")
	}

	totalSize := 0
	for _, cube := range processedCubes {
		totalSize += cubeSize(cube)
	}

	fmt.Printf("total size %d\n", totalSize)
}

func manualCubeFill(cube Cube, cubesMap map[Coord3D]bool) {
	for x := cube.minPoint.x; x < cube.minPoint.x+cube.xLength; x++ {
		for y := cube.minPoint.y; y < cube.minPoint.y+cube.yLength; y++ {
			for z := cube.minPoint.z; z < cube.minPoint.z+cube.zLength; z++ {
				cubesMap[Coord3D{x, y, z}] = cube.mode == "on"

			}
		}
	}
}

func cubeCountOn(cubesMap map[Coord3D]bool) int {
	count := 0
	for _, value := range cubesMap {
		if value {
			count++
		}
	}
	return count
}
func cubesOn(cubesMap map[Coord3D]bool) []Coord3D {
	coords := make([]Coord3D, 0)
	for coord, value := range cubesMap {
		if value {
			coords = append(coords, coord)
		}
	}
	return coords
}

func hasCubeIntersection(first, second Cube) bool {
	if first.minPoint.x+first.xLength < second.minPoint.x || first.minPoint.x > second.minPoint.x+second.xLength {
		return false
	}
	if first.minPoint.y+first.yLength < second.minPoint.y || first.minPoint.y > second.minPoint.y+second.yLength {
		return false
	}
	if first.minPoint.z+first.zLength < second.minPoint.z || first.minPoint.z > second.minPoint.z+second.zLength {
		return false
	}
	return true
}

func sumCubes(cubes []Cube) int {
	totalSize := 0
	for _, cube := range cubes {
		size := cubeSize(cube)
		// fmt.Printf("cube size %d\n", size)
		totalSize += size
	}

	return totalSize
}

func splitCube(first Cube, second Cube) []Cube {

	minX := maxInt(first.minPoint.x, second.minPoint.x)
	maxX := minInt(first.minPoint.x+first.xLength-1, second.minPoint.x+second.xLength-1)
	minY := maxInt(first.minPoint.y, second.minPoint.y)
	maxY := minInt(first.minPoint.y+first.yLength-1, second.minPoint.y+second.yLength-1)
	minZ := maxInt(first.minPoint.z, second.minPoint.z)
	maxZ := minInt(first.minPoint.z+first.zLength-1, second.minPoint.z+second.zLength-1)

	// fmt.Printf("Intersection: minX: %d maxX: %d minY: %d maxY: %d minZ: %d maxZ: %d \n", minX, maxX, minY, maxY, minZ, maxZ)
	// fmt.Printf("First %v, second %v, mode: %s\n", first, second, second.mode)

	cubes := make([]Cube, 0)

	// X slices
	if first.minPoint.x < minX {
		cube := Cube{first.minPoint, minX - first.minPoint.x, first.yLength, first.zLength, "on"}
		// fmt.Printf(" - Case 1, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}
	if first.minPoint.x+first.xLength-1 > maxX {
		cube := Cube{Coord3D{maxX + 1, first.minPoint.y, first.minPoint.z}, (first.minPoint.x + first.xLength) - maxX - 1, first.yLength, first.zLength, "on"}
		// fmt.Printf(" - Case 2, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}
	if second.mode == "on" && second.minPoint.x < minX {
		cube := Cube{second.minPoint, minX - second.minPoint.x, second.yLength, second.zLength, "on"}
		// fmt.Printf(" - Case 3, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}
	if second.mode == "on" && second.minPoint.x+second.xLength-1 > maxX {
		cube := Cube{Coord3D{maxX + 1, second.minPoint.y, second.minPoint.z}, second.minPoint.x + second.xLength - maxX - 1, second.yLength, second.zLength, "on"}
		// fmt.Printf(" - Case 4, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	// Y Slices
	if first.minPoint.y < minY {
		cube := Cube{Coord3D{minX, first.minPoint.y, first.minPoint.z}, maxX - minX + 1, minY - first.minPoint.y, first.zLength, "on"}
		// fmt.Printf(" - Case 5, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	if first.minPoint.y+first.yLength-1 > maxY {
		cube := Cube{Coord3D{minX, maxY + 1, first.minPoint.z}, maxX - minX + 1, (first.minPoint.y + first.yLength - 1) - maxY, first.zLength, "on"}
		// fmt.Printf(" - Case 6, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	if second.mode == "on" && second.minPoint.y < minY {
		cube := Cube{Coord3D{minX, second.minPoint.y, second.minPoint.z}, maxX - minX + 1, minY - second.minPoint.y, second.zLength, "on"}
		// fmt.Printf(" - Case 7, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}
	if second.mode == "on" && second.minPoint.y+second.yLength-1 > maxY {
		cube := Cube{Coord3D{minX, maxY + 1, second.minPoint.z}, maxX - minX + 1, second.minPoint.y + second.yLength - 1 - maxY, second.zLength, "on"}
		// fmt.Printf(" - Case 8, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	// Z Slices
	if first.minPoint.z < minZ {
		cube := Cube{Coord3D{minX, minY, first.minPoint.z}, maxX - minX + 1, maxY - minY + 1, minZ - first.minPoint.z, "on"}
		// fmt.Printf(" - Case 9, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	if first.minPoint.z+first.zLength-1 > maxZ {
		cube := Cube{Coord3D{minX, minY, maxZ + 1}, maxX - minX + 1, maxY - minY + 1, first.minPoint.z + first.zLength - maxZ - 1, "on"}
		// fmt.Printf(" - Case 10, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	if second.mode == "on" && second.minPoint.z < minZ {
		cube := Cube{Coord3D{minX, minY, second.minPoint.z}, maxX - minX + 1, maxY - minY + 1, minZ - second.minPoint.z, "on"}
		// fmt.Printf(" - Case 11, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}
	if second.mode == "on" && second.minPoint.z+second.zLength-1 > maxZ {
		cube := Cube{Coord3D{minX, minY, maxZ + 1}, maxX - minX + 1, maxY - minY + 1, second.minPoint.z + second.zLength - maxZ - 1, "on"}
		// fmt.Printf(" - Case 12, %v %d\n", cube, cubeSize(cube))
		if cubeSize(cube) > 0 {
			cubes = append(cubes, cube)
		}
	}

	// Add intersection
	if second.mode == "on" {
		cubes = append(cubes, Cube{Coord3D{minX, minY, minZ}, maxX - minX + 1, maxY - minY + 1, maxZ - minZ + 1, "on"})
	}

	// filteredCubes := make([]Cube, 0)
	// for _, cube := range cubes {
	// 	if cubeSize(cube) > 0 {
	// 		filteredCubes = append(filteredCubes, cube)
	// 	}
	// }

	// fmt.Printf("filtered cubes %v\n", filteredCubes)
	return cubes

	// cubeMap := make(map[Coord3D]bool)
	// manualCubeFill(first, cubeMap)
	// manualCubeFill(second, cubeMap)

	// area := sumCubes(filteredCubes)
	// if area != cubeCountOn(cubeMap) {

	// 	fmt.Printf("\nFail expected %d, got %d!\n\n", cubeCountOn(cubeMap), area)
	// 	os.Exit(1)
	// 	// fmt.Printf("Fail\n")
	// } else {
	// 	// fmt.Printf("\nCorrect!\n\n")
	// }
	// return filteredCubes
}

func cubeSize(cube Cube) int {
	return (cube.xLength) * (cube.yLength) * (cube.zLength)
}
