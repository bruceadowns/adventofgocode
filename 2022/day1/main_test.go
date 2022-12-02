package main

import (
	"bytes"
	"testing"
)

var in Input

func init() {
	stdin := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := 24000

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(in)
	expected := 45000

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
