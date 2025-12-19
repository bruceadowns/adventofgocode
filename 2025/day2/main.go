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

type Item struct {
	i int
	s string
}

type Input []Item

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
			res = append(res, Item{i, strconv.Itoa(i)})
		}
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in {
		first := v.s[0 : len(v.s)/2]
		second := v.s[len(v.s)/2:]
		if strings.EqualFold(first, second) {
			res += v.i
		}
	}

	return
}

func Part2(in Input) (res int) {
v:
	for _, v := range in {
	i:
		for i := 1; i < len(v.s)/2+1; i++ {
			if len(v.s)%i != 0 {
				continue
			}

			curr := v.s[0:i]
			for j := i; j < len(v.s)-len(curr)+1; j += len(curr) {
				comp := v.s[j : j+len(curr)]
				if curr != comp {
					continue i
				}
			}

			res += v.i
			continue v
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
