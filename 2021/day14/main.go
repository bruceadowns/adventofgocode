package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type pair [2]uint8

func (p pair) String() string {
	return fmt.Sprintf("%c%c", p[0], p[1])
}

type Manual struct {
	template []uint8
	rules    map[pair]uint8
}

func Part1(in Manual) (res int) {
	polymer := in.template

	for i := 0; i < 10; i++ {
		var newPolymer []uint8
		for j := 0; j < len(polymer)-1; j++ {
			newPolymer = append(newPolymer, polymer[j])

			if v, exists := in.rules[pair{polymer[j], polymer[j+1]}]; exists {
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
	pairs := make(map[pair]int)
	for i := 0; i < len(in.template)-1; i++ {
		pairs[pair{in.template[i], in.template[i+1]}]++
	}

	for i := 0; i < 40; i++ {
		newPairs := make(map[pair]int)
		for k, v := range pairs {
			newPairs[k] = v
		}

		for k, v := range pairs {
			if vv, exists := in.rules[k]; exists {
				newPairs[k] -= v
				newPairs[pair{k[0], vv}] += v
				newPairs[pair{vv, k[1]}] += v
			}
		}

		pairs = newPairs
	}

	counts := make(map[uint8]int)
	for k, v := range pairs {
		counts[k[0]] += v
		counts[k[1]] += v
	}
	counts[in.template[0]]++
	counts[in.template[len(in.template)-1]]++

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

	return max/2 - min/2
}

func In(r io.Reader) (res Manual) {
	res.rules = make(map[pair]uint8)

	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}

	line := scanner.Text()
	for i := 0; i < len(line); i++ {
		res.template = append(res.template, line[i])
	}

	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			continue
		}

		var p string
		var i uint8
		n, err := fmt.Sscanf(line, "%s -> %c", &p, &i)
		if err != nil {
			log.Fatal("invalid input")
		}
		if n != 2 {
			log.Fatal("invalid input")
		}
		if len(p) != 2 {
			log.Fatal("invalid input")
		}

		res.rules[pair{p[0], p[1]}] = i
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
