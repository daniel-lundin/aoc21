package main

import (
	"fmt"
	"testing"
)

func TestPrevIndex(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{3, 2},
		{2, 1},
		{1, 0},
		{0, 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := prevIndex(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
