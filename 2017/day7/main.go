package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type initnode struct {
	name     string
	weight   int
	children []string
}

type initnodes []*initnode

func (inode *initnode) init(line string) {
	i, err := fmt.Sscanf(line, "%s (%d)", &inode.name, &inode.weight)
	if i != 2 {
		log.Fatalf("invalid input: %s", line)
	}
	if err != nil {
		log.Fatalf("invalid input: %s [%v]", line, err)
	}

	idx := strings.Index(line, " -> ")
	if idx != -1 {
		s := line[idx+len(" -> "):]
		if s == "" {
			log.Fatalf("invalid input: %s", line)
		}

		children := strings.Split(s, ",")
		if len(children) < 1 {
			log.Fatalf("invalid input: %s", line)
		}

		for _, child := range children {
			inode.children = append(inode.children, strings.TrimSpace(child))
		}
	}
}

func getInput(r io.Reader) (res initnodes) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		in := &initnode{}
		in.init(scanner.Text())
		res = append(res, in)
	}

	return
}

type treenode struct {
	name       string
	weight     int
	treeWeight int
	parent     *treenode
	children   []*treenode
}

func (tnode *treenode) String() string {
	return fmt.Sprintf("name: %s weight: %d treeWeight: %d",
		tnode.name, tnode.weight, tnode.treeWeight)
}

func (tnode *treenode) calcTreeWeight() (res int) {
	res = tnode.weight
	for _, child := range tnode.children {
		res += child.calcTreeWeight()
	}

	tnode.treeWeight = res
	return
}

func (tnode *treenode) unbalancedNode() (*treenode, int) {
	siblingWeights := make(map[int]int)
	for _, child := range tnode.children {
		siblingWeights[child.treeWeight]++
	}

	if len(siblingWeights) == 1 {
		return nil, 0
	} else if len(siblingWeights) == 2 {
		var uniqWeight, siblingWeight int
		for k, v := range siblingWeights {
			if v == 1 {
				uniqWeight = k
			} else {
				siblingWeight = k
			}
		}
		if uniqWeight == 0 || siblingWeight == 0 {
			log.Fatalf("invalid input: %s", tnode.name)
		}

		var uniqSibling *treenode
		for _, child := range tnode.children {
			if child.treeWeight == uniqWeight {
				uniqSibling = child
				break
			}
		}

		if uniqSiblingChild, uniqSiblingChildWeight := uniqSibling.unbalancedNode(); uniqSiblingChild != nil {
			return uniqSiblingChild, uniqSiblingChildWeight
		} else {
			return uniqSibling, siblingWeight
		}
	} else {
		log.Fatalf("invalid input: %s", tnode.name)
	}

	return nil, 0
}

func hasParent(node *initnode, nodes initnodes) bool {
	for _, inode := range nodes {
		for _, child := range inode.children {
			if strings.EqualFold(child, node.name) {
				return true
			}
		}
	}

	return false
}

func findRoot(inodes initnodes) *initnode {
	for _, inode := range inodes {
		if !hasParent(inode, inodes) {
			return inode
		}
	}

	return nil
}

func makeTree(child string, inodes initnodes) (tnode *treenode) {
	for _, inode := range inodes {
		if strings.EqualFold(child, inode.name) {
			tnode = &treenode{name: inode.name, weight: inode.weight}
			for _, child := range inode.children {
				tnode.children = append(tnode.children, makeTree(child, inodes))
			}

			break
		}
	}

	return
}

func part1() {
	input := getInput(os.Stdin)
	iroot := findRoot(input)
	if iroot == nil {
		log.Fatal("no root found")
	}

	log.Printf("part1 root: %s", iroot.name)
}

func part2() {
	input := getInput(os.Stdin)

	iroot := findRoot(input)
	if iroot == nil {
		log.Fatal("no root found")
	}

	troot := &treenode{name: iroot.name, weight: iroot.weight}
	for _, ichild := range iroot.children {
		troot.children = append(troot.children, makeTree(ichild, input))
	}

	troot.calcTreeWeight()

	unode, balancedTreeWeight := troot.unbalancedNode()
	log.Printf("part2 unbalanced node: [%s]: weight should be: %d",
		unode, unode.weight-(unode.treeWeight-balancedTreeWeight))
}

func main() {
	bPart1 := false
	if len(os.Args) == 2 && os.Args[1] == "1" {
		bPart1 = true
	}

	if bPart1 {
		part1()
	} else {
		part2()
	}
}
