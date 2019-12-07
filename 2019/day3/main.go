package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	right = iota
	up
	left
	down
)

type move struct {
	direction int
	scalar    int
}

func (m move) String() string {
	var dir string
	switch m.direction {
	case right:
		dir = "right"
	case up:
		dir = "up"
	case left:
		dir = "left"
	case down:
		dir = "down"
	}

	return fmt.Sprintf("move %s %d", dir, m.scalar)
}

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c coord) dist() (res int) {
	x := c.x
	if x > 1 {
		res += x - 1
	} else if x < 1 {
		res -= x - 1
	}

	y := c.y
	if y > 1 {
		res += y - 1
	} else if y < 1 {
		res -= y - 1
	}

	return
}

func part1(in [2][]move) (res int) {
	m := make(map[coord]map[int]struct{})
	for idx := 0; idx < 2; idx++ {
		pos := coord{1, 1}

		record := func() {
			if _, ok := m[pos]; !ok {
				m[pos] = make(map[int]struct{})
			}
			m[pos][idx] = struct{}{}
		}

		for _, v := range in[idx] {
			switch v.direction {
			case right:
				for i := 0; i < v.scalar; i++ {
					pos.x++
					record()
				}
			case up:
				for i := 0; i < v.scalar; i++ {
					pos.y++
					record()
				}
			case left:
				for i := 0; i < v.scalar; i++ {
					pos.x--
					record()
				}
			case down:
				for i := 0; i < v.scalar; i++ {
					pos.y--
					record()
				}
			}
		}
	}

	res = math.MaxInt32
	for k, v := range m {
		if len(v) > 1 {
			d := k.dist()
			if res > d {
				res = d
			}
		}
	}

	return
}

func part2(in [2][]move) (res int) {
	m := make(map[coord]map[int]int)
	for idx := 0; idx < 2; idx++ {
		pos := coord{1, 1}
		step := 0

		record := func() {
			step++

			if _, ok := m[pos]; !ok {
				m[pos] = make(map[int]int)
			}
			m[pos][idx] = step
		}

		for _, v := range in[idx] {
			switch v.direction {
			case right:
				for i := 0; i < v.scalar; i++ {
					pos.x++
					record()
				}
			case up:
				for i := 0; i < v.scalar; i++ {
					pos.y++
					record()
				}
			case left:
				for i := 0; i < v.scalar; i++ {
					pos.x--
					record()
				}
			case down:
				for i := 0; i < v.scalar; i++ {
					pos.y--
					record()
				}
			}
		}
	}

	res = math.MaxInt32
	for _, v := range m {
		if len(v) > 1 {
			var d int
			for _, vv := range v {
				d += vv
			}
			if res > d {
				res = d
			}
		}
	}

	return
}

func getin(r io.Reader) (res [2][]move) {
	var idx int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var moves []move
		for _, v := range strings.Split(line, ",") {
			var m move
			switch v[0] {
			case 'R':
				m.direction = right
			case 'U':
				m.direction = up
			case 'L':
				m.direction = left
			case 'D':
				m.direction = down
			default:
				log.Fatal("invalid input")
			}

			if n, err := strconv.Atoi(v[1:]); err == nil {
				m.scalar = n
			} else {
				log.Fatal("invalid input")
			}

			moves = append(moves, m)
		}

		res[idx] = moves
		idx++
	}

	return
}

func main() {
	in := getin(os.Stdin)
	log.Printf("part1: %d", part1(in))
	log.Printf("part2: %d", part2(in))
}
