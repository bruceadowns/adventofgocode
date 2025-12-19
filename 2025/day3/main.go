package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Input [][]int

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, v := range line {
			row = append(row, int(v-'0'))
		}

		res = append(res, row)
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in {
		var m int
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				s := v[i]*10 + v[j]
				if s > m {
					m = s
				}
			}
		}

		res += m
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
