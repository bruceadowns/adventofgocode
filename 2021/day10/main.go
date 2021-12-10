package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
)

type queue struct {
	data []rune
}

func (q *queue) push(r rune) {
	q.data = append(q.data, r)
}

func (q *queue) pop() (res rune) {
	if len(q.data) > 0 {
		res = q.data[len(q.data)-1]
		q.data = q.data[:len(q.data)-1]
	}

	return
}

func (q queue) empty() bool {
	return len(q.data) == 0
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
		var q queue
		for _, vv := range v {
			switch vv {
			case '(', '[', '{', '<':
				q.push(vv)
			case ')', ']', '}', '>':
				if q.pop() != pairs[vv] {
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
	var incompleteQueues []queue
outer:
	for _, v := range in {
		var q queue
		var fullQ queue
		for _, vv := range v {
			fullQ.push(vv)

			switch vv {
			case '(', '[', '{', '<':
				q.push(vv)
			case ')', ']', '}', '>':
				if q.pop() != pairs[vv] {
					continue outer
				}
			default:
				log.Fatal("invalid input")
			}
		}

		incompleteQueues = append(incompleteQueues, fullQ)
	}

	var completionQueues []queue
	for _, v := range incompleteQueues {
		var reverseQueue queue
		var completionQueue queue
		for p := v.pop(); p != 0; p = v.pop() {
			switch p {
			case ')', ']', '}', '>':
				reverseQueue.push(p)
			case '(', '[', '{', '<':
				if reverseQueue.empty() {
					completionQueue.push(pairs[p])
				} else {
					_ = reverseQueue.pop()
				}
			}
		}

		completionQueues = append(completionQueues, completionQueue)
	}

	var points = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	var completionScores []int
	for _, v := range completionQueues {
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
