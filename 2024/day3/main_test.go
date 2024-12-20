package main

import (
	"bytes"
	"testing"
)

func Test_Part1(t *testing.T) {
	stdin := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	actual := Part1(In(bytes.NewBufferString(stdin)))
	expected := 161

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	stdin := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	actual := Part2(In(bytes.NewBufferString(stdin)))
	expected := 48

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
