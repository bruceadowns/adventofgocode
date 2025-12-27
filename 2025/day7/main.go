package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Input []string

type manifold map[int]int

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
			m[x] = 1
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
					m[x-1] = 1
					m[x+1] = 1
					delete(m, x)
					res++
				}
			default:
				log.Fatal()
			}
		}
	}

	return
}

func Part2(in Input) (res int) {
	m := make(manifold)
	for x := 0; x < len(in[0]); x++ {
		switch in[0][x] {
		case '.':
		case 'S':
			m[x] = 1
		default:
			log.Fatal()
		}
	}

	for y := 1; y < len(in); y++ {
		for x := 0; x < len(in[y]); x++ {
			switch in[y][x] {
			case '.':
			case '^':
				m[x-1] += m[x]
				m[x+1] += m[x]
				delete(m, x)
			default:
				log.Fatal()
			}
		}
	}

	for _, v := range m {
		res += v
	}
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
