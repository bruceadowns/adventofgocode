package main

import (
	"bytes"
	"testing"
)

var in []int

func init() {
	stdin := `16,1,2,0,4,2,7,1,2,14`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := 37

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(in)
	expected := 168

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
