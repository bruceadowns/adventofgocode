package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Input []string

type (
	manifold  map[int]struct{}
	manifolds []manifold
)

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return
}

func Part1(in Input) (res int) {
	m := make(manifold)
	for x := 0; x < len(in[0]); x++ {
		switch in[0][x] {
		case '.':
		case 'S':
			m[x] = struct{}{}
		default:
			log.Fatal()
		}
	}

	for y := 1; y < len(in); y++ {
		for x := 0; x < len(in[y]); x++ {
			switch in[y][x] {
			case '.':
			case '^':
				if _, ok := m[x]; ok {
					delete(m, x)
					m[x-1] = struct{}{}
					m[x+1] = struct{}{}
					res++
				}
			default:
				log.Fatal()
			}
		}
	}

	return
}

func Part2(in Input) int {
	firstWorld := make(manifold)
	for x := 0; x < len(in[0]); x++ {
		switch in[0][x] {
		case '.':
		case 'S':
			firstWorld[x] = struct{}{}
		default:
			log.Fatal()
		}
	}

	var worlds manifolds
	worlds = append(worlds, firstWorld)

	for y := 1; y < len(in); y++ {
		for x := 0; x < len(in[y]); x++ {
			var newWorlds manifolds
			for _, v := range worlds {
				if _, ok := v[x]; ok {
					switch in[y][x] {
					case '.':
					case '^':
						delete(v, x)

						newWorld := make(manifold)
						for kk, vv := range v {
							newWorld[kk] = vv
						}
						newWorld[x+1] = struct{}{}
						newWorlds = append(newWorlds, newWorld)

						v[x-1] = struct{}{}
					default:
						log.Fatal()
					}
				}
			}

			if newWorlds != nil {
				worlds = append(worlds, newWorlds...)
			}
		}
	}

	return len(worlds)
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
