package main

import (
	"bytes"
	"testing"
)

func Test_Part1(t *testing.T) {
	stdin := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	`
	lines := In(bytes.NewBufferString(stdin))
	actual := Part1(lines)
	expected := 142

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	stdin := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`
	lines := In(bytes.NewBufferString(stdin))
	actual := Part2(lines)
	expected := 281

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
