package main

import (
	"bytes"
	"testing"
)

var lines Input

func init() {
	stdin := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(lines)
	expected := 2

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(lines)
	expected := 4

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
