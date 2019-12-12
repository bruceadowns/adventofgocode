package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	empty = iota
	asteroid
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func traj(src, dst coord) (res float64) {
	x := float64(dst.x - src.x)
	y := float64(dst.y - src.y)

	if x == 0.0 {
		if y < 0.0 {
			res = 0.0
		} else {
			res = 180.0
		}
	} else if y == 0.0 {
		if x < 0.0 {
			res = 270.0
		} else {
			res = 90.0
		}
	} else {
		res = math.Atan(math.Abs(x)/math.Abs(y)) * (180 / math.Pi)

		if x < 0.0 {
			if y < 0.0 {
				res += 180.0
			} else {
				res += 270.0
			}
		} else {
			if y < 0.0 {
				res += 0.0
			} else {
				res += 90.0
			}
		}
	}

	return
}

func part1(in map[coord]int) (res int) {
	res = math.MinInt32
	for k, v := range in {
		if v == asteroid {
			m := make(map[float64][]coord)

			for kk, vv := range in {
				if k != kk && vv == asteroid {
					t := traj(k, kk)
					m[t] = append(m[t], kk)
				}
			}

			if len(m) > res {
				res = len(m)
			}
		}
	}

	return
}

func part2(in map[coord]int) int {
	// find moniker coordinate
	// sort trajectories
	// iteratively vaporize cyclicly
	// return 200th vaporized x*100 + y
	return 0
}

func getin(r io.Reader) (res map[coord]int) {
	res = make(map[coord]int)

	scanner := bufio.NewScanner(r)
	var y int
	for scanner.Scan() {
		var x int
		for _, v := range scanner.Text() {
			switch v {
			case '#':
				res[coord{x, y}] = asteroid
			case '.':
				res[coord{x, y}] = empty
			default:
				log.Fatal("invalid input")
			}
			x++
		}
		y++
	}

	return
}

func main() {
	in := getin(os.Stdin)
	log.Printf("part1: %d", part1(in))
	log.Printf("part2: %d", part2(in))
}
