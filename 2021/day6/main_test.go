package main

import (
	"bytes"
	"testing"
)

var in map[int]int

func init() {
	stdin := `3,4,3,1,2`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Solution(t *testing.T) {
	testTable := []struct {
		days     int
		expected int
	}{
		{18, 26},
		{80, 5_934},
		{256, 26_984_457_539},
	}

	for _, test := range testTable {
		actual := Solution(test.days, in)
		if actual != test.expected {
			t.Errorf("actual: %d expected: %d", actual, test.expected)
		}
	}
}
