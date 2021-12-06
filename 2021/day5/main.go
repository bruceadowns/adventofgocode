package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Coord struct {
	x int
	y int
}
type Line [2]Coord
type Lines []Line

const (
	unknown = iota
	right
	left
	down
	up
	rightdown
	rightup
	leftdown
	leftup
)

func (c Coord) equal(that Coord) bool {
	return c.x == that.x && c.y == that.y
}

func (c *Coord) inc(dir int) {
	switch dir {
	case right:
		c.x++
	case left:
		c.x--
	case down:
		c.y++
	case up:
		c.y--
	case rightdown:
		c.x++
		c.y++
	case rightup:
		c.x++
		c.y--
	case leftdown:
		c.x--
		c.y++
	case leftup:
		c.x--
		c.y--
	case unknown:
		log.Fatal("invalid input")
	}
}

func straightLineCoords(start, end Coord) (res []Coord) {
	var dir = unknown
	if start.x == end.x {
		if start.y < end.y {
			dir = down
		} else {
			dir = up
		}
	} else if start.y == end.y {
		if start.x < end.x {
			dir = right
		} else {
			dir = left
		}
	} else {
		return
	}

	for curr := start; !curr.equal(end); {
		res = append(res, curr)
		curr.inc(dir)
	}
	res = append(res, end)

	return
}

func diagonalLineCoords(start, end Coord) (res []Coord) {
	var dir = unknown
	if start.x == end.x {
		if start.y < end.y {
			dir = down
		} else {
			dir = up
		}
	} else if start.y == end.y {
		if start.x < end.x {
			dir = right
		} else {
			dir = left
		}
	} else {
		if start.x < end.x {
			if start.y < end.y {
				dir = rightdown
			} else {
				dir = rightup
			}
		} else {
			if start.y < end.y {
				dir = leftdown
			} else {
				dir = leftup
			}
		}
	}

	for curr := start; !curr.equal(end); {
		res = append(res, curr)
		curr.inc(dir)
	}
	res = append(res, end)

	return
}

func Part1(lines Lines) (res int) {
	diagram := make(map[Coord]int)
	for _, v := range lines {
		for _, vv := range straightLineCoords(v[0], v[1]) {
			diagram[vv]++
		}
	}

	for _, v := range diagram {
		if v >= 2 {
			res++
		}
	}

	return res
}

func Part2(lines Lines) (res int) {
	diagram := make(map[Coord]int)
	for _, v := range lines {
		for _, vv := range diagonalLineCoords(v[0], v[1]) {
			diagram[vv]++
		}
	}

	for _, v := range diagram {
		if v >= 2 {
			res++
		}
	}

	return
}

func In(r io.Reader) (lines Lines) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		in := scanner.Text()

		var x1, y1, x2, y2 int
		n, err := fmt.Sscanf(in, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatalf("invalid input: %s", in)
		}
		if n != 4 {
			log.Fatalf("invalid input: %s", in)
		}

		lines = append(lines, Line{Coord{x1, y1}, Coord{x2, y2}})
	}

	return
}

func main() {
	lines := In(os.Stdin)
	log.Printf("part1: %d", Part1(lines))
	log.Printf("part2: %d", Part2(lines))
}
