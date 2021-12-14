package main

import (
	"bytes"
	"fmt"
	"testing"
)

var tp TP

func init() {
	stdin := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
`
	tp = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(tp)
	expected := 17

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(tp)
	fmt.Println(actual)
	// O
	// ABKJFBGC
}
