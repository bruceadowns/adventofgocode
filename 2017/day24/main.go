package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type pin int
type port [2]pin
type components []port
type bridge []port
type bridges []bridge

func (b bridge) copy() (res bridge) {
	res = make(bridge, len(b))
	copy(res, b)

	return
}

func (c components) copy() (res components) {
	res = make(components, len(c))
	copy(res, c)

	return
}

func (b bridge) strength() (res pin) {
	for _, v := range b {
		res += v[0] + v[1]
	}

	return
}

func (bs bridges) maxStrength() (res pin) {
	for _, v := range bs {
		if res < v.strength() {
			res = v.strength()
		}
	}

	return
}

func (bs bridges) maxLenStrength() (res pin) {
	var maxLen int
	for _, v := range bs {
		if maxLen < len(v) {
			maxLen = len(v)
			res = v.strength()
		} else if maxLen == len(v) {
			if res < v.strength() {
				res = v.strength()
			}
		}
	}

	return
}

func input(r io.Reader) (res components) {
	res = make(components, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var pin1, pin2 pin
		n, err := fmt.Sscanf(line, "%d/%d", &pin1, &pin2)
		if err != nil {
			log.Fatal("invalid input")
		}
		if n != 2 {
			log.Fatal("invalid input")
		}

		comp := port{}
		comp[0] = pin1
		comp[1] = pin2
		res = append(res, comp)
	}

	return
}

func build(bs bridges, b bridge, cs components, p pin) (res bridges) {
	res = append(bs, b)

	for k, v := range cs {
		if v[0] == p || v[1] == p {
			var nextPin pin
			if v[0] == p {
				nextPin = v[1]
			} else if v[1] == p {
				nextPin = v[0]
			}

			subCs := cs.copy()
			subCs = append(subCs[:k], subCs[k+1:]...)

			bNew := b.copy()
			bNew = append(bNew, v)

			res = build(res, bNew, subCs, nextPin)
		}
	}

	return
}

func main() {
	in := input(os.Stdin)

	bs := make(bridges, 0)
	for k, v := range in {
		if v[0] == 0 || v[1] == 0 {
			var nextPin pin
			if v[0] == 0 {
				nextPin = v[1]
			} else if v[1] == 0 {
				nextPin = v[0]
			}

			subIn := in.copy()
			subIn = append(subIn[:k], subIn[k+1:]...)

			bs = build(bs, bridge{v}, subIn, nextPin)
		}
	}

	log.Printf("part1 strength of the strongest bridge is: %d", bs.maxStrength())
	log.Printf("part2 strength of the longest/strongest bridge is: %d", bs.maxLenStrength())
}
