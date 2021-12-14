package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

type fold struct {
	axis   rune
	scalar int
}

type TP struct {
	grid  map[coord]struct{}
	folds []fold
}

func newTP() TP {
	return TP{
		grid: make(map[coord]struct{}),
	}
}

func Part1(tp TP) (res int) {
	grid := make(map[coord]struct{})
	for k, v := range tp.grid {
		grid[k] = v
	}

	firstFold := tp.folds[0]
	switch firstFold.axis {
	case 'x':
		for k := range grid {
			if k.x > firstFold.scalar {
				newX := firstFold.scalar - (k.x - firstFold.scalar)
				grid[coord{newX, k.y}] = struct{}{}

				delete(grid, coord{k.x, k.y})
			}
		}
	case 'y':
		for k := range grid {
			if k.y > firstFold.scalar {
				newY := firstFold.scalar - (k.y - firstFold.scalar)
				grid[coord{k.x, newY}] = struct{}{}

				delete(grid, coord{k.x, k.y})
			}
		}
	}

	return len(grid)
}

func Part2(tp TP) (res string) {
	grid := make(map[coord]struct{})
	for k, v := range tp.grid {
		grid[k] = v
	}

	for _, v := range tp.folds {
		switch v.axis {
		case 'x':
			for k := range grid {
				if k.x > v.scalar {
					newX := v.scalar - (k.x - v.scalar)
					grid[coord{newX, k.y}] = struct{}{}

					delete(grid, coord{k.x, k.y})
				}
			}
		case 'y':
			for k := range grid {
				if k.y > v.scalar {
					newY := v.scalar - (k.y - v.scalar)
					grid[coord{k.x, newY}] = struct{}{}

					delete(grid, coord{k.x, k.y})
				}
			}
		}
	}

	var maxX int
	var maxY int
	for k := range grid {
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}

	var sb strings.Builder
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if _, exists := grid[coord{x, y}]; exists {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func In(r io.Reader) (res TP) {
	res = newTP()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		var x int
		var y int
		n, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatal("invalid input")
		}

		res.grid[coord{x, y}] = struct{}{}
	}

	for scanner.Scan() {
		line := scanner.Text()

		var axis rune
		var scalar int
		n, err := fmt.Sscanf(line, "fold along %c=%d", &axis, &scalar)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatal("invalid input")
		}

		res.folds = append(res.folds, fold{axis, scalar})
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: \n%s", Part2(i))
}
