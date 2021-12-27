package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d, %d", c.x, c.y)
}

type algoType map[int]struct{}
type imageType map[coord]struct{}

func (image imageType) bounds() (int, int, int, int) {
	minX, maxX := math.MaxInt, 0
	minY, maxY := math.MaxInt, 0
	for k := range image {
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

	return minX, maxX, minY, maxY
}

func (image imageType) String() string {
	var sb strings.Builder

	sb.WriteRune('\n')
	minX, maxX, minY, maxY := image.bounds()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, exists := image[coord{x, y}]; exists {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (image imageType) bitIndex(c coord) (res int) {
	for y := c.y - 1; y <= c.y+1; y++ {
		for x := c.x - 1; x <= c.x+1; x++ {
			res <<= 1
			if _, exists := image[coord{x, y}]; exists {
				res++
			}
		}
	}

	return res
}

type Input struct {
	algo  algoType
	image imageType
}

func solution(in Input, iterations int) int {
	currImage := make(imageType)
	for k, v := range in.image {
		currImage[k] = v
	}
	//log.Printf("%d", len(currImage))
	//log.Printf("%s", currImage)

	var turnVoidOn bool
	if _, exists := in.algo[0]; exists {
		turnVoidOn = true
	}

	for i := 0; i < iterations; i++ {
		minX, maxX, minY, maxY := currImage.bounds()

		if turnVoidOn && i%2 == 1 {
			for x := minX - 2; x <= maxX+2; x++ {
				currImage[coord{x, minY - 2}] = struct{}{}
				currImage[coord{x, minY - 1}] = struct{}{}
				currImage[coord{x, maxY + 1}] = struct{}{}
				currImage[coord{x, maxY + 2}] = struct{}{}
			}
			for y := minY - 2; y <= maxY+2; y++ {
				currImage[coord{minX - 2, y}] = struct{}{}
				currImage[coord{minX - 1, y}] = struct{}{}
				currImage[coord{maxX + 1, y}] = struct{}{}
				currImage[coord{maxX + 2, y}] = struct{}{}
			}
		}
		//log.Printf("%s", currImage)

		newImage := make(imageType)
		for y := minY - 1; y <= maxY+1; y++ {
			for x := minX - 1; x <= maxX+1; x++ {
				c := coord{x, y}
				if _, exists := in.algo[currImage.bitIndex(c)]; exists {
					newImage[c] = struct{}{}
				}
			}
		}

		currImage = newImage
		//log.Printf("%d", len(currImage))
		//log.Printf("%s", currImage)
	}

	return len(currImage)
}

func Part1(in Input) (res int) {
	return solution(in, 2)
}

func Part2(in Input) (res int) {
	return solution(in, 50)
}

func In(r io.Reader) (res Input) {
	res.algo = make(algoType)
	res.image = make(imageType)

	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		line := scanner.Text()
		if len(line) != 512 {
			log.Fatal("invalid input")
		}

		for k, v := range line {
			if v == '#' {
				res.algo[k] = struct{}{}
			}
		}
	} else {
		log.Fatal("invalid input")
	}

	if !scanner.Scan() {
		log.Fatal("invalid input")
	}

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		for k, v := range line {
			if v == '#' {
				res.image[coord{k, row}] = struct{}{}
			}
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
