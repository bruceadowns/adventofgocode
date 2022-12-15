package main

import (
	"bytes"
	"testing"
)

var lines Input

func init() {
	stdin := `30373
25512
65332
33549
35390
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(lines)
	expected := 21

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(lines)
	expected := 8

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
