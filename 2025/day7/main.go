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
type Input []string

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return
}

func Part1(in Input) (res int) {
	beams := make(map[coord]struct{})
	for x := 0; x < len(in[0]); x++ {
		switch in[0][x] {
		case '.':
		case 'S':
			beams[coord{x, 0}] = struct{}{}
		default:
			log.Fatal()
		}
	}

	for y := 1; y < len(in); y++ {
		for x := 0; x < len(in[y]); x++ {
			switch in[y][x] {
			case '.':
				if _, ok := beams[coord{x, y - 1}]; ok {
					beams[coord{x, y}] = struct{}{}
				}
			case '^':
				if _, ok := beams[coord{x, y - 1}]; ok {
					beams[coord{x - 1, y}] = struct{}{}
					beams[coord{x + 1, y}] = struct{}{}
					res++
				}
			default:
				log.Fatal()
			}
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
