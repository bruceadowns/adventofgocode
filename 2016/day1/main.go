package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	north = iota
	east
	south
	west
)

type coordinate struct {
	set  bool
	x, y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c *coordinate) setIfUnset(u coordinate) {
	//log.Printf("intersection %s", u)
	if !c.set {
		c.set = true
		c.x = u.x
		c.y = u.y
	}
}

func (c *coordinate) distance() int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

func visit(current coordinate,
	firstIntersection *coordinate,
	visited map[coordinate]struct{}) {
	if _, ok := visited[current]; ok {
		(*firstIntersection).setIfUnset(current)
	} else {
		visited[current] = struct{}{}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		direction := north
		var current, firstIntersection coordinate
		visited := make(map[coordinate]struct{})

		for _, v := range strings.Split(line, ",") {
			token := strings.TrimSpace(v)
			//log.Printf("element %s", token)
			if len(token) < 2 {
				log.Fatalf("Invalid input: %s", line)
			}

			distance, err := strconv.Atoi(token[1:])
			if err != nil {
				log.Fatalf("Invalid input: %s", token)
			}

			switch token[0] {
			case 'R':
				//log.Printf("Turn right and move %d", distance)
				switch direction {
				case north:
					direction = east
				case east:
					direction = south
				case south:
					direction = west
				case west:
					direction = north
				default:
					log.Fatal("Invalid logic")
				}

			case 'L':
				//log.Printf("Turn left and move %d", distance)
				switch direction {
				case north:
					direction = west
				case east:
					direction = north
				case south:
					direction = east
				case west:
					direction = south
				default:
					log.Fatal("Invalid logic")
				}
			default:
				log.Fatalf("Invalid input: %s", token)
			}

			switch direction {
			case north:
				for i := 0; i < distance; i++ {
					current.y++
					visit(current, &firstIntersection, visited)
				}
			case east:
				for i := 0; i < distance; i++ {
					current.x++
					visit(current, &firstIntersection, visited)
				}
			case west:
				for i := 0; i < distance; i++ {
					current.x--
					visit(current, &firstIntersection, visited)
				}
			case south:
				for i := 0; i < distance; i++ {
					current.y--
					visit(current, &firstIntersection, visited)
				}
			default:
				log.Fatalf("Invalid logic %d", direction)
			}
		}

		log.Printf("End coordinates %s - total block: %d",
			current, current.distance())
		log.Printf("First intersection %s - total blocks: %d",
			firstIntersection, firstIntersection.distance())
	}
}
