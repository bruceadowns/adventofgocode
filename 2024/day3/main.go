package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

type Input []string

func In(r io.Reader) (res Input) {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		res = append(res, re.FindAllString(scanner.Text(), -1)...)
	}

	return
}

func multiply(s string) int {
	var one, two int
	num, err := fmt.Sscanf(s, "mul(%d,%d)", &one, &two)
	if err != nil {
		log.Fatal(err)
	}
	if num != 2 {
		log.Fatal("invalid input")
	}

	return one * two
}

func Part1(in Input) (res int) {
	for _, v := range in {
		switch v {
		case "do()":
		case "don't()":
			break
		default:
			res += multiply(v)
		}
	}

	return
}

func Part2(in Input) (res int) {
	do := true
	for _, v := range in {
		switch v {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if do {
				res += multiply(v)
			}
		}
	}

	return
}

func main() {
	input := In(os.Stdin)
	log.Printf("part1: %d", Part1(input))
	log.Printf("part2: %d", Part2(input))
}
