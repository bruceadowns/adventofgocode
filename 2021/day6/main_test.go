package main

import (
	"bytes"
	"testing"
)

var initialPool []*Fish

func init() {
	stdin := `3,4,3,1,2`
	initialPool = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	testTable := []struct {
		days     int
		expected int
	}{
		{18, 26},
		//{80, 5934},
		//{256, 26984457539},
	}

	for _, test := range testTable {
		actual := Part1(test.days, initialPool)
		if actual != test.expected {
			t.Errorf("actual: %d expected: %d", actual, test.expected)
		}
	}
}

func Test_Part2(t *testing.T) {
}
