package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	forward = iota
	down
	up
)

type direction struct {
	baring int
	vector int
}
type course []direction
type position struct {
	horizontal int
	depth      int
	aim        int
}

func part1(in course) (res int) {
	var curr position
	for _, dir := range in {
		switch dir.baring {
		case forward:
			curr.horizontal += dir.vector
		case down:
			curr.depth += dir.vector
		case up:
			curr.depth -= dir.vector
		}
	}

	return curr.horizontal * curr.depth
}

func part2(in course) (res int) {
	var curr position
	for _, dir := range in {
		switch dir.baring {
		case forward:
			curr.horizontal += dir.vector
			curr.depth += curr.aim * dir.vector
		case down:
			curr.aim += dir.vector
		case up:
			curr.aim -= dir.vector
		}
	}

	return curr.horizontal * curr.depth
}

func in(r io.Reader) (res course) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var sBaring string
		var vector int
		n, err := fmt.Sscanf(line, "%s %d", &sBaring, &vector)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatalf("invalid input: %s", line)
		}

		var iBaring int
		switch sBaring {
		case "forward":
			iBaring = forward
		case "down":
			iBaring = down
		case "up":
			iBaring = up
		default:
			log.Fatalf("unknown baring %s", sBaring)
		}

		res = append(res, direction{iBaring, vector})
	}

	return
}

func main() {
	i := in(os.Stdin)
	log.Printf("part1: %d", part1(i))
	log.Printf("part2: %d", part2(i))
}
