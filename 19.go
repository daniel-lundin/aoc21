package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Scanner struct {
	position      []int
	beacons       [][]int
	positionKnown bool
	id            int
}

type Coord3D struct {
	x int
	y int
	z int
}

func Nineteen() {
	file, err := os.Open("./input-19.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanners := make([]Scanner, 0)
	for scanner.Scan() {
		line := scanner.Text()
		scanners = append(scanners, Scanner{})

		beacons := make([][]int, 0)
		for scanner.Scan() {
			line = scanner.Text()
			if line == "" {
				break
			}

			parts := strings.Split(line, ",")
			beacons = append(beacons, StringListToNumbers(parts))
		}
		scanners[len(scanners)-1].beacons = beacons
	}
	scanners[0].position = []int{0, 0, 0}
	scanners[0].positionKnown = true
	for index, _ := range scanners {
		scanners[index].id = index
	}

	unknown := getUnknownScanners(&scanners)
	known := getKnownScanners(&scanners)
	for len(unknown) > 0 {
		found := false
		for _, unknownScanner := range unknown {
			for _, knownScanner := range known {
				if overlaps(knownScanner, unknownScanner) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		unknown = getUnknownScanners(&scanners)
		known = getKnownScanners(&scanners)
	}

	uniqueBeacons := make(map[Coord3D]bool)
	for _, scanner := range scanners {
		for _, beacon := range scanner.beacons {
			uniqueBeacons[Coord3D{beacon[0], beacon[1], beacon[2]}] = true
		}
	}

	fmt.Printf("%d unique beacons. \n", len(uniqueBeacons))

	largestManhatten := 0
	for _, scanner := range scanners {
		for _, otherScanner := range scanners {
			if scanner.id != otherScanner.id {
				manhattan := math.Abs(float64(scanner.position[0])-float64(otherScanner.position[0])) + math.Abs(float64(scanner.position[1])-float64(otherScanner.position[1])) + math.Abs(float64(scanner.position[2])-float64(otherScanner.position[2]))
				largestManhatten = maxInt(int(manhattan), largestManhatten)
			}
		}
	}

	fmt.Printf("largest manhattan %d\n", largestManhatten)

}

func overlaps(known *Scanner, unknown *Scanner) bool {
	for _, orientation := range allOrientations(unknown.beacons) {
		for _, unknownBeacon := range orientation {
			for _, knownBeacon := range known.beacons {
				xDiff := knownBeacon[0] - unknownBeacon[0]
				yDiff := knownBeacon[1] - unknownBeacon[1]
				zDiff := knownBeacon[2] - unknownBeacon[2]

				overlaps := overlappingBeacons(known.beacons, orientation, xDiff, yDiff, zDiff)
				if overlaps > 11 {
					fmt.Printf("Overlap, updating  scanner %d!\n", unknown.id)
					unknown.position = []int{
						xDiff,
						yDiff,
						zDiff,
					}
					unknown.beacons = offset(orientation, xDiff, yDiff, zDiff)
					unknown.positionKnown = true
					return true
				}
			}
		}
	}
	return false
}

func offset(beacons [][]int, x, y, z int) [][]int {
	offsetted := make([][]int, 0)
	for _, beacon := range beacons {
		offset := []int{beacon[0] + x, beacon[1] + y, beacon[2] + z}
		offsetted = append(offsetted, offset)
	}
	return offsetted
}
func getKnownScanners(scanners *[]Scanner) []*Scanner {
	known := make([]*Scanner, 0)
	for index, scanner := range *scanners {
		if scanner.positionKnown {
			known = append(known, &(*scanners)[index])
		}
	}
	return known
}

func getUnknownScanners(scanners *[]Scanner) []*Scanner {
	unknown := make([]*Scanner, 0)
	for index, scanner := range *scanners {
		if !scanner.positionKnown {
			unknown = append(unknown, &(*scanners)[index])
		}
	}
	return unknown
}

func overlappingBeacons(beacons1 [][]int, beacons2 [][]int, xOffset, yOffset, zOffset int) int {
	overlaps := 0
	for _, beacon1 := range beacons1 {
		for _, beacon2 := range beacons2 {
			xMatch := beacon1[0] == beacon2[0]+xOffset
			yMatch := beacon1[1] == beacon2[1]+yOffset
			zMatch := beacon1[2] == beacon2[2]+zOffset
			if xMatch && yMatch && zMatch {

				overlaps++
			}
		}
	}
	return overlaps
}

func allOrientations(coords [][]int) [][][]int {
	tilts := []func([]int) []int{
		func(coords []int) []int {
			return coords
		},
		func(coords []int) []int {
			return []int{
				-coords[1],
				coords[0],
				coords[2],
			}
		},
		func(coords []int) []int {
			return []int{
				-coords[0],
				-coords[1],
				coords[2],
			}
		},
		func(coords []int) []int {
			return []int{
				coords[1],
				-coords[0],
				coords[2],
			}
		},
	}
	rotations := []func([]int) []int{
		func(coords []int) []int {
			return coords
		},
		func(coords []int) []int {
			return []int{
				coords[0],
				coords[2],
				-coords[1],
			}
		},
		func(coords []int) []int {
			return []int{
				coords[0],
				-coords[1],
				-coords[2],
			}
		},
		func(coords []int) []int {
			return []int{
				coords[0],
				-coords[2],
				coords[1],
			}
		},
		func(coords []int) []int {
			return []int{
				-coords[2],
				coords[1],
				coords[0],
			}
		},
		func(coords []int) []int {
			return []int{
				coords[2],
				coords[1],
				-coords[0],
			}
		},
	}
	orientations := make([][][]int, 24)
	for _, coord := range coords {
		for tiltIndex, tilt := range tilts {
			for rotationIndex, rotation := range rotations {
				updatedCoord := tilt(rotation(coord))
				orientations[tiltIndex*6+rotationIndex] = append(orientations[tiltIndex*6+rotationIndex], updatedCoord)
			}
		}
	}
	return orientations
}
