package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type coord struct {
	x, y int
}
type Input map[coord]struct{}

func In(r io.Reader) (res Input) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var y int
	res = make(map[coord]struct{})
	for i := len(lines) - 1; i >= 0; i-- {
		for k, v := range lines[i] {
			switch v {
			case '@':
				res[coord{k, y}] = struct{}{}
			}
		}

		y++
	}

	return
}

func adjacent(in Input, c coord) (res int) {
	if _, ok := in[coord{c.x - 1, c.y + 1}]; ok {
		res++
	}
	if _, ok := in[coord{c.x, c.y + 1}]; ok {
		res++
	}
	if _, ok := in[coord{c.x + 1, c.y + 1}]; ok {
		res++
	}
	if _, ok := in[coord{c.x - 1, c.y}]; ok {
		res++
	}
	if _, ok := in[coord{c.x + 1, c.y}]; ok {
		res++
	}
	if _, ok := in[coord{c.x - 1, c.y - 1}]; ok {
		res++
	}
	if _, ok := in[coord{c.x, c.y - 1}]; ok {
		res++
	}
	if _, ok := in[coord{c.x + 1, c.y - 1}]; ok {
		res++
	}

	return
}

func Part1(in Input) (res int) {
	for k := range in {
		if adjacent(in, k) < 4 {
			res++
		}
	}

	return
}

func inCopy(in Input) (res Input) {
	res = make(Input)
	for k, v := range in {
		res[k] = v
	}

	return
}

func Part2(in Input) (res int) {
	nin := inCopy(in)

	for {
		tin := inCopy(nin)
		for k := range nin {
			if adjacent(nin, k) < 4 {
				delete(tin, k)
			}
		}

		if len(tin) == len(nin) {
			break
		}

		nin = tin
	}

	return len(in) - len(nin)
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
