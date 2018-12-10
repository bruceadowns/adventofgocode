package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type input struct {
	playerCount, lastMarble int
}

func in(r io.Reader) (res []input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		i := input{}
		num, err := fmt.Sscanf(
			scanner.Text(),
			"%d players; last marble is worth %d points",
			&i.playerCount, &i.lastMarble)
		if err != nil {
			log.Fatal(err)
		}
		if num != 2 {
			log.Fatal("invalid input")
		}

		res = append(res, i)
	}

	return
}

type node struct {
	data int
	next *node
	prev *node
}

func (n *node) insertAfter(nn *node) (res *node) {
	res = nn
	res.prev = n
	res.next = n.next
	n.next.prev = nn
	n.next = nn

	return
}

func (n *node) remove() (res *node) {
	res = n.next
	n.next.prev = n.prev
	n.prev.next = n.next

	return
}

func (n *node) inc(idx int) (res *node) {
	res = n
	for i := 0; i < idx; i++ {
		res = res.next
	}

	return
}

func (n *node) dec(idx int) (res *node) {
	res = n
	for i := 0; i < idx; i++ {
		res = res.prev
	}

	return
}

func winScore(in input) (res int) {
	curr := &node{data: 0}
	curr.next = curr
	curr.prev = curr

	players := make(map[int]int)
	for i := 1; i <= in.lastMarble; i++ {
		if i%23 == 0 {
			player := i % in.playerCount
			players[player] += i
			curr = curr.dec(7)
			players[player] += curr.data
			curr = curr.remove()
		} else {
			curr = curr.inc(1).insertAfter(&node{data: i})
		}
	}

	res = math.MinInt32
	for _, v := range players {
		if v > res {
			res = v
		}
	}

	return
}

func main() {
	i := in(os.Stdin)
	for _, v := range i {
		log.Printf("winning elf score: %d", winScore(v))
	}
}
