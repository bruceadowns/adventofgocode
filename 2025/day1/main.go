package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type rotation struct {
	direction rune
	scalar    int
}

type Input []rotation

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var direction rune
		var scalar int
		num, err := fmt.Sscanf(line, "%c%d", &direction, &scalar)
		if err != nil {
			log.Fatal("invalid")
		}
		if num != 2 {
			log.Fatal("invalid")
		}

		res = append(res, rotation{direction, scalar})
	}

	return
}

func Part1(in Input) (res int) {
	// count how many times we land on zero
	curr := 50

	for _, v := range in {
		switch v.direction {
		case 'L':
			curr -= v.scalar
			for curr < 0 {
				curr += 100
			}
		case 'R':
			curr += v.scalar
			for curr > 99 {
				curr -= 100
			}
		default:
			log.Fatal("invalid")
		}

		if curr == 0 {
			res++
		}
	}

	return
}

func Part2(in Input) (res int) {
	// count how many times we pass zero
	curr := 50

	for _, v := range in {
		for i := v.scalar; i > 0; i-- {
			switch v.direction {
			case 'L':
				curr--
			case 'R':
				curr++
			default:
				log.Fatal("invalid")
			}

			if curr == -1 {
				curr = 99
			} else if curr == 100 {
				curr = 0
			}
			if curr == 0 {
				res++
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
