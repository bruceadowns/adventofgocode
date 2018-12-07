package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type coord struct {
	x, y int
}

type data struct {
	id       coord
	distance int
}

var dup = coord{-1, -1}

func absDiff(l, r int) int {
	if l < r {
		return r - l
	}
	return l - r
}

func manhattanDistance(l, r coord) (res int) {
	if l.x < r.x {
		res += r.x - l.x
	} else {
		res += l.x - r.x
	}

	if l.y < r.y {
		res += r.y - l.y
	} else {
		res += l.y - r.y
	}

	return
}

func minMax(in []coord) (minX, minY, maxX, maxY int) {
	minX, minY = math.MaxInt32, math.MaxInt32
	maxX, maxY = math.MinInt32, math.MinInt32
	for _, v := range in {
		if v.x < minX {
			minX = v.x
		}
		if v.y < minY {
			minY = v.y
		}
		if v.x > maxX {
			maxX = v.x
		}
		if v.y > maxY {
			maxY = v.y
		}
	}

	return
}

func part1(in []coord) (res int) {
	minX, minY, maxX, maxY := minMax(in)

	// calculate all distances for all inputs
	grid := make(map[coord]data)
	for _, i := range in {
		for y := minY - 1; y < maxY+2; y++ {
			for x := minX - 1; x < maxX+2; x++ {
				c := coord{x, y}
				md := manhattanDistance(i, c)

				if d, ok := grid[c]; ok {
					if md < d.distance {
						grid[c] = data{i, md}
					} else if md == d.distance {
						grid[c] = data{dup, md}
					}
				} else {
					grid[c] = data{i, md}
				}
			}
		}
	}

	// count identifiers
	m := make(map[coord]int)
	for y := minY - 1; y < maxY+2; y++ {
		for x := minX - 1; x < maxX+2; x++ {
			m[grid[coord{x, y}].id]++
		}
	}

	// remove infinite
	for x := minX - 1; x < maxX+2; x++ {
		delete(m, grid[coord{x, minY - 1}].id)
		delete(m, grid[coord{x, maxY + 1}].id)
	}
	for y := minY - 1; y < maxY+2; y++ {
		delete(m, grid[coord{minX - 1, y}].id)
		delete(m, grid[coord{maxX + 1, y}].id)
	}

	// find max
	res = math.MinInt32
	for _, v := range m {
		if v > res {
			res = v
		}
	}

	return
}

func part2(in []coord) (res int) {
	minX, minY, maxX, maxY := minMax(in)

	// count abs difference for each point to each input
	m := make(map[coord]struct{})
	for y := minY - 1; y < maxY+2; y++ {
		for x := minX - 1; x < maxX+2; x++ {
			var total int
			for _, i := range in {
				total += absDiff(x, i.x) + absDiff(y, i.y)
			}
			if total < 10000 {
				m[coord{x, y}] = struct{}{}
			}
		}
	}

	return len(m)
}

func in(r io.Reader) (res []coord) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var c coord
		num, err := fmt.Sscanf(scanner.Text(), "%d, %d", &c.x, &c.y)
		if err != nil {
			log.Fatal(err)
		}
		if num != 2 {
			log.Fatal("invalid input")
		}

		res = append(res, c)
	}

	return
}

func main() {
	cs := in(os.Stdin)
	log.Printf("size of largest non-infinite area: %d", part1(cs))
	log.Printf("part2: %d", part2(cs))
}
