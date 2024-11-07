package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Input []string

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return
}

func Part1(in Input) (res int) {
	for _, v := range in {
		first := 10 * func(v string) int {
			for i := 0; i < len(v); i++ {
				switch v[i] {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					return int(v[i] - '0')
				}
			}
			return -1
		}(v)
		second := func(v string) int {
			for i := len(v) - 1; i > -1; i-- {
				switch v[i] {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					return int(v[i] - '0')
				}
			}
			return -1
		}(v)

		if first == -1 || second == -1 {
			log.Fatalf("invalid input %d:%d", first, second)
		}

		res += first + second
	}

	return
}

func Part2(in Input) (res int) {
	for _, v := range in {
		first := 10 * func(v string) int {
			for i := 0; i < len(v); i++ {
				if v[i] == 48 {
					return 0
				}
				if v[i] == 49 {
					return 1
				}
				if v[i] == 50 {
					return 2
				}
				if v[i] == 51 {
					return 3
				}
				if v[i] == 52 {
					return 4
				}
				if v[i] == 53 {
					return 5
				}
				if v[i] == 54 {
					return 6
				}
				if v[i] == 55 {
					return 7
				}
				if v[i] == 56 {
					return 8
				}
				if v[i] == 57 {
					return 9
				}
				if i <= len(v)-3 && v[i] == 'o' && v[i+1] == 'n' && v[i+2] == 'e' {
					return 1
				}
				if i <= len(v)-3 && v[i] == 't' && v[i+1] == 'w' && v[i+2] == 'o' {
					return 2
				}
				if i <= len(v)-3 && v[i] == 's' && v[i+1] == 'i' && v[i+2] == 'x' {
					return 6
				}
				if i <= len(v)-4 && v[i] == 'z' && v[i+1] == 'e' && v[i+2] == 'r' && v[i+3] == 'o' {
					return 0
				}
				if i <= len(v)-4 && v[i] == 'f' && v[i+1] == 'o' && v[i+2] == 'u' && v[i+3] == 'r' {
					return 4
				}
				if i <= len(v)-4 && v[i] == 'f' && v[i+1] == 'i' && v[i+2] == 'v' && v[i+3] == 'e' {
					return 5
				}
				if i <= len(v)-4 && v[i] == 'n' && v[i+1] == 'i' && v[i+2] == 'n' && v[i+3] == 'e' {
					return 9
				}
				if i <= len(v)-5 && v[i] == 't' && v[i+1] == 'h' && v[i+2] == 'r' && v[i+3] == 'e' && v[i+4] == 'e' {
					return 3
				}
				if i <= len(v)-5 && v[i] == 's' && v[i+1] == 'e' && v[i+2] == 'v' && v[i+3] == 'e' && v[i+4] == 'n' {
					return 7
				}
				if i <= len(v)-5 && v[i] == 'e' && v[i+1] == 'i' && v[i+2] == 'g' && v[i+3] == 'h' && v[i+4] == 't' {
					return 8
				}
			}
			return -1
		}(v)
		second := func(v string) int {
			for i := len(v) - 1; i > -1; i-- {
				if v[i] == 48 {
					return 0
				}
				if v[i] == 49 {
					return 1
				}
				if v[i] == 50 {
					return 2
				}
				if v[i] == 51 {
					return 3
				}
				if v[i] == 52 {
					return 4
				}
				if v[i] == 53 {
					return 5
				}
				if v[i] == 54 {
					return 6
				}
				if v[i] == 55 {
					return 7
				}
				if v[i] == 56 {
					return 8
				}
				if v[i] == 57 {
					return 9
				}
				if i <= len(v)-3 && v[i] == 'o' && v[i+1] == 'n' && v[i+2] == 'e' {
					return 1
				}
				if i <= len(v)-3 && v[i] == 't' && v[i+1] == 'w' && v[i+2] == 'o' {
					return 2
				}
				if i <= len(v)-3 && v[i] == 's' && v[i+1] == 'i' && v[i+2] == 'x' {
					return 6
				}
				if i <= len(v)-4 && v[i] == 'z' && v[i+1] == 'e' && v[i+2] == 'r' && v[i+3] == 'o' {
					return 0
				}
				if i <= len(v)-4 && v[i] == 'f' && v[i+1] == 'o' && v[i+2] == 'u' && v[i+3] == 'r' {
					return 4
				}
				if i <= len(v)-4 && v[i] == 'f' && v[i+1] == 'i' && v[i+2] == 'v' && v[i+3] == 'e' {
					return 5
				}
				if i <= len(v)-4 && v[i] == 'n' && v[i+1] == 'i' && v[i+2] == 'n' && v[i+3] == 'e' {
					return 9
				}
				if i <= len(v)-5 && v[i] == 't' && v[i+1] == 'h' && v[i+2] == 'r' && v[i+3] == 'e' && v[i+4] == 'e' {
					return 3
				}
				if i <= len(v)-5 && v[i] == 's' && v[i+1] == 'e' && v[i+2] == 'v' && v[i+3] == 'e' && v[i+4] == 'n' {
					return 7
				}
				if i <= len(v)-5 && v[i] == 'e' && v[i+1] == 'i' && v[i+2] == 'g' && v[i+3] == 'h' && v[i+4] == 't' {
					return 8
				}
			}
			return -1
		}(v)

		if first == -1 || second == -1 {
			log.Fatalf("invalid input %d:%d", first, second)
		}

		res += first + second
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
