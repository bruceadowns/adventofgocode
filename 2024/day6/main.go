package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

type Input struct {
	maxX         int
	maxY         int
	obstructions map[coord]struct{}
	guard        struct {
		position  coord
		direction int
	}
	visited map[coord]int
}

func In(r io.Reader) (res Input) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	res.maxX = len(lines[0])
	res.maxY = len(lines)
	res.obstructions = make(map[coord]struct{})
	res.guard.direction = UP
	res.visited = make(map[coord]int)

	var y int
	for i := len(lines) - 1; i >= 0; i-- {
		for x, r := range lines[i] {
			switch r {
			case '.':
				break
			case '#':
				res.obstructions[coord{x, y}] = struct{}{}
			case '^':
				res.guard.position = coord{x, y}
				res.visited[res.guard.position]++
			default:
				log.Fatal("invalid input")
			}
		}

		y++
	}

	return
}

func Part1(in Input) (res int) {
	for {
		switch in.guard.direction {
		case UP:
			cand := in.guard.position
			cand.y++
			if _, ok := in.obstructions[cand]; ok {
				in.guard.position.x++
				in.guard.direction = RIGHT
			} else {
				in.guard.position.y++
			}
		case DOWN:
			cand := in.guard.position
			cand.y--
			if _, ok := in.obstructions[cand]; ok {
				in.guard.position.x--
				in.guard.direction = LEFT
			} else {
				in.guard.position.y--
			}
		case LEFT:
			cand := in.guard.position
			cand.x--
			if _, ok := in.obstructions[cand]; ok {
				in.guard.position.y++
				in.guard.direction = UP
			} else {
				in.guard.position.x--
			}
		case RIGHT:
			cand := in.guard.position
			cand.x++
			if _, ok := in.obstructions[cand]; ok {
				in.guard.position.y--
				in.guard.direction = DOWN
			} else {
				in.guard.position.x++
			}
		default:
			log.Fatal("invalid input")
		}

		if in.guard.position.x < 0 || in.guard.position.x > in.maxX ||
			in.guard.position.y < 0 || in.guard.position.y > in.maxY {
			break
		}

		in.visited[in.guard.position]++
	}

	return len(in.visited)
}

func Part2(in Input) (res int) {
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
