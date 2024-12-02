package main

import (
	"bytes"
	"testing"
)

var lines Input

func init() {
	stdin := `3   4
4   3
2   5
1   3
3   9
3   3
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(lines)
	expected := 11

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(lines)
	expected := 31

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
