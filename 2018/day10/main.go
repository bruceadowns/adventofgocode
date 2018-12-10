package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type coord struct {
	x, y int
}

type point struct {
	position, velocity coord
}

func (p *point) crank() {
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
}

func adiff(l, r int) (res int) {
	if l < r {
		return r - l
	}
	return l - r
}

type points []*point

func (p points) plot(idx int) int {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32
	m := make(map[coord]struct{})
	for _, v := range p {
		if v.position.x < minX {
			minX = v.position.x
		}
		if v.position.x > maxX {
			maxX = v.position.x
		}
		if v.position.y < minY {
			minY = v.position.y
		}
		if v.position.y > maxY {
			maxY = v.position.y
		}

		m[v.position] = struct{}{}
	}

	if adiff(minX, maxX) < 100 && adiff(minY, maxY) < 100 {
		fmt.Printf("plot # %d\n", idx)
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if _, ok := m[coord{x, y}]; ok {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	return adiff(minX, maxX) + adiff(minY, maxY)
}

func reveal(in points) {
	idx, prevSize := 0, math.MaxInt32
	for {
		for _, v := range in {
			v.crank()
		}

		idx++
		newSize := in.plot(idx)
		if newSize > prevSize {
			break
		}
		prevSize = newSize
	}
}

func in(r io.Reader) (res points) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		p := &point{}
		num, err := fmt.Sscanf(
			scanner.Text(),
			"position=<%d, %d> velocity=<%d, %d>",
			&p.position.x, &p.position.y,
			&p.velocity.x, &p.velocity.y)
		if err != nil {
			log.Fatal(err)
		}
		if num != 4 {
			log.Fatal("invalid input")
		}

		res = append(res, p)
	}

	return
}

func main() {
	reveal(in(os.Stdin))
}
