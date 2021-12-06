package main

import (
	"bytes"
	"testing"
)

var in Report

func init() {
	// test input len = 5
	// puzzle input len = 12

	stdin := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := 198

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(in)
	expected := 230

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
