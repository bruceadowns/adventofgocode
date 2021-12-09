package main

import (
	"bytes"
	"testing"
)

var in [][]int

func init() {
	stdin := `2199943210
3987894921
9856789892
8767896789
9899965678
`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := 15

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
}
