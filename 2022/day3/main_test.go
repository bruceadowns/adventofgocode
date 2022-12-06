package main

import (
	"bytes"
	"testing"
)

var stdin = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func Test_Part1(t *testing.T) {
	lines := In(bytes.NewBufferString(stdin))
	actual := Part1(lines)
	expected := 157

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	lines := In2(bytes.NewBufferString(stdin))
	actual := Part2(lines)
	expected := 70

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
