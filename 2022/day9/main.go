package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type direction int

const (
	R direction = iota
	L
	D
	U
)

type coord struct {
	x int
	y int
}

type Input []direction

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var ch rune
		var v int
		n, err := fmt.Sscanf(line, "%c %d", &ch, &v)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatal("invalid input")
		}

		var dir direction
		switch ch {
		case 'R':
			dir = R
		case 'L':
			dir = L
		case 'D':
			dir = D
		case 'U':
			dir = U
		default:
			log.Fatal("invalid input")
		}

		for i := 0; i < v; i++ {
			res = append(res, dir)
		}
	}

	return
}

func Part1(in Input) (res int) {
	var head coord
	var tail coord
	var trail = make(map[coord]struct{})
	for _, motion := range in {
		switch motion {
		case R:
			head.x++
		case L:
			head.x--
		case U:
			head.y--
		case D:
			head.y++
		default:
			log.Fatal("invalid input")
		}

		if head.x-tail.x == 2 {
			if head.y == tail.y {
				tail.x++
			} else if head.y-tail.y == 1 {
				tail.x++
				tail.y++
			} else if head.y-tail.y == -1 {
				tail.x++
				tail.y--
			} else {
				log.Fatal("invalid input")
			}
		} else if head.y-tail.y == 2 {
			if head.x == tail.x {
				tail.y++
			} else if head.x-tail.x == 1 {
				tail.x++
				tail.y++
			} else if head.x-tail.x == -1 {
				tail.x--
				tail.y++
			} else {
				log.Fatal("invalid input")
			}
		} else if head.x-tail.x == -2 {
			if head.y == tail.y {
				tail.x--
			} else if head.y-tail.y == 1 {
				tail.x--
				tail.y++
			} else if head.y-tail.y == -1 {
				tail.x--
				tail.y--
			} else {
				log.Fatal("invalid input")
			}
		} else if head.y-tail.y == -2 {
			if head.x == tail.x {
				tail.y--
			} else if head.x-tail.x == 1 {
				tail.x++
				tail.y--
			} else if head.x-tail.x == -1 {
				tail.x--
				tail.y--
			} else {
				log.Fatal("invalid input")
			}
		}

		trail[tail] = struct{}{}
	}

	return len(trail)
}

func Part2(in Input) (res int) {
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
