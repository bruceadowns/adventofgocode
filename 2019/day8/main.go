package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type layer [25][6]uint8

func (l layer) String() string {
	var sb strings.Builder
	sb.WriteByte('\n')
	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			sb.WriteByte(l[j][i])
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (l layer) count() (int, int, int) {
	m := make(map[uint8]int)
	for i := 0; i < 25; i++ {
		for j := 0; j < 6; j++ {
			m[l[i][j]]++
		}
	}

	return m['0'], m['1'], m['2']
}

func part1(image []layer) (res int) {
	min0 := math.MaxInt32
	var min1, min2 int
	for _, v := range image {
		c0, c1, c2 := v.count()
		if c0 < min0 {
			min0 = c0
			min1 = c1
			min2 = c2
		}
	}

	return min1 * min2
}

func part2(image []layer) (res layer) {
	for i := 0; i < 25; i++ {
		for j := 0; j < 6; j++ {
			var pixel uint8
			for _, v := range image {
				pixel = v[i][j]
				if pixel != '2' {
					break
				}
			}
			res[i][j] = pixel
		}
	}

	return
}

func getin(r io.Reader) (res []layer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()

	var in []uint8
	for i := 0; i < len(line); i++ {
		in = append(in, line[i])
	}

	var curr layer
	var row, column uint8
	for _, v := range in {
		curr[row][column] = v

		row++
		if row > 24 {
			row = 0
			column++
		}

		if column > 5 {
			column = 0
			res = append(res, curr)
			curr = layer{}
		}
	}

	return
}

func main() {
	in := getin(os.Stdin)
	log.Printf("part1: %d", part1(in))
	log.Printf("part2: %s", part2(in))
}
