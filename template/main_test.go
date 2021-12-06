package main

import (
	"bytes"
	"testing"
)

var lines []string

func init() {
	stdin := `test
one
two
three
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(lines)
	expected := 0

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(lines)
	expected := 0

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
