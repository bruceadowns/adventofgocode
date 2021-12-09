package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Part1(in [][]int) (res int) {
	rowLen := len(in)
	columnLen := len(in[0])
	var lowPoints []int
	for x := 1; x < columnLen-1; x++ {
		for y := 1; y < rowLen-1; y++ {
			target := in[y][x]
			up := in[y-1][x]
			down := in[y+1][x]
			left := in[y][x-1]
			right := in[y][x+1]

			if target < left && target < right && target < up && target < down {
				lowPoints = append(lowPoints, target)
			}
		}
	}

	for _, v := range lowPoints {
		res += v + 1
	}

	return
}

func Part2(in [][]int) (res int) {
	return
}

func In(r io.Reader) (res [][]int) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//rowLen := len(lines)
	columnLen := len(lines[0])

	const max int = 9

	var blankRow []int
	for i := 0; i < columnLen+2; i++ {
		blankRow = append(blankRow, max)
	}

	res = append(res, blankRow)
	for _, v := range lines {
		var row []int

		row = append(row, max)
		for _, vv := range v {
			row = append(row, int(vv-48))
		}
		row = append(row, max)

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
