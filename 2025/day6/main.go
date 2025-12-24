package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = iota
	MULTIPLY
)

type Equation struct {
	args []int
	op   int
}
type Input1 []Equation

type Card struct {
	minX int
	maxX int
	op   int
}
type Input2 struct {
	rlines []string
	cards  []Card
}

func In(r io.Reader) (res []string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return
}

func In1(lines []string) (res Input1) {
	for range strings.Fields(lines[1]) {
		res = append(res, Equation{})
	}

	i := 0
	for ; i < len(lines)-1; i++ {
		for k, v := range strings.Fields(lines[i]) {
			j, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal()
			}

			res[k].args = append(res[k].args, j)
		}
	}

	for k, v := range strings.Fields(lines[i]) {
		switch v {
		case "+":
			res[k].op = ADD
		case "*":
			res[k].op = MULTIPLY
		default:
			log.Fatal()
		}
	}

	return
}

func In2(lines []string) (res Input2) {
	for i := len(lines) - 2; i >= 0; i-- {
		res.rlines = append(res.rlines, lines[i])
	}
	ops := lines[len(lines)-1]

	var indices []int
	for k, v := range ops {
		switch v {
		case '+', '*':
			indices = append(indices, k)
		}
	}
	indices = append(indices, len(ops)+1)

	for i := 0; i < len(indices)-1; i++ {
		switch ops[indices[i]] {
		case '+':
			res.cards = append(res.cards, Card{indices[i], indices[i+1] - 2, ADD})
		case '*':
			res.cards = append(res.cards, Card{indices[i], indices[i+1] - 2, MULTIPLY})
		}
	}

	return
}

func Part1(in Input1) (res int) {
	for _, v := range in {
		var total int
		switch v.op {
		case ADD:
			for _, vv := range v.args {
				total += vv
			}
		case MULTIPLY:
			total = 1
			for _, vv := range v.args {
				total *= vv
			}
		default:
			log.Fatal()
		}

		res += total
	}

	return
}

func Part2(in Input2) (res int) {
	for _, v := range in.cards {
		var nums []int
		for x := v.maxX; x >= v.minX; x-- {
			var num int
			for y := len(in.rlines) - 1; y >= 0; y-- {
				n := in.rlines[y][x]
				switch n {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					num = num*10 + int(n-'0')
				}
			}

			nums = append(nums, num)
		}

		var total int
		switch v.op {
		case ADD:
			for _, v := range nums {
				total += v
			}
		case MULTIPLY:
			total = 1
			for _, v := range nums {
				total *= v
			}
		}

		res += total
	}

	return
}

func Part21(in Input2) (res int) {
	for _, v := range in.cards {
		var nums []int
		for y := len(in.rlines) - 1; y >= 0; y-- {
			var num int
			for x := v.minX; x <= v.maxX; x++ {
				n := in.rlines[y][x]
				switch n {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					num = num*10 + int(n-'0')
				}
			}

			nums = append(nums, num)
		}

		var total int
		switch v.op {
		case ADD:
			for _, v := range nums {
				total += v
			}
		case MULTIPLY:
			total = 1
			for _, v := range nums {
				total *= v
			}
		}

		res += total
	}

	return
}

func main() {
	lines := In(os.Stdin)

	in1 := In1(lines)
	log.Printf("part1: %d", Part1(in1))

	in2 := In2(lines)
	log.Printf("part2: %d", Part2(in2))

	log.Printf("part21: %d", Part21(in2))
}
