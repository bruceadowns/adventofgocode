package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const caseDiff rune = 'a' - 'A'

type node struct {
	next *node
	prev *node
	data rune
}

type list struct {
	first *node
	last  *node
}

func (l *list) String() string {
	var sb strings.Builder
	for n := l.first; n != nil; n = n.next {
		sb.WriteRune(n.data)
	}
	return sb.String()
}

func (l *list) insertAfter(n *node, nn *node) {
	nn.prev = n
	if n.next == nil {
		l.last = nn
	} else {
		nn.next = n.next
		n.next.prev = nn
	}
	n.next = nn
}

func (l *list) insertBefore(n *node, nn *node) {
	nn.next = n
	if n.prev == nil {
		l.first = nn
	} else {
		nn.prev = n.prev
		n.prev.next = nn
	}
	n.prev = nn
}

func (l *list) insertBeginning(nn *node) {
	if l.first == nil {
		l.first = nn
		l.last = nn
	} else {
		l.insertBefore(l.first, nn)
	}
}

func (l *list) insertEnd(nn *node) {
	if l.last == nil {
		l.insertBeginning(nn)
	} else {
		l.insertAfter(l.last, nn)
	}
}

func (l *list) remove(n *node) {
	if n.prev == nil {
		l.first = n.next
	} else {
		n.prev.next = n.next
	}
	if n.next == nil {
		l.last = n.prev
	} else {
		n.next.prev = n.prev
	}
}

func react(a, b rune) bool {
	if a > b {
		return a-b == caseDiff
	}
	return b-a == caseDiff
}

func build(s string) (res list) {
	for _, v := range s {
		res.insertEnd(&node{data: v})
	}

	return
}

func part1(l list) int {
	for n := l.first; n != nil && n.next != nil; {
		if react(n.data, n.next.data) {
			r := n

			// go back or skip forward
			if n.prev == nil {
				n = n.next.next
			} else {
				n = n.prev
			}

			l.remove(r.next)
			l.remove(r)
		} else {
			n = n.next
		}
	}

	return len(l.String())
}

func part2(s string) (res int) {
	res = math.MaxInt32
	for i := 'A'; i <= 'Z'; i++ {
		in := strings.NewReplacer(
			fmt.Sprintf("%c", i), "",
			fmt.Sprintf("%c", i+caseDiff), "").Replace(s)
		l := part1(build(in))
		if l < res {
			res = l
		}
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("reduced polymer length: %d", part1(build(line)))
		log.Printf("shortest reduced polymer length: %d", part2(line))
	}
}
