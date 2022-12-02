package main

import (
	"bytes"
	"testing"
)

var lines Input

func init() {
	stdin := `A Y
B X
C Z
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(lines)
	expected := 15

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
