package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type node int

const (
	clean node = iota
	weakened
	infected
	flagged
)

type row []node
type grid []row
type coord struct {
	x, y int
}

type cluster struct {
	g     grid
	curr  coord
	dir   direction
	count int
}

func (g grid) fault(c coord, d direction) bool {
	switch d {
	case up:
		return c.x == 0
	case left:
		return c.y == 0
	case down:
		return c.x == len(g)-1
	case right:
		return c.y == len(g)-1
	}

	return false
}

func (c *cluster) expand() {
	ng := make(grid, len(c.g)*2)
	for i := 0; i < len(ng); i++ {
		ng[i] = make(row, len(c.g)*2)
	}

	offset := len(c.g) * 2 / 4
	for x := 0; x < len(c.g); x++ {
		for y := 0; y < len(c.g[x]); y++ {
			ng[x+offset][y+offset] = c.g[x][y]
		}
	}
	c.g = ng

	c.curr.x += offset
	c.curr.y += offset

	return
}

func (c *cluster) turn1() {
	switch c.g[c.curr.x][c.curr.y] {
	case clean:
		// turn left
		switch c.dir {
		case up:
			c.dir = left
		case left:
			c.dir = down
		case down:
			c.dir = right
		case right:
			c.dir = up
		}
	case infected:
		// turn right
		switch c.dir {
		case up:
			c.dir = right
		case right:
			c.dir = down
		case down:
			c.dir = left
		case left:
			c.dir = up
		}
	default:
		log.Fatal("invalid input")
	}
}

func (c *cluster) turn2() {
	switch c.g[c.curr.x][c.curr.y] {
	case clean:
		// turn left
		switch c.dir {
		case up:
			c.dir = left
		case left:
			c.dir = down
		case down:
			c.dir = right
		case right:
			c.dir = up
		}
	case weakened:
	case infected:
		// turn right
		switch c.dir {
		case up:
			c.dir = right
		case right:
			c.dir = down
		case down:
			c.dir = left
		case left:
			c.dir = up
		}
	case flagged:
		// reverse
		switch c.dir {
		case up:
			c.dir = down
		case right:
			c.dir = left
		case down:
			c.dir = up
		case left:
			c.dir = right
		}
	default:
		log.Fatal("invalid input")
	}
}

func (c *cluster) toggle1() {
	switch c.g[c.curr.x][c.curr.y] {
	case clean:
		c.count++
		c.g[c.curr.x][c.curr.y] = infected
	case infected:
		c.g[c.curr.x][c.curr.y] = clean
	default:
		log.Fatal("invalid input")
	}
}

func (c *cluster) toggle2() {
	switch c.g[c.curr.x][c.curr.y] {
	case clean:
		c.g[c.curr.x][c.curr.y] = weakened
	case weakened:
		c.count++
		c.g[c.curr.x][c.curr.y] = infected
	case infected:
		c.g[c.curr.x][c.curr.y] = flagged
	case flagged:
		c.g[c.curr.x][c.curr.y] = clean
	default:
		log.Fatal("invalid input")
	}
}

func (c *cluster) move() {
	if c.g.fault(c.curr, c.dir) {
		c.expand()
	}

	switch c.dir {
	case up:
		c.curr.x--
	case left:
		c.curr.y--
	case down:
		c.curr.x++
	case right:
		c.curr.y++
	}
}

func input(r io.Reader) (res cluster) {
	res.g = make(grid, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		r := make(row, len(line))
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '#':
				r[i] = infected
			case '.':
				r[i] = clean
			default:
				log.Fatal("invalid input")
			}
		}

		res.g = append(res.g, r)
	}

	res.dir = up
	res.curr.x = len(res.g) / 2
	res.curr.y = len(res.g) / 2

	return
}

func part1(c cluster, burstCount int) {
	for i := 0; i < burstCount; i++ {
		c.turn1()
		c.toggle1()
		c.move()
	}

	log.Printf("part1 %d infections caused", c.count)
}

func part2(c cluster, burstCount int) {
	for i := 0; i < burstCount; i++ {
		c.turn2()
		c.toggle2()
		c.move()
	}

	log.Printf("part2 %d infections caused", c.count)
}

func main() {
	burstCount := 10000
	if len(os.Args) == 2 || len(os.Args) == 3 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("invalid input")
		}
		burstCount = n
	}

	bPart1 := true
	if len(os.Args) == 3 && os.Args[2] == "2" {
		bPart1 = false
	}

	c := input(os.Stdin)
	if len(c.g) == 0 {
		log.Fatal("invalid input")
	}

	if bPart1 {
		part1(c, burstCount)
	} else {
		part2(c, burstCount)
	}
}
