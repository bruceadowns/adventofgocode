package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type (
	coord struct {
		x, y int
	}
	Input map[coord]rune
)

func In(r io.Reader) (res Input) {
	res = make(Input)
	scanner := bufio.NewScanner(r)

	// X 88
	// M 77
	// A 65
	// S 83
	y := 0
	for scanner.Scan() {
		for x, v := range scanner.Text() {
			res[coord{x, y}] = v
		}
		y++
	}

	return
}

func MAS(in Input, c coord) (res int) {
	if in[coord{c.x - 1, c.y - 1}] == 'M' && in[coord{c.x - 2, c.y - 2}] == 'A' && in[coord{c.x - 3, c.y - 3}] == 'S' {
		res++
	}
	if in[coord{c.x, c.y - 1}] == 'M' && in[coord{c.x, c.y - 2}] == 'A' && in[coord{c.x, c.y - 3}] == 'S' {
		res++
	}
	if in[coord{c.x + 1, c.y - 1}] == 'M' && in[coord{c.x + 2, c.y - 2}] == 'A' && in[coord{c.x + 3, c.y - 3}] == 'S' {
		res++
	}
	if in[coord{c.x - 1, c.y}] == 'M' && in[coord{c.x - 2, c.y}] == 'A' && in[coord{c.x - 3, c.y}] == 'S' {
		res++
	}
	if in[coord{c.x + 1, c.y}] == 'M' && in[coord{c.x + 2, c.y}] == 'A' && in[coord{c.x + 3, c.y}] == 'S' {
		res++
	}
	if in[coord{c.x - 1, c.y + 1}] == 'M' && in[coord{c.x - 2, c.y + 2}] == 'A' && in[coord{c.x - 3, c.y + 3}] == 'S' {
		res++
	}
	if in[coord{c.x, c.y + 1}] == 'M' && in[coord{c.x, c.y + 2}] == 'A' && in[coord{c.x, c.y + 3}] == 'S' {
		res++
	}
	if in[coord{c.x + 1, c.y + 1}] == 'M' && in[coord{c.x + 2, c.y + 2}] == 'A' && in[coord{c.x + 3, c.y + 3}] == 'S' {
		res++
	}

	return
}

func MS(in Input, c coord) (res int) {
	if (in[coord{c.x - 1, c.y - 1}] == 'M' && in[coord{c.x + 1, c.y + 1}] == 'S' ||
		in[coord{c.x - 1, c.y - 1}] == 'S' && in[coord{c.x + 1, c.y + 1}] == 'M') &&
		(in[coord{c.x - 1, c.y + 1}] == 'M' && in[coord{c.x + 1, c.y - 1}] == 'S' ||
			in[coord{c.x - 1, c.y + 1}] == 'S' && in[coord{c.x + 1, c.y - 1}] == 'M') {
		res++
	}

	return
}

func Part1(in Input) (res int) {
	for xCoord, X := range in {
		if X == 'X' {
			res += MAS(in, xCoord)
		}
	}

	return
}

func Part2(in Input) (res int) {
	for mCoord, A := range in {
		if A == 'A' {
			res += MS(in, mCoord)
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
