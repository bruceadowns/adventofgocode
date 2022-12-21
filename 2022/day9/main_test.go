package main

import (
	"bytes"
	"testing"
)

func Test_Part1(t *testing.T) {
	stdin := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`
	in := In(bytes.NewBufferString(stdin))
	actual := Part1(in)
	expected := 13

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	stdin := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
	in := In(bytes.NewBufferString(stdin))
	actual := Part2(in)
	expected := 36

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
