package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = iota
	MULTIPLY
)

type Equation struct {
	args  []int
	op    int
	total int
}
type Input []Equation

func In(r io.Reader) (res Input) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for range strings.Fields(lines[1]) {
		res = append(res, Equation{})
	}

	i := 0
	for ; i < len(lines)-1; i++ {
		for k, v := range strings.Fields(lines[i]) {
			j, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("invalid")
			}

			res[k].args = append(res[k].args, j)
		}
	}

	for k, v := range strings.Fields(lines[i]) {
		switch v {
		case "+":
			res[k].op = ADD

			for _, vv := range res[k].args {
				res[k].total += vv
			}
		case "*":
			res[k].op = MULTIPLY

			res[k].total = 1
			for _, vv := range res[k].args {
				res[k].total *= vv
			}
		default:
			log.Fatal("invalid")
		}
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in {
		res += v.total
	}

	return
}

func Part2(in Input) (res int) {
	for _, v := range in {
		res += v.total
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
