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

type coord struct {
	x, y int
}
type grid map[coord]node
type cluster struct {
	g     grid
	curr  coord
	dir   direction
	count int
}

func (c *cluster) turn1() {
	switch c.g[c.curr] {
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
	switch c.g[c.curr] {
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
	switch c.g[c.curr] {
	case clean:
		c.count++
		c.g[c.curr] = infected
	case infected:
		delete(c.g, c.curr)
	default:
		log.Fatal("invalid input")
	}
}

func (c *cluster) toggle2() {
	switch c.g[c.curr] {
	case clean:
		c.g[c.curr] = weakened
	case weakened:
		c.count++
		c.g[c.curr] = infected
	case infected:
		c.g[c.curr] = flagged
	case flagged:
		delete(c.g, c.curr)
	default:
		log.Fatal("invalid input")
	}
}

func (c *cluster) move() {
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
	res.g = make(grid)

	var x, y int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		for y = 0; y < len(line); y++ {
			switch line[y] {
			case '#':
				res.g[coord{x, y}] = infected
			case '.':
			default:
				log.Fatal("invalid input")
			}
		}

		x++
	}

	res.curr.x = x / 2
	res.curr.y = y / 2
	res.dir = up

	return
}

func part1(c cluster, burstCount int) {
	for i := 0; i < burstCount; i++ {
		c.turn1()
		c.toggle1()
		c.move()
	}

	log.Printf("part1 %d infections caused after %d bursts", c.count, burstCount)
}

func part2(c cluster, burstCount int) {
	for i := 0; i < burstCount; i++ {
		c.turn2()
		c.toggle2()
		c.move()
	}

	log.Printf("part2 %d infections caused after %d bursts", c.count, burstCount)
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
