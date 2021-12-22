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

const (
	on  = iota
	off = iota
)

type Cube struct {
	minPoint Coord3D
	xLength  int
	yLength  int
	zLength  int
	mode     int
}

func TwentyTwo() {
	file, err := os.Open("./input-22-example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cubes := make(map[Coord3D]bool)
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

		xOutside := (xFrom < -50 && xTo < -50) || (xFrom > 50 && xTo > 50)
		yOutside := (yFrom < -50 && yTo < -50) || (yFrom > 50 && yTo > 50)
		zOutside := (zFrom < -50 && zTo < -50) || (zFrom > 50 && zTo > 50)
		if xOutside || yOutside || zOutside {
			fmt.Printf("skipping %s\n", line)
			continue
		}
		fmt.Printf("%d %d %d %d %d %d turning %s\n", xFrom, xTo, yFrom, yTo, zFrom, zTo, mode)
		for x := xFrom; x <= xTo; x++ {
			for y := yFrom; y <= yTo; y++ {
				for z := zFrom; z <= zTo; z++ {
					cubes[Coord3D{x, y, z}] = mode == "on"
				}
			}
		}
	}
	count := 0
	for _, value := range cubes {
		if value {
			count++
		}
	}
	fmt.Printf("cubes %v\n", count)

}

func hasCubeIntersection(first Cube, second Cube) bool {
	outsideX := first.minPoint.x+first.xLength < second.minPoint.x || first.minPoint.x > second.minPoint.x+second.xLength
	outsideY := first.minPoint.y+first.yLength < second.minPoint.y || first.minPoint.y > second.minPoint.y+second.yLength
	outsideZ := first.minPoint.z+first.zLength < second.minPoint.z || first.minPoint.z > second.minPoint.z+second.zLength

	return !(outsideX && outsideY && outsideZ)

}
