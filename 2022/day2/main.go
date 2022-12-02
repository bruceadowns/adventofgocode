package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	rock = iota + 1
	paper
	scissors
)

type round struct {
	left  int
	right int
}

func (r round) score() int {
	// Rock defeats Scissors
	// Paper defeats Rock
	// Scissors defeats Paper

	// 0 if you lost
	// 3 if the round was a draw
	// and 6 if you won

	var score int
	if r.right == r.left {
		score = 3
	} else if (r.right == rock && r.left == scissors) ||
		(r.right == paper && r.left == rock) ||
		(r.right == scissors && r.left == paper) {
		score = 6
	} else {
		score = 0
	}

	return score + r.right
}

func (r round) prescore() (res round) {
	// X (rock) means you need to lose
	// Y (paper) means you need to end the round in a draw
	// Z (scissors) means you need to win

	res.left = r.left
	if r.right == rock {
		switch r.left {
		case rock:
			res.right = scissors
		case paper:
			res.right = rock
		case scissors:
			res.right = paper
		}
	} else if r.right == paper {
		res.right = r.left
	} else {
		switch r.left {
		case rock:
			res.right = paper
		case paper:
			res.right = scissors
		case scissors:
			res.right = rock
		}
	}

	return
}

type Input []round

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var player1, player2 rune
		if n, err := fmt.Sscanf(line, "%c %c", &player1, &player2); err == nil && n == 2 {
			var playerMap1 int
			switch player1 {
			case 'A':
				playerMap1 = rock
			case 'B':
				playerMap1 = paper
			case 'C':
				playerMap1 = scissors
			default:
				log.Fatal("invalid input")
			}

			var playerMap2 int
			switch player2 {
			case 'X':
				playerMap2 = rock
			case 'Y':
				playerMap2 = paper
			case 'Z':
				playerMap2 = scissors
			default:
				log.Fatal("invalid input")
			}

			res = append(res, round{playerMap1, playerMap2})
		}
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in {
		res += v.score()
	}

	return
}

func Part2(in Input) (res int) {
	for _, v := range in {
		res += v.prescore().score()
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
