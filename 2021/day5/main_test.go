package main

import (
	"bytes"
	"testing"
)

var lines Lines

func init() {
	stdin := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(lines)
	expected := 5

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(lines)
	expected := 12

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
