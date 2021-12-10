package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
)

type stack struct {
	data []rune
}

func (s *stack) push(r rune) {
	s.data = append(s.data, r)
}

func (s *stack) pop() (res rune) {
	if len(s.data) > 0 {
		res = s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
	}

	return
}

func (s stack) empty() bool {
	return len(s.data) == 0
}

var pairs = map[rune]rune{
	'(': ')',
	')': '(',
	'[': ']',
	']': '[',
	'{': '}',
	'}': '{',
	'<': '>',
	'>': '<',
}

func Part1(in []string) (res int) {
	var illegals []rune
outer:
	for _, v := range in {
		var s stack
		for _, vv := range v {
			switch vv {
			case '(', '[', '{', '<':
				s.push(vv)
			case ')', ']', '}', '>':
				if s.pop() != pairs[vv] {
					illegals = append(illegals, vv)
					continue outer
				}
			default:
				log.Fatal("invalid input")
			}
		}
	}

	var points = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for _, v := range illegals {
		res += points[v]
	}

	return
}

func Part2(in []string) (res int) {
	var incompleteStacks []stack
outer:
	for _, v := range in {
		var s stack
		var fullStack stack
		for _, vv := range v {
			fullStack.push(vv)

			switch vv {
			case '(', '[', '{', '<':
				s.push(vv)
			case ')', ']', '}', '>':
				if s.pop() != pairs[vv] {
					continue outer
				}
			default:
				log.Fatal("invalid input")
			}
		}

		incompleteStacks = append(incompleteStacks, fullStack)
	}

	var completionStacks []stack
	for _, v := range incompleteStacks {
		var reverseStack stack
		var completionStack stack
		for p := v.pop(); p != 0; p = v.pop() {
			switch p {
			case ')', ']', '}', '>':
				reverseStack.push(p)
			case '(', '[', '{', '<':
				if reverseStack.empty() {
					completionStack.push(pairs[p])
				} else {
					_ = reverseStack.pop()
				}
			}
		}

		completionStacks = append(completionStacks, completionStack)
	}

	var points = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	var completionScores []int
	for _, v := range completionStacks {
		var completionScore int
		for _, vv := range v.data {
			completionScore *= 5
			completionScore += points[vv]
		}
		completionScores = append(completionScores, completionScore)
	}

	sort.Ints(completionScores)
	middle := len(completionScores) / 2

	return completionScores[middle]
}

func In(r io.Reader) (res []string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
