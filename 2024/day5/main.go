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

type (
	Input struct {
		rules   map[int]map[int]struct{}
		updates [][]int
	}
)

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)

	res.rules = make(map[int]map[int]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		// 47|53
		var pageNumberX, pageNumberY int
		num, err := fmt.Sscanf(line, "%d|%d", &pageNumberX, &pageNumberY)
		if err != nil {
			log.Fatal(err)
		}
		if num != 2 {
			log.Fatal("invalid input")
		}

		if _, ok := res.rules[pageNumberX]; !ok {
			res.rules[pageNumberX] = make(map[int]struct{})
		}
		res.rules[pageNumberX][pageNumberY] = struct{}{}
	}

	for scanner.Scan() {
		line := scanner.Text()

		// 75,47,61,53,29
		var fields []int
		for _, v := range strings.Split(line, ",") {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			fields = append(fields, i)
		}
		res.updates = append(res.updates, fields)
	}

	return
}

func correct(u []int, r map[int]map[int]struct{}) bool {
	for i := 0; i < len(u); i++ {
		for j := i + 1; j < len(u); j++ {
			if _, ok := r[u[j]][u[i]]; ok {
				return false
			}
		}
	}

	return true
}

func reorder(u []int, r map[int]map[int]struct{}) (res []int) {
	res = make([]int, len(u))
	copy(res, u)

	for i := 0; i < len(u); i++ {
		for j := i + 1; j < len(u); j++ {
			if _, ok := r[u[j]][u[i]]; ok {
				res[i], res[j] = res[j], res[i]
				return
			}
		}
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in.updates {
		if correct(v, in.rules) {
			res += v[len(v)/2]
		}
	}

	return
}

func Part2(in Input) (res int) {
	for _, v := range in.updates {
		if !correct(v, in.rules) {
			vv := v
			for {
				vv = reorder(vv, in.rules)
				if correct(vv, in.rules) {
					break
				}
			}
			res += vv[len(v)/2]
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
