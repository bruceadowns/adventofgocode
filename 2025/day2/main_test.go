package main

import (
	"bytes"
	"testing"
)

var in Input

func init() {
	stdin := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := 1227775554

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(in)
	expected := 4174379265

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
