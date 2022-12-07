package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type stack struct {
	data []uint8
}

func (s stack) copy() (res *stack) {
	data := make([]uint8, len(s.data))
	copy(data, s.data)

	return &stack{data}
}

func (s *stack) push(i uint8) {
	s.data = append(s.data, i)
}

func (s *stack) pushStack(i []uint8) {
	for j := len(i) - 1; j >= 0; j-- {
		s.data = append(s.data, i[j])
	}
}

func (s *stack) pop() (res uint8) {
	res = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return
}

func (s *stack) popStack(count int) (res []uint8) {
	for i := 0; i < count; i++ {
		res = append(res, s.data[len(s.data)-1])
		s.data = s.data[:len(s.data)-1]
	}

	return
}

func (s stack) peek() uint8 {
	return s.data[len(s.data)-1]
}

type instruction struct {
	count int
	from  int
	to    int
}

type Input struct {
	stacks       map[int]*stack
	instructions []instruction
}

func (in Input) copyStack() (res map[int]*stack) {
	res = make(map[int]*stack)
	for k, v := range in.stacks {
		res[k] = v.copy()
	}

	return
}

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)

	var rows []string
	var names []int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.IndexAny(line, "[") == -1 && strings.IndexAny(line, "]") == -1 {
			for _, v := range strings.Fields(line) {
				if n, err := strconv.Atoi(v); err == nil {
					names = append(names, n)
				} else {
					log.Fatal("invalid input")
				}
			}
			break
		} else {
			rows = append(rows, line)
		}
	}

	res.stacks = make(map[int]*stack)
	for i := len(rows) - 1; i >= 0; i-- {
		for j := 0; j < len(rows[i]); j += 4 {
			if rows[i][j] == '[' && rows[i][j+2] == ']' {
				if cur, exists := res.stacks[j/4+1]; exists {
					cur.push(rows[i][j+1])
				} else {
					res.stacks[j/4+1] = &stack{[]uint8{rows[i][j+1]}}
				}
			}
		}
	}

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()

		var count, from, to int
		n, err := fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
		if err != nil {
			log.Fatal(err)
		}
		if n != 3 {
			log.Fatal("invalid input")
		}

		res.instructions = append(res.instructions, instruction{count, from, to})
	}

	return
}

func orderedTips(i map[int]*stack) (res string) {
	stackKeys := make([]int, 0)
	for k := range i {
		stackKeys = append(stackKeys, k)
	}
	sort.Ints(stackKeys)

	var sb strings.Builder
	for _, k := range stackKeys {
		sb.WriteRune(rune(i[k].peek()))
	}
	return sb.String()
}

func Part1(in Input) (res string) {
	stacks := in.copyStack()
	for _, v := range in.instructions {
		for i := 0; i < v.count; i++ {
			stacks[v.to].push(stacks[v.from].pop())
		}
	}

	return orderedTips(stacks)
}

func Part2(in Input) (res string) {
	stacks := in.copyStack()
	for _, v := range in.instructions {
		stacks[v.to].pushStack(stacks[v.from].popStack(v.count))
	}

	return orderedTips(stacks)
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %s", Part1(i))
	log.Printf("part2: %s", Part2(i))
}
