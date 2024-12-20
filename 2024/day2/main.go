package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	report []int
)

type Input []report

func parse(s string) (res report) {
	for _, v := range strings.Fields(s) {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, n)
	}

	return
}

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, parse(line))
	}

	return
}

const (
	UNKNOWN = iota
	UP
	DOWN
)

func isSafe(r report) bool {
	prev := 0
	direction := UNKNOWN
	for _, v := range r {
		if prev == 0 {
			prev = v
			continue
		}
		if direction == UNKNOWN {
			if prev < v {
				direction = UP
			} else if prev > v {
				direction = DOWN
			} else {
				return false
			}
		}

		var diff int
		switch direction {
		case UP:
			if prev >= v {
				return false
			}
			diff = v - prev
		case DOWN:
			if prev <= v {
				return false
			}
			diff = prev - v
		default:
			log.Fatal("logic error")
		}

		if diff < 1 || diff > 3 {
			return false
		}

		prev = v
	}

	return true
}

func Part1(in Input) (res int) {
	for _, v := range in {
		if isSafe(v) {
			res++
		}
	}

	return
}

func Part2(in Input) (res int) {
	for _, v := range in {
		if isSafe(v) {
			res++
		} else {
			for i := 0; i < len(v); i++ {
				var vv []int
				vv = append(vv, v[:i]...)
				vv = append(vv, v[i+1:]...)
				if isSafe(vv) {
					res++
					break
				}
			}
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
