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

func Part1(in []int) (res int) {
	res = math.MaxInt64
	for _, v := range in {
		var curr int
		for _, vv := range in {
			curr += int(math.Abs(float64(v - vv)))
		}
		if curr < res {
			res = curr
		}
	}

	return
}

func Part2(in []int) (res int) {
	start := math.MaxInt64
	end := math.MinInt64
	for _, v := range in {
		if v < start {
			start = v
		}
		if v > end {
			end = v
		}
	}

	cost := make(map[int]int)
	for i := start - 1; i <= end; i++ {
		cost[i+1] = cost[i] + (i + 1)
	}

	res = math.MaxInt64
	for i := start; i <= end; i++ {
		var curr int
		for _, v := range in {
			curr += cost[int(math.Abs(float64(i-v)))]
		}
		if curr < res {
			res = curr
		}
	}

	return
}

func In(r io.Reader) (res []int) {
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
		res = append(res, n)
	}

	return
}

func main() {
	in := In(os.Stdin)
	log.Printf("part1: %d", Part1(in))
	log.Printf("part2: %d", Part2(in))
}
