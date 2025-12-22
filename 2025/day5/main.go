package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type ingredientRange struct {
	start, end int
}

func (ir ingredientRange) in(i int) (res bool) {
	if i >= ir.start && i <= ir.end {
		return true
	}

	return false
}

type Input struct {
	fresh []ingredientRange
	avail []int
}

func In(r io.Reader) (res Input) {
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
		res.fresh = append(res.fresh, ingredientRange{start, end})
	}

	for scanner.Scan() {
		line := scanner.Text()
		if num, err := strconv.Atoi(line); err == nil {
			res.avail = append(res.avail, num)
		} else {
			log.Fatal("invalid")
		}
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in.avail {
		for _, vv := range in.fresh {
			if vv.in(v) {
				res++
				break
			}
		}
	}

	return
}

func Part2(in Input) (res int) {
	uniq := make(map[int]struct{})
	for _, v := range in.fresh {
		for i := v.start; i <= v.end; i++ {
			uniq[i] = struct{}{}
		}
	}

	return len(uniq)
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
