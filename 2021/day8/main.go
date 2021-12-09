package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Part1(in []string) (res int) {
	for k, v := range in {
		log.Printf("%d: %s", k, v)
	}

	return
}

func Part2(in []string) (res int) {
	return
}

func In(r io.Reader) (res []string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
