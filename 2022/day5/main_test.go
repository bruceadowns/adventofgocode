package main

import (
	"bytes"
	"testing"
)

var in Input

func init() {
	stdin := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := "CMZ"

	if actual != expected {
		t.Errorf("actual: %s expected: %s", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(in)
	expected := "MCD"

	if actual != expected {
		t.Errorf("actual: %s expected: %s", actual, expected)
	}
}
