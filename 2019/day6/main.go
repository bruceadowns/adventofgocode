package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type node struct {
	parent   *node
	children []*node
}

func depth(n *node) (res int) {
	for t := n; t.parent != nil; t = t.parent {
		res++
	}

	return
}

func part1(in map[string]*node) (res int) {
	for _, v := range in {
		res += depth(v)
	}

	return
}

func part2(you, san *node) (res int) {
	var com *node
outer:
	for n := you; n.parent != nil; n = n.parent {
		for m := san; m.parent != nil; m = m.parent {
			if n == m {
				com = n
				break outer
			}
		}
	}

	youDepth := depth(you)
	sanDepth := depth(san)
	comDepth := depth(com)

	return (youDepth - comDepth - 1) + (sanDepth - comDepth - 1)
}

func getin(r io.Reader) (res map[string]*node) {
	res = make(map[string]*node)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		in := strings.Split(scanner.Text(), ")")
		if len(in) != 2 {
			log.Fatal("invalid input")
		}

		parent, okParent := res[in[0]]
		if !okParent {
			parent = &node{}
		}
		res[in[0]] = parent

		child, okChild := res[in[1]]
		if !okChild {
			child = &node{}
		}
		res[in[1]] = child

		child.parent = parent
		parent.children = append(parent.children, child)
	}

	return
}

func main() {
	in := getin(os.Stdin)
	log.Printf("part1: %d", part1(in))
	log.Printf("part2: %d", part2(in["YOU"], in["SAN"]))
}
