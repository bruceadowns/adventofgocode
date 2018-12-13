package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

type coords []coord

func (c coords) Len() int {
	return len(c)
}

func (c coords) Less(i, j int) bool {
	if c[i].y == c[j].y {
		return c[i].x < c[j].x
	}
	return c[i].y < c[j].y
}

func (c coords) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type track int
type course map[coord]track

const (
	updown track = iota
	leftright
	lrcurve
	rlcurve
	intersection
)

func runeToTrack(r rune) (res track) {
	switch r {
	case '|':
		res = updown
	case '-':
		res = leftright
	case '/':
		res = lrcurve
	case '\\':
		res = rlcurve
	case '+':
		res = intersection
	}

	return
}

type direction int

const (
	left direction = iota
	right
	up
	down
)

func runeToDirection(r rune) (res direction) {
	switch r {
	case '<':
		res = left
	case '>':
		res = right
	case '^':
		res = up
	case 'v':
		res = down
	}

	return
}

type cart struct {
	dir  direction
	turn int
}

func (c *cart) move(curr coord, co course) (res coord) {
	res = curr
	switch c.dir {
	case left:
		res.x--

		switch co[res] {
		case leftright:
		case lrcurve:
			c.dir = down
		case rlcurve:
			c.dir = up
		case intersection:
			switch c.turn % 3 {
			case 0:
				c.dir = down
			case 1:
			case 2:
				c.dir = up
			default:
				log.Fatal("invalid logic")
			}

			c.turn++
		default:
			log.Fatal("invalid logic")
		}
	case right:
		res.x++

		switch co[res] {
		case leftright:
		case lrcurve:
			c.dir = up
		case rlcurve:
			c.dir = down
		case intersection:
			switch c.turn % 3 {
			case 0:
				c.dir = up
			case 1:
			case 2:
				c.dir = down
			default:
				log.Fatal("invalid logic")
			}

			c.turn++
		default:
			log.Fatal("invalid logic")
		}
	case up:
		res.y--

		switch co[res] {
		case updown:
		case lrcurve:
			c.dir = right
		case rlcurve:
			c.dir = left
		case intersection:
			switch c.turn % 3 {
			case 0:
				c.dir = left
			case 1:
			case 2:
				c.dir = right
			default:
				log.Fatal("invalid logic")
			}

			c.turn++
		default:
			log.Printf("%v %v %d", curr, res, co[res])
			log.Fatalf("invalid logic")
		}
	case down:
		res.y++

		switch co[res] {
		case updown:
		case lrcurve:
			c.dir = left
		case rlcurve:
			c.dir = right
		case intersection:
			switch c.turn % 3 {
			case 0:
				c.dir = right
			case 1:
			case 2:
				c.dir = left
			default:
				log.Fatal("invalid logic")
			}

			c.turn++
		default:
			log.Fatal("invalid logic")
		}
	}

	return
}

type carts map[coord]cart

func in(r io.Reader) (co course, ca carts) {
	co = make(course)
	ca = make(carts)

	scanner := bufio.NewScanner(r)
	var y int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		for k, v := range scanner.Text() {
			switch v {
			case '|', '-', '/', '\\', '+':
				co[coord{k, y}] = runeToTrack(v)
			case '<', '>':
				co[coord{k, y}] = leftright
				ca[coord{k, y}] = cart{dir: runeToDirection(v)}
			case 'v', '^':
				co[coord{k, y}] = updown
				ca[coord{k, y}] = cart{dir: runeToDirection(v)}
			case ' ':
			default:
				log.Fatalf("invalid input %c", v)
			}
		}

		y++
	}

	return
}

func printMap(co course, ca carts) {
	minX, maxX := math.MaxInt32, math.MinInt32
	minY, maxY := math.MaxInt32, math.MinInt32
	for k := range co {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}

	for y := minY; y < maxY+1; y++ {
		var sb strings.Builder
		for x := minX; x < maxX+1; x++ {
			if v, ok := ca[coord{x, y}]; ok {
				switch v.dir {
				case left:
					sb.WriteRune('<')
				case right:
					sb.WriteRune('>')
				case up:
					sb.WriteRune('^')
				case down:
					sb.WriteRune('v')
				}
			} else if v, ok := co[coord{x, y}]; ok {
				switch v {
				case leftright:
					sb.WriteRune('-')
				case updown:
					sb.WriteRune('|')
				case lrcurve:
					sb.WriteRune('/')
				case rlcurve:
					sb.WriteRune('\\')
				case intersection:
					sb.WriteRune('+')
				}
			} else {
				sb.WriteRune(' ')
			}
		}

		fmt.Println(sb.String())
	}
	fmt.Println()
}

func part1(co course, ca carts) (res coord) {
	caCopy := make(carts)
	for k, v := range ca {
		caCopy[k] = v
	}

	for {
		orderedCa := make(coords, 0)
		for k := range caCopy {
			orderedCa = append(orderedCa, k)
		}
		sort.Sort(orderedCa)

		for _, v := range orderedCa {
			currCa := caCopy[v]
			newCoord := currCa.move(v, co)
			if _, ok := caCopy[newCoord]; ok {
				return newCoord
			}

			delete(caCopy, v)
			caCopy[newCoord] = currCa
		}
	}
}

func part2(co course, ca carts) (res coord) {
	caCopy := make(carts)
	for k, v := range ca {
		caCopy[k] = v
	}

	for {
		orderedCa := make(coords, 0)
		for k := range caCopy {
			orderedCa = append(orderedCa, k)
		}
		sort.Sort(orderedCa)

		for _, c := range orderedCa {
			currCa, ok := caCopy[c]
			if !ok {
				continue
			}

			newCoord := currCa.move(c, co)
			delete(caCopy, c)

			if _, ok := caCopy[newCoord]; ok {
				delete(caCopy, newCoord)
			} else {
				caCopy[newCoord] = currCa
			}
		}

		if len(caCopy) == 1 {
			sliceCa := make(coords, 0)
			for k := range caCopy {
				sliceCa = append(sliceCa, k)
			}
			return sliceCa[0]
		}
	}
}

func main() {
	co, ca := in(os.Stdin)
	printMap(co, ca)
	log.Printf("coordinate of first collision: %s", part1(co, ca))
	log.Printf("coordinate of surviving cart: %s", part2(co, ca))
}
