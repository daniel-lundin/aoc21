package main

import (
	"fmt"
	"testing"
)

func TestCubeIntersection(t *testing.T) {
	first := Cube{0, 10, 10, 10}
	second := Cube{5, 10, 10, 10}

	t.Run("Cube intersection", func(t *testing.T) {
		cubeIntersect(first Cube, second Cube)
	})
}
