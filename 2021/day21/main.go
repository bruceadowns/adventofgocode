package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type player struct {
	position int
	score    int
}

type Input map[int]player

func normalize(in, mod int) (res int) {
	res = in % mod
	if res == 0 {
		res = mod
	}

	return
}

func Part1(in Input) (res int) {
	game := make(map[int]*player)
	for k, v := range in {
		game[k] = &player{v.position, v.score}
	}

	ddie := 0
	turn := 1
	loser := 0

	for {
		ddie++
		game[turn].position += ddie
		ddie++
		game[turn].position += ddie
		ddie++
		game[turn].position += ddie

		game[turn].score += normalize(normalize(game[turn].position, 100), 10)

		if game[turn].score >= 1000 {
			if turn == 1 {
				loser = 2
			} else {
				loser = 1
			}

			break
		}

		if turn == 1 {
			turn = 2
		} else {
			turn = 1
		}
	}

	return game[loser].score * ddie
}

func Part2(in Input) (res int) {
	return
}

func In(r io.Reader) (res Input) {
	res = make(Input)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var id int
		var position int
		n, err := fmt.Sscanf(line, "Player %d starting position: %d", &id, &position)
		if err != nil {
			log.Fatal("invalid input")
		}
		if n != 2 {
			log.Fatal("invalid input")
		}

		res[id] = player{position, 0}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
