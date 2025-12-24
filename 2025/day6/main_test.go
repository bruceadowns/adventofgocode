package main

import (
	"bytes"
	"testing"
)

var (
	lines []string
	in1   Input1
	in2   Input2
)

func init() {
	stdin := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`
	lines = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	in1 = In1(lines)
	actual := Part1(in1)
	expected := 4277556

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	in2 = In2(lines)
	actual := Part2(in2)
	expected := 3263827

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part21(t *testing.T) {
	in2 = In2(lines)
	actual := Part21(in2)
	expected := 4277556

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
