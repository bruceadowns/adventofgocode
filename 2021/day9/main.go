package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
)

type Coord struct {
	x int
	y int
}
type Cell struct {
	value   int
	visited bool
}
type HeightMap [][]*Cell

const max int = 9

func Part1(in HeightMap) (res int) {
	rowLen := len(in)
	columnLen := len(in[0])
	var lowPoints []Coord
	for x := 1; x < columnLen-1; x++ {
		for y := 1; y < rowLen-1; y++ {
			target := in[y][x].value
			up := in[y-1][x].value
			down := in[y+1][x].value
			left := in[y][x-1].value
			right := in[y][x+1].value

			if target < left && target < right && target < up && target < down {
				lowPoints = append(lowPoints, Coord{x, y})
			}
		}
	}

	for _, v := range lowPoints {
		res += in[v.y][v.x].value + 1
	}

	return
}

func Part2(in HeightMap) (res int) {
	rowLen := len(in)
	columnLen := len(in[0])
	var lowPoints []Coord
	for x := 1; x < columnLen-1; x++ {
		for y := 1; y < rowLen-1; y++ {
			target := in[y][x].value
			up := in[y-1][x].value
			down := in[y+1][x].value
			left := in[y][x-1].value
			right := in[y][x+1].value

			if target < left && target < right && target < up && target < down {
				lowPoints = append(lowPoints, Coord{x, y})
			}
		}
	}

	var basin func(c Coord) (res []Coord)
	basin = func(c Coord) (res []Coord) {
		if in[c.y][c.x].visited {
			return
		}
		in[c.y][c.x].visited = true

		target := in[c.y][c.x].value
		if target == max {
			return
		}

		up := in[c.y-1][c.x].value
		down := in[c.y+1][c.x].value
		left := in[c.y][c.x-1].value
		right := in[c.y][c.x+1].value

		res = append(res, c)
		if target < up {
			res = append(res, basin(Coord{c.x, c.y - 1})...)
		}
		if target < down {
			res = append(res, basin(Coord{c.x, c.y + 1})...)
		}
		if target < left {
			res = append(res, basin(Coord{c.x - 1, c.y})...)
		}
		if target < right {
			res = append(res, basin(Coord{c.x + 1, c.y})...)
		}

		return
	}

	var basins [][]Coord
	for _, v := range lowPoints {
		basins = append(basins, basin(v))
	}

	var basinSizes []int
	for _, v := range basins {
		basinSizes = append(basinSizes, len(v))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	res = 1
	for i := 0; i < 3; i++ {
		res *= basinSizes[i]
	}

	return
}

func In(r io.Reader) (res HeightMap) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//rowLen := len(lines)
	columnLen := len(lines[0])

	var blankRow []*Cell
	for i := 0; i < columnLen+2; i++ {
		blankRow = append(blankRow, &Cell{max, true})
	}

	res = append(res, blankRow)
	for _, v := range lines {
		var row []*Cell

		row = append(row, &Cell{max, true})
		for _, vv := range v {
			row = append(row, &Cell{int(vv - 48), false})
		}
		row = append(row, &Cell{max, true})

		res = append(res, row)
	}
	res = append(res, blankRow)

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
