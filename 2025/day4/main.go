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

func Part1(in Input) (res int) {
	for k := range in {
		var count int
		if _, ok := in[coord{k.x - 1, k.y + 1}]; ok {
			count++
		}
		if _, ok := in[coord{k.x, k.y + 1}]; ok {
			count++
		}
		if _, ok := in[coord{k.x + 1, k.y + 1}]; ok {
			count++
		}
		if _, ok := in[coord{k.x - 1, k.y}]; ok {
			count++
		}
		if _, ok := in[coord{k.x + 1, k.y}]; ok {
			count++
		}
		if _, ok := in[coord{k.x - 1, k.y - 1}]; ok {
			count++
		}
		if _, ok := in[coord{k.x, k.y - 1}]; ok {
			count++
		}
		if _, ok := in[coord{k.x + 1, k.y - 1}]; ok {
			count++
		}

		if count < 4 {
			res++
		}
	}

	return
}

func Part2(in Input) (res int) {
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
