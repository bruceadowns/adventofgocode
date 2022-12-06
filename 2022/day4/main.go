package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type assignment struct {
	lo int
	hi int
}

func (a assignment) in(that assignment) bool {
	return a.lo >= that.lo && a.hi <= that.hi
}

func (a assignment) touch(that assignment) bool {
	return (a.lo <= that.lo && a.hi >= that.lo) ||
		(a.hi >= that.hi && a.lo <= that.hi)
}

type pair [2]assignment

type Input []pair

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var lo1, hi1, lo2, hi2 int
		n, err := fmt.Sscanf(line, "%d-%d,%d-%d", &lo1, &hi1, &lo2, &hi2)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("invalid input")
		}

		res = append(res, pair{assignment{lo1, hi1}, assignment{lo2, hi2}})
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in {
		if v[0].in(v[1]) || v[1].in(v[0]) {
			res++
		}
	}

	return
}

func Part2(in Input) (res int) {
	for _, v := range in {
		if v[0].touch(v[1]) || v[1].touch(v[0]) {
			res++
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
