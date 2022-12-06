package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type compartment struct {
	items map[uint8]int
}

type rucksack [2]compartment

type Input []rucksack

type group [3]compartment
type Input2 []group

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var comp1 compartment
		comp1.items = make(map[uint8]int)
		for i := 0; i < len(line)/2; i++ {
			comp1.items[line[i]]++
		}

		var comp2 compartment
		comp2.items = make(map[uint8]int)
		for i := len(line) / 2; i < len(line); i++ {
			comp2.items[line[i]]++
		}

		res = append(res, rucksack{comp1, comp2})
	}

	return
}

func In2(r io.Reader) (res Input2) {
	scanner := bufio.NewScanner(r)

outer:
	for {
		var group group
		for i := 0; i < 3; i++ {
			if !scanner.Scan() {
				break outer
			}
			line := scanner.Text()

			group[i].items = make(map[uint8]int)
			for j := 0; j < len(line); j++ {
				group[i].items[line[j]]++
			}
		}
		res = append(res, group)
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

func Part2(in Input2) (res int) {
outer:
	for _, v := range in {
		for k := range v[0].items {
			_, exists1 := v[1].items[k]
			_, exists2 := v[2].items[k]
			if exists1 && exists2 {
				res += mapPriority(k)
				continue outer
			}
		}
	}

	return
}

func main() {
	//i := In(os.Stdin)
	//log.Printf("part1: %d", Part1(i))

	i2 := In2(os.Stdin)
	log.Printf("part2: %d", Part2(i2))
}
