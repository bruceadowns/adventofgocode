package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	children []node
	metadata []int
}

func build(in []int, idx *int) (res node) {
	numChildren := in[*idx]
	*idx++
	numMetadata := in[*idx]
	*idx++

	for i := 0; i < numChildren; i++ {
		res.children = append(res.children, build(in, idx))
	}

	for i := 0; i < numMetadata; i++ {
		res.metadata = append(res.metadata, in[*idx])
		*idx++
	}

	return
}

func count(n node) (res int) {
	for _, v := range n.metadata {
		res += v
	}

	for _, v := range n.children {
		res += count(v)
	}

	return
}

func value(n node) (res int) {
	if len(n.children) == 0 {
		for _, v := range n.metadata {
			res += v
		}
	} else {
		for _, v := range n.metadata {
			if v <= len(n.children) {
				res += value(n.children[v-1])
			}
		}
	}

	return
}

func part1(in []int) (res int) {
	var idx int
	root := build(in, &idx)
	if idx != len(in) {
		log.Fatal("invalid logic")
	}

	res = count(root)

	return
}

func part2(in []int) (res int) {
	var idx int
	root := build(in, &idx)
	if idx != len(in) {
		log.Fatal("invalid logic")
	}

	res = value(root)

	return
}

func in(r io.Reader) (res []int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for _, v := range strings.Fields(scanner.Text()) {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			res = append(res, n)
		}
	}

	return
}

func main() {
	i := in(os.Stdin)
	log.Printf("sum of all metadata entries: %d", part1(i))
	log.Printf("value of the root node: %d", part2(i))
}
