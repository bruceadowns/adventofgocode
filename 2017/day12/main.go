package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type initnode struct {
	pid   int
	cpids []int
}

type initnodes []*initnode

type graphnode struct {
	pid     int
	peer    []*graphnode
	visited bool
}

func (n *initnode) init(s string) (err error) {
	// 0 <-> 2
	// 2 <-> 0, 3, 4

	const sep = " <-> "
	if len(s) < len(sep)+2 {
		return fmt.Errorf("invalid input")
	}
	idx := strings.Index(s, sep)

	spid := s[:idx]
	n.pid, err = strconv.Atoi(spid)
	if err != nil {
		return err
	}

	sCpids := s[idx+len(sep):]
	for _, v := range strings.Split(sCpids, ",") {
		cpid, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return err
		}

		n.cpids = append(n.cpids, cpid)
	}

	return nil
}

func input(r io.Reader) (res initnodes) {
	res = make(initnodes, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		inode := &initnode{}
		inode.init(line)
		res = append(res, inode)
	}

	return
}

func makeGraph(n initnodes) (res map[int]*graphnode) {
	res = make(map[int]*graphnode, 0)
	for _, pinode := range n {
		pnode, ok := res[pinode.pid]
		if !ok {
			pnode = &graphnode{pid: pinode.pid}
			res[pinode.pid] = pnode
		}

		for _, cpid := range pinode.cpids {
			cnode, ok := res[cpid]
			if !ok {
				cnode = &graphnode{pid: cpid}
				res[cpid] = cnode
			}

			pnode.peer = append(pnode.peer, cnode)
			cnode.peer = append(cnode.peer, pnode)
		}
	}

	return
}

func visit(gnode *graphnode) (res int) {
	if gnode.visited {
		log.Fatal("invalid logic")
	}

	res = 1
	gnode.visited = true

	for _, cnode := range gnode.peer {
		if !cnode.visited {
			res += visit(cnode)
		}
	}

	return
}

func part1() {
	inodes := input(os.Stdin)
	gnodes := makeGraph(inodes)

	ppid := 0
	start, ok := gnodes[ppid]
	if !ok {
		log.Fatal("invalid input")
	}
	num := visit(start)
	if num < 1 {
		log.Fatal("invalid logic")
	}

	log.Printf("part1 %d pids in %d's pid graph", num, ppid)
}

func part2() {
	inodes := input(os.Stdin)
	gnodes := makeGraph(inodes)

	groups := 0
	for _, v := range gnodes {
		if !v.visited {
			visit(v)
			groups++
		}
	}
	if groups < 1 {
		log.Fatal("invalid logic")
	}

	log.Printf("part2 %d groups in graph", groups)
}

func main() {
	bPart1 := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		bPart1 = false
	}

	if bPart1 {
		part1()
	} else {
		part2()
	}
}
