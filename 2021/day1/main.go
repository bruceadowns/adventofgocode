package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func part1(in []int) (res int) {
	for i := 0; i < len(in)-1; i++ {
		if in[i] < in[i+1] {
			res++
		}
	}

	return
}

func part2(in []int) (res int) {
	for i := 0; i < len(in)-3; i++ {
		if in[i]+in[i+1]+in[i+2] < in[i+1]+in[i+2]+in[i+3] {
			res++
		}
	}

	return
}

func in(r io.Reader) (res []int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, i)
	}

	return
}

func main() {
	i := in(os.Stdin)
	log.Printf("part1: %d", part1(i))
	log.Printf("part2: %d", part2(i))
}
