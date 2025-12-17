package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input []int

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line := scanner.Text()

	for _, v := range strings.Split(line, ",") {
		var first, last int
		num, err := fmt.Sscanf(v, "%d-%d", &first, &last)
		if err != nil {
			log.Fatal()
		}
		if num != 2 {
			log.Fatal()
		}

		for i := first; i <= last; i++ {
			res = append(res, i)
		}
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in {
		s := strconv.Itoa(v)
		first := s[0 : len(s)/2]
		second := s[len(s)/2:]
		if strings.EqualFold(first, second) {
			res += v
		}
	}

	return
}

func Part2(in Input) (res int) {
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
