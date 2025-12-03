package main

import (
	"bytes"
	"testing"
)

var in Input

func init() {
	stdin := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := 3

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(in)
	expected := 6

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
