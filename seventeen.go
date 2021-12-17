package main

import (
	"fmt"
)

type TargetArea struct {
	XMin int
	XMax int
	YMin int
	YMax int
}

func Seventeen() {
	// input := "target area: x=20..30, y=-10..-5"
	// arget area: x=25..67, y=-260..-200
	// targetArea := TargetArea{
	// 	20,
	// 	30,
	// 	-10,
	// 	-5,
	// }
	targetArea := TargetArea{
		25,
		67,
		-260,
		-200,
	}

	highestY := 0
	hits := 0
	for velX := 0; velX < 500; velX++ {
		for velY := -300; velY < 50000; velY++ {
			hit, localHighestY := simulate(velX, velY, targetArea)

			if hit {
				hits++
				highestY = maxInt(localHighestY, highestY)
			}
		}
	}

	// simulate(6, 0, targetArea)

	fmt.Printf("hightest %d %d\n", hits, highestY)
}

func simulate(velX int, velY int, targetArea TargetArea) (bool, int) {
	// initialVelX := velX
	// initialVelY := velY
	probeX := 0
	probeY := 0

	highestY := 0
	for probeX <= targetArea.XMax && probeY >= targetArea.YMin {
		probeX += velX
		probeY += velY
		highestY = maxInt(probeY, highestY)
		velX = decrease(velX)
		velY -= 1
		// fmt.Printf("%v %v\n", probeX, probeY)

		if targetArea.XMin <= probeX && probeX <= targetArea.XMax &&
			targetArea.YMin <= probeY && probeY <= targetArea.YMax {
			// fmt.Printf("hit %v %v\n", initialVelX, initialVelY)
			return true, highestY

		}

	}
	return false, highestY
}

func decrease(val int) int {
	if val == 0 {
		return val
	}
	if val > 0 {
		return val - 1
	}
	return val + 1
}

// 10,-1
// 10,-2
// 11,-1
// 11,-2
// 11,-3
// 11,-4
// 12,-2
// 12,-3
// 12,-4
// 13,-2
// 13,-3
// 13,-4
// 14,-2
// 14,-3
// 14,-4
// 15,-2
// 15,-3
// 15,-4
// 20,-10
// 20,-5
// 20,-6
// 20,-7
// 20,-8
// 20,-9
// 21,-10
// 21,-5
// 21,-6
// 21,-7
// 21,-8
// 21,-9
// 22,-10
// 22,-5
// 22,-6
// 22,-7
// 22,-8
// 22,-9
// 23,-10
// 23,-5
// 23,-6
// 23,-7
// 23,-8
// 23,-9
// 24,-10
// 24,-5
// 24,-6
// 24,-7
// 24,-8
// 24,-9
// 25,-10
// 25,-5
// 25,-6
// 25,-7
// 25,-8
// 25,-9
// 26,-10
// 26,-5
// 26,-6
// 26,-7
// 26,-8
// 26,-9
// 27,-10
// 27,-5
// 27,-6
// 27,-7
// 27,-8
// 27,-9
// 28,-10
// 28,-5
// 28,-6
// 28,-7
// 28,-8
// 28,-9
// 29,-10
// 29,-5
// 29,-6
// 29,-7
// 29,-8
// 29,-9
// 30,-10
// 30,-5
// 30,-6
// 30,-7
// 30,-8
// 30,-9
// 6,0
// 6,1
// 6,2
// 6,3
// 6,4
// 6,5
// 6,6
// 6,7
// 6,8
// 6,9
// 7,-1
// 7,0
// 7,1
// 7,2
// 7,3
// 7,4
// 7,5
// 7,6
// 7,7
// 7,8
// 7,9
// 8,-1
// 8,-2
// 8,0
// 8,1
// 9,-1
// 9,-2
// 9,0
