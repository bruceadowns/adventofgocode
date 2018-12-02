package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func part1(is []int) (res int) {
	for _, v := range is {
		res += v
	}

	return
}

func part2(is []int) (res int) {
	m := make(map[int]struct{})
	m[res] = struct{}{}

	for {
		for _, v := range is {
			res += v
			if _, ok := m[res]; ok {
				return
			}
			m[res] = struct{}{}
		}
	}
}

func in(r io.Reader) (res []int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, num)
	}

	return
}

func main() {
	is := in(os.Stdin)
	log.Printf("resulting freq: %d", part1(is))
	log.Printf("first duplicate freq: %d", part2(is))
}
