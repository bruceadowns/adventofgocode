package main

import (
	"bytes"
	"testing"
)

func Test_Part1(t *testing.T) {
	var testTable = []struct {
		dataStream string
		expected   int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, test := range testTable {
		in := In(bytes.NewBufferString(test.dataStream))
		actual := Part1(in)
		expected := test.expected

		if actual != expected {
			t.Errorf("actual: %d expected: %d", actual, expected)
		}
	}
}

func Test_Part2(t *testing.T) {
	var testTable = []struct {
		dataStream string
		expected   int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, test := range testTable {
		in := In(bytes.NewBufferString(test.dataStream))
		actual := Part2(in)
		expected := test.expected

		if actual != expected {
			t.Errorf("actual: %d expected: %d", actual, expected)
		}
	}
}
