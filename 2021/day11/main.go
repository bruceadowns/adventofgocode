package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type cell struct {
	value   int
	flashed bool
}

func (c cell) rune() rune {
	if c.value > 9 {
		return '*'
	}

	return rune(c.value + 48)
}

func (c *cell) set(i uint8) {
	c.value = int(i - 48)
}

func (c *cell) flash() bool {
	if !c.flashed && c.value > 9 {
		c.flashed = true
		return true
	}

	return false
}

func (c *cell) inc() {
	c.value++
}

func (c *cell) reset() {
	c.value = 0
	c.flashed = false
}

func newCell() *cell {
	return &cell{math.MinInt, false}
}

const SIZE = 10

type Grid [SIZE + 2][SIZE + 2]*cell

func (g Grid) String() string {
	var sb strings.Builder
	for x := 1; x <= SIZE; x++ {
		for y := 1; y <= SIZE; y++ {
			sb.WriteRune(g[y][x].rune())
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func Part1(in Grid) (res int) {
	//fmt.Println(in)

	var incAdj func(x, y int)
	incAdj = func(x, y int) {
		if in[y][x].flash() {
			in[y][x+1].inc()
			in[y-1][x+1].inc()
			in[y-1][x].inc()
			in[y-1][x-1].inc()
			in[y][x-1].inc()
			in[y+1][x-1].inc()
			in[y+1][x].inc()
			in[y+1][x+1].inc()

			incAdj(x+1, y)
			incAdj(x+1, y-1)
			incAdj(x, y-1)
			incAdj(x-1, y-1)
			incAdj(x-1, y)
			incAdj(x-1, y+1)
			incAdj(x, y+1)
			incAdj(x+1, y+1)
		}
	}

	for step := 0; step < 100; step++ {
		for x := 1; x <= SIZE; x++ {
			for y := 1; y <= SIZE; y++ {
				in[y][x].inc()
			}
		}

		for x := 1; x <= SIZE; x++ {
			for y := 1; y <= SIZE; y++ {
				incAdj(x, y)
			}
		}

		for x := 1; x <= SIZE; x++ {
			for y := 1; y <= SIZE; y++ {
				if in[y][x].flashed {
					in[y][x].reset()
					res++
				}
			}
		}
	}

	return
}

func Part2(in Grid) (res int) {
	var incAdj func(x, y int)
	incAdj = func(x, y int) {
		if in[y][x].flash() {
			in[y][x+1].inc()
			in[y-1][x+1].inc()
			in[y-1][x].inc()
			in[y-1][x-1].inc()
			in[y][x-1].inc()
			in[y+1][x-1].inc()
			in[y+1][x].inc()
			in[y+1][x+1].inc()

			incAdj(x+1, y)
			incAdj(x+1, y-1)
			incAdj(x, y-1)
			incAdj(x-1, y-1)
			incAdj(x-1, y)
			incAdj(x-1, y+1)
			incAdj(x, y+1)
			incAdj(x+1, y+1)
		}
	}

	for res = 1; ; res++ {
		for x := 1; x <= SIZE; x++ {
			for y := 1; y <= SIZE; y++ {
				in[y][x].inc()
			}
		}

		for x := 1; x <= SIZE; x++ {
			for y := 1; y <= SIZE; y++ {
				incAdj(x, y)
			}
		}

		var total int
		for x := 1; x <= SIZE; x++ {
			for y := 1; y <= SIZE; y++ {
				if in[y][x].flashed {
					in[y][x].reset()
					total++
				}
			}
		}

		if total == SIZE*SIZE {
			break
		}
	}

	return
}

func MakeGrid(in []string) (res Grid) {
	for x := 0; x < SIZE+2; x++ {
		for y := 0; y < SIZE+2; y++ {
			res[y][x] = newCell()
		}
	}

	for x := 1; x <= SIZE; x++ {
		for y := 1; y <= SIZE; y++ {
			res[y][x].set(in[x-1][y-1])
		}
	}

	return
}

func In(r io.Reader) (res []string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return
}

func main() {
	in := In(os.Stdin)
	log.Printf("part1: %d", Part1(MakeGrid(in)))
	log.Printf("part2: %d", Part2(MakeGrid(in)))
}
