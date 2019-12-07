package main

import (
	"fmt"
	"log"
)

func part1(low, high int) (res int) {
outer:
	for i := low; i <= high; i++ {
		s := fmt.Sprintf("%d", i)
		if len(s) != 6 {
			log.Fatal("invalid input")
		}

		var adj bool
		for j := 0; j < len(s)-1; j++ {
			if s[j] == s[j+1] {
				adj = true
			} else if s[j] > s[j+1] {
				continue outer
			}
		}

		if adj {
			res++
		}
	}

	return
}

func part2(low, high int) (res int) {
outer:
	for i := low; i <= high; i++ {
		s := fmt.Sprintf("%d", i)
		if len(s) != 6 {
			log.Fatal("invalid input")
		}

		m := make(map[uint8]int)
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				continue outer
			}

			m[s[j]]++
		}
		m[s[len(s)-1]]++

		for _, v := range m {
			if v == 2 {
				res++
				break
			}
		}
	}

	return
}

func main() {
	log.Printf("part1: %d", part1(136760, 595730))
	log.Printf("part2: %d", part2(136760, 595730))
}
