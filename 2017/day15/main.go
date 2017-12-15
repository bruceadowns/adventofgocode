package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	iterCountPart1 = 40000000
	iterCountPart2 = 5000000
	aFactor        = 16807
	bFactor        = 48271
	remainder      = 2147483647
	lsb16          = 0xffff
)

func input(r io.Reader) (a, b int) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	n, err := fmt.Sscanf(scanner.Text(), "Generator A starts with %d", &a)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("invalid input")
	}

	scanner.Scan()
	n, err = fmt.Sscanf(scanner.Text(), "Generator B starts with %d", &b)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("invalid input")
	}

	return
}

func part1() {
	currA, currB := input(os.Stdin)

	count := 0
	for i := 0; i < iterCountPart1; i++ {
		currA = currA * aFactor % remainder
		currB = currB * bFactor % remainder
		if currA&lsb16 == currB&lsb16 {
			count++
		}
	}

	log.Printf("part judge matched %d in %d iterations", count, iterCountPart1)
}

func part2() {
	currA, currB := input(os.Stdin)

	count := 0
	for i := 0; i < iterCountPart2; i++ {
		for {
			currA = currA * aFactor % remainder
			if currA%4 == 0 {
				break
			}
		}

		for {
			currB = currB * bFactor % remainder
			if currB%8 == 0 {
				break
			}
		}

		if currA&lsb16 == currB&lsb16 {
			count++
		}
	}

	log.Printf("part2 judge matched %d in %d iterations", count, iterCountPart2)
}

func main() {
	bPart1 := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		bPart1 = false
	}

	if bPart1 {
		part1()
	} else {
		part2()
	}
}
