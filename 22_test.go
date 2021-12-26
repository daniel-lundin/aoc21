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
