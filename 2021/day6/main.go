package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Fish struct {
	cycle int
}

func (f *Fish) dec() (res bool) {
	f.cycle--
	if f.cycle < 0 {
		f.cycle = 6
		res = true
	}

	return
}

func Part1(days int, pool []*Fish) (res int) {
	poolRes := pool
	for i := 0; i < days; i++ {
		poolNew := poolRes
		for _, v := range poolRes {
			if v.dec() {
				poolNew = append(poolNew, &Fish{8})
			}
		}
		poolRes = poolNew
	}

	return len(poolRes)
}

func Part2() (res int) {
	return
}

func In(r io.Reader) (res []*Fish) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}

	line := scanner.Text()
	for _, v := range strings.Split(line, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, &Fish{n})
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(80, i))
	log.Printf("part2: %d", Part2())
}
