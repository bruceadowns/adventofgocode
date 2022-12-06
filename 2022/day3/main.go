package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type items map[uint8]struct{}

type compartment struct {
	items items
}

type rucksack [2]compartment

func (r rucksack) all() (res []uint8) {
	for k := range r[0].items {
		res = append(res, k)
	}
	for k := range r[1].items {
		res = append(res, k)
	}

	return
}

func (r rucksack) exists(i uint8) bool {
	if _, exists := r[0].items[i]; exists {
		return true
	}
	if _, exists := r[1].items[i]; exists {
		return true
	}

	return false
}

type Input []rucksack

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var comp1 compartment
		comp1.items = make(items)
		for i := 0; i < len(line)/2; i++ {
			if _, exists := comp1.items[line[i]]; !exists {
				comp1.items[line[i]] = struct{}{}
			}
		}

		var comp2 compartment
		comp2.items = make(items)
		for i := len(line) / 2; i < len(line); i++ {
			if _, exists := comp2.items[line[i]]; !exists {
				comp2.items[line[i]] = struct{}{}
			}
		}

		res = append(res, rucksack{comp1, comp2})
	}

	return
}

func mapPriority(in uint8) (res int) {
	//Lowercase item types a through z have priorities 1 through 26.
	//Uppercase item types A through Z have priorities 27 through 52.
	//https://www.asciitable.com/

	if 'a' <= in && in <= 'z' {
		res = int(in - 'a' + 1)
	} else {
		res = int(in - 'A' + 1 + 26)
	}

	return
}

func Part1(in Input) (res int) {
outer:
	for _, v := range in {
		for k := range v[0].items {
			if _, exists := v[1].items[k]; exists {
				res += mapPriority(k)
				continue outer
			}
		}
	}

	return
}

func Part2(in Input) (res int) {
outer:
	for i := 0; i < len(in); i += 3 {
		for _, v := range in[i].all() {
			if in[i+1].exists(v) && in[i+2].exists(v) {
				res += mapPriority(v)
				continue outer
			}
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
