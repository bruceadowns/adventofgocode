package main

import (
	"bytes"
	"testing"
)

var lines Input

func init() {
	stdin := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + 
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(lines)
	expected := 4277556

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(lines)
	expected := 3263827

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
