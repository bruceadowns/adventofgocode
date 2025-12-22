package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Input struct {
	fresh map[int]struct{}
	avail []int
}

func In(r io.Reader) (res Input) {
	res.fresh = make(map[int]struct{})

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var start, end int
		num, err := fmt.Sscanf(line, "%d-%d", &start, &end)
		if err != nil {
			log.Fatal("invalid")
		}
		if num != 2 {
			log.Fatal("invalid")
		}
		log.Printf("start %d end %d", start, end)
		for i := start; i <= end; i++ {
			res.fresh[i] = struct{}{}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if num, err := strconv.Atoi(line); err == nil {
			res.avail = append(res.avail, num)
		} else {
			log.Fatal("invalid")
		}
	}

	log.Print(len(res.fresh))
	log.Print(len(res.avail))

	return
}

func Part1(in Input) (res int) {
	for _, v := range in.avail {
		if _, ok := in.fresh[v]; ok {
			res++
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
