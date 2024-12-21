package main

import (
	"bytes"
	"testing"
)

var input Input

func init() {
	stdin := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	input = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(input)
	expected := 18

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(input)
	expected := 9

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
