package main

import (
	"bytes"
	"testing"
)

var in Input

func init() {
	stdin := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := 41

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(in)
	expected := 0

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
