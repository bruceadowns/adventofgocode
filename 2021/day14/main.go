package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type rule [2]uint8

type Manual struct {
	template []uint8
	rules    map[rule]uint8
}

func Part1(in Manual) (res int) {
	polymer := in.template

	for i := 0; i < 10; i++ {
		var newPolymer []uint8
		for j := 0; j < len(polymer)-1; j++ {
			newPolymer = append(newPolymer, polymer[j])

			if v, exists := in.rules[rule{polymer[j], polymer[j+1]}]; exists {
				newPolymer = append(newPolymer, v)
			}
		}
		polymer = append(newPolymer, polymer[len(polymer)-1])
	}

	counts := make(map[uint8]int)
	for _, v := range polymer {
		counts[v]++
	}

	min := math.MaxInt
	max := 0
	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min
}

func Part2(in Manual) (res int) {
	return
}

func In(r io.Reader) (res Manual) {
	res.rules = make(map[rule]uint8)

	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}

	line := scanner.Text()
	for i := 0; i < len(line); i++ {
		res.template = append(res.template, line[i])
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		var pair string
		var insert uint8
		n, err := fmt.Sscanf(line, "%s -> %c", &pair, &insert)
		if err != nil {
			log.Fatal("invalid input")
		}
		if n != 2 {
			log.Fatal("invalid input")
		}
		if len(pair) != 2 {
			log.Fatal("invalid input")
		}

		res.rules[rule{pair[0], pair[1]}] = insert
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
