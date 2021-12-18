package main

import (
	"fmt"
	"testing"
)

func Explore(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"
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
