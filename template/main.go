package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Input []string

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return
}

func Part1(in Input) (res int) {
	for k, v := range in {
		log.Printf("%d: %s", k, v)
	}

	return
}

func Part2(in Input) (res int) {
	for k, v := range in {
		log.Printf("%d: %s", k, v)
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
