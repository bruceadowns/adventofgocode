package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type ingredientRange struct {
	start, end int
}
type ingredientRanges []*ingredientRange

func (ir ingredientRange) in(i int) (res bool) {
	if i >= ir.start && i <= ir.end {
		return true
	}

	return false
}

func (ir ingredientRanges) consolidate() (res ingredientRanges) {
	sort.Sort(ir)

	curr := ir[0]
	res = append(res, curr)

	for i := 1; i < len(ir); i++ {
		if ir[i].start <= curr.end {
			curr.end = max(curr.end, ir[i].end)
		} else {
			curr = ir[i]
			res = append(res, curr)
		}
	}

	return
}

func (ir ingredientRanges) Len() int {
	return len(ir)
}

func (ir ingredientRanges) Less(i, j int) bool {
	return ir[i].start < ir[j].start
}

func (ir ingredientRanges) Swap(i, j int) {
	ir[i], ir[j] = ir[j], ir[i]
}

type Input struct {
	fresh ingredientRanges
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
		res.fresh = append(res.fresh, &ingredientRange{start, end})
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
	for _, v := range in.fresh.consolidate() {
		res += v.end - v.start + 1
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
