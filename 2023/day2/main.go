package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Turn map[string]int
type Game []Turn
type Input map[int]Game

var maxTurn = Turn{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func In(r io.Reader) (res Input) {
	res = Input{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		gamesSplit := strings.Split(line, ":")
		if len(gamesSplit) != 2 {
			log.Fatal("invalid input")
		}

		var gameNumber int
		num, err := fmt.Sscanf(gamesSplit[0], "Game %d", &gameNumber)
		if err != nil {
			log.Fatal("invalid input")
		}
		if num != 1 {
			log.Fatal("invalid input")
		}

		var game Game
		gameSetsSplit := strings.Split(gamesSplit[1], ";")
		for _, gameSets := range gameSetsSplit {
			turn := Turn{}
			setSplit := strings.Split(gameSets, ",")
			for _, set := range setSplit {
				var count int
				var color string
				num, err := fmt.Sscanf(set, "%d %s", &count, &color)
				if err != nil {
					log.Fatal("invalid input")
				}
				if num != 2 {
					log.Fatal("invalid input")
				}

				turn[color] = count
			}

			game = append(game, turn)
		}

		res[gameNumber] = game
	}

	return
}

func Part1(in Input, maxTurn Turn) (res int) {
	for k, v := range in {
		if func() bool {
			for _, vv := range v {
				for kkk, vvv := range vv {
					if vvvv, ok := maxTurn[kkk]; ok {
						if vvv > vvvv {
							return false
						}
					} else {
						return false
					}
				}
			}
			return true
		}() {
			res += k
		}
	}

	return
}

func Part2(in Input) (res int) {
	for _, v := range in {
		var maxCubes = make(map[string]int)
		for _, vv := range v {
			for kkk, vvv := range vv {
				if vvvv, ok := maxCubes[kkk]; ok {
					if vvv > vvvv {
						maxCubes[kkk] = vvv
					}
				} else {
					maxCubes[kkk] = vvv
				}
			}
		}

		var power = 1
		for _, v := range maxCubes {
			power *= v
		}
		res += power
	}

	return
}

func main() {
	in := In(os.Stdin)
	log.Printf("part1: %d", Part1(in, maxTurn))
	log.Printf("part2: %d", Part2(in))
}
