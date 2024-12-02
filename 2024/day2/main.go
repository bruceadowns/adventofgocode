package main

import (
	"bufio"
	"io"
	"log"
	"math"
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

func Part1(in Input) (res int) {
	for _, v := range in {
		if func() bool {
			prev := 0
			direction := UNKNOWN
			for _, vv := range v {
				if prev == 0 {
					prev = vv
					continue
				}
				if direction == UNKNOWN {
					if prev < vv {
						direction = UP
					} else if prev > vv {
						direction = DOWN
					} else {
						return false
					}
				}

				switch direction {
				case UP:
					if prev > vv {
						return false
					}
				case DOWN:
					if prev < vv {
						return false
					}
				case UNKNOWN:
					log.Fatal("logic error")
				default:
					log.Fatal("logic error")
				}

				diff := math.Abs(float64(prev - vv))
				if diff < 1 || diff > 3 {
					return false
				}

				prev = vv
			}

			return true
		}() {
			res++
		}
	}

	return
}

func Part2(in Input) (res int) {
	return 4
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
