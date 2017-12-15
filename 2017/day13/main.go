package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type direction int

const (
	up direction = iota
	down
)

type initLayer struct {
	fwdepth int
	fwrange int
}

func (il *initLayer) init(s string) {
	n, err := fmt.Sscanf(s, "%d: %d", &il.fwdepth, &il.fwrange)
	if err != nil {
		log.Fatal(err)
	}
	if n != 2 {
		log.Fatal("invalid input")
	}
}

type layer struct {
	curr    int
	scanDir direction
	rnge    []bool
}

func (l *layer) init(i int) {
	l.rnge = make([]bool, i)
	l.rnge[0] = true
}

type firewall struct {
	packet int
	layers []*layer
}

func (fw *firewall) init(i int) {
	fw.layers = make([]*layer, i+1)
}

func input(r io.Reader) (ils []initLayer, maxDepth int) {
	maxDepth = math.MinInt32
	ils = make([]initLayer, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		il := initLayer{}
		il.init(line)

		ils = append(ils, il)

		if maxDepth < il.fwdepth {
			maxDepth = il.fwdepth
		}
	}

	return
}

func buildFirewall(ils []initLayer, depth int) (fw firewall) {
	fw.init(depth)
	for _, il := range ils {
		l := &layer{}
		l.init(il.fwrange)

		fw.layers[il.fwdepth] = l
	}

	return
}

func copyLayer(in *layer) (out *layer) {
	if in == nil {
		return
	}

	out = &layer{}
	out.curr = in.curr
	out.scanDir = in.scanDir
	out.rnge = make([]bool, len(in.rnge))
	copy(out.rnge, in.rnge)

	return
}

func copyFirewall(fwin firewall) (fwout firewall) {
	fwout.init(len(fwin.layers))
	for i := 0; i < len(fwin.layers); i++ {
		fwout.layers[i] = copyLayer(fwin.layers[i])
	}

	return
}

func (fw firewall) scan() {
	for _, layer := range fw.layers {
		if layer != nil {
			layer.rnge[layer.curr] = false

			switch layer.scanDir {
			case up:
				if layer.curr < len(layer.rnge)-1 {
					layer.curr++
				} else {
					layer.scanDir = down
					layer.curr--
				}
			case down:
				if layer.curr > 0 {
					layer.curr--
				} else {
					layer.scanDir = up
					layer.curr++
				}
			}

			layer.rnge[layer.curr] = true
		}
	}
}

func (fw *firewall) advance() (res int) {
	res = -1

	l := fw.layers[fw.packet]
	if l != nil && l.rnge[0] {
		res = fw.packet * len(l.rnge)
	}

	fw.packet++

	return
}

func part1() {
	ils, maxDepth := input(os.Stdin)
	fw := buildFirewall(ils, maxDepth)

	severity := 0
	for picosecond := 0; picosecond < len(fw.layers); picosecond++ {
		if penalty := fw.advance(); penalty > 0 {
			severity += penalty
		}

		fw.scan()
	}

	log.Printf("part1 severity: %d", severity)
}

func part2() {
	ils, maxDepth := input(os.Stdin)
	fw := buildFirewall(ils, maxDepth)

	delay := 0
	for {
		caught := false
		fwCopy := copyFirewall(fw)
		for picosecond := 0; picosecond < len(fw.layers); picosecond++ {
			if penalty := fwCopy.advance(); penalty > -1 {
				caught = true
				break
			}

			fwCopy.scan()
		}

		if !caught {
			break
		}

		fw.scan()
		delay++
	}

	log.Printf("part2 minimal delay: %d", delay)
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
