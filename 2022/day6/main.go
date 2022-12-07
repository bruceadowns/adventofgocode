package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Input []rune

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()
	for _, v := range line {
		res = append(res, v)
	}

	return
}

func findDistinct(in Input, iter int) (res int) {
	window := make(map[rune]int)

	//seed
	for i := 0; i < iter; i++ {
		window[in[i]]++
	}

	//iter
	for i := iter; i < len(in); i++ {
		if len(window) == iter {
			return i
		}

		window[in[i]]++
		window[in[i-iter]]--
		if window[in[i-iter]] == 0 {
			delete(window, in[i-iter])
		}
	}

	return
}

func Part1(in Input) (res int) {
	return findDistinct(in, 4)
}

func Part2(in Input) (res int) {
	return findDistinct(in, 14)
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
