package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func part1(in []string) (res int) {
	for _, v := range in {
		if i, err := strconv.Atoi(v); err == nil {
			res += i/3 - 2
		} else {
			log.Fatal("invalid input")
		}
	}

	return
}

func part2(in []string) (res int) {
	for _, v := range in {
		if i, err := strconv.Atoi(v); err == nil {
			mass := i/3 - 2
			for {
				res += mass

				mass = mass/3 - 2
				if mass < 0 {
					break
				}
			}
		} else {
			log.Fatal("invalid input")
		}
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
