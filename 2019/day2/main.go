package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(in []int, p1, p2 int) int {
	mem := make([]int, len(in))
	copy(mem, in)

	mem[1] = p1
	mem[2] = p2

	var iPtr int
	for {
		opcode := mem[iPtr]
		switch opcode {
		case 1, 2:
			param1 := mem[iPtr+1]
			param2 := mem[iPtr+2]
			param3 := mem[iPtr+3]

			if opcode == 1 {
				mem[param3] = mem[param1] + mem[param2]
			} else if opcode == 2 {
				mem[param3] = mem[param1] * mem[param2]
			}

			iPtr += 4
		case 99:
			return mem[0]
		default:
			log.Fatal("invalid input")
		}
	}

	return 0
}

func part2(in []int) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			p1 := part1(in, noun, verb)

			if p1 == 19690720 {
				return 100*noun + verb
			}
		}
	}

	return 0
}

func getin(r io.Reader) (res [][]int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var line []int
		for _, v := range strings.Split(scanner.Text(), ",") {
			if n, err := strconv.Atoi(v); err == nil {
				line = append(line, n)
			} else {
				log.Fatal("invalid input")
			}
		}
		res = append(res, line)
	}

	return
}

func main() {
	i := getin(os.Stdin)
	for _, v := range i {
		log.Printf("part1: %d", part1(v, 12, 2))
		log.Printf("part2: %d", part2(v))
	}
}
