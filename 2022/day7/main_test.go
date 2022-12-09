package main

import (
	"bytes"
	"testing"
)

var in *Dir

func init() {
	stdin := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
	in = In(bytes.NewBufferString(stdin))
}

func Test_Part1(t *testing.T) {
	actual := Part1(in)
	expected := 95437

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2(in)
	expected := 24933642

	if actual != expected {
		t.Errorf("actual: %d expected: %d", actual, expected)
	}
}
