package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Input map[int]int

func In(r io.Reader) (res Input) {
	res = make(Input)

	scanner := bufio.NewScanner(r)
	var elf int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elf++
			continue
		}

		var val int
		var err error
		if val, err = strconv.Atoi(line); err != nil {
			log.Fatal(err)
		}
		res[elf] += val
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in {
		if v > res {
			res = v
		}
	}

	return
}

func Part2(in Input) (res int) {
	values := make([]int, 0, len(in))
	for _, v := range in {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	return values[0] + values[1] + values[2]
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
