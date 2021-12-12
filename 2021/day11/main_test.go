package main

import (
	"bytes"
	"testing"
)

var in []string

func init() {
	stdin := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(MakeGrid(in))
	expected := 1656

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(MakeGrid(in))
	expected := 195

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
