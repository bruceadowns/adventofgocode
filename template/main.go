package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func part1(in []string) (res int) {
	for k, v := range in {
		log.Printf("%d: %s", k, v)
	}

	return
}

func part2(in []string) (res int) {
	for k, v := range in {
		log.Printf("%d: %s", k, v)
	}

	return
}

func in(r io.Reader) (res []string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return
}

func main() {
	i := in(os.Stdin)
	log.Printf("part1: %d", part1(i))
	log.Printf("part2: %d", part2(i))
}
