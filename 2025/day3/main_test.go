package main

import (
	"bytes"
	"testing"
)

var lines Input

func init() {
	stdin := `987654321111111
811111111111119
234234234234278
818181911112111
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(lines)
	expected := 357

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
