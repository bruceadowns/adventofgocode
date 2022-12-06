package main

import (
	"bytes"
	"testing"
)

var lines Input

func init() {
	stdin := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
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
