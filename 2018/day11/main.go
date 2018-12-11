package main

import (
	"log"
	"math"
)

func initGrid(serialNumber int) (res [300][300]int) {
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			rackID := x + 10
			value := rackID * y
			value += serialNumber
			value *= rackID
			value = (value / 100) % 10
			value -= 5

			res[x-1][y-1] = value
		}
	}

	return
}

func part1(in int) (maxX, maxY int) {
	g := initGrid(in)

	max := math.MinInt32
	for y := 1; y <= 300-3; y++ {
		for x := 1; x <= 300-3; x++ {
			var work int
			for iy := y; iy < y+3; iy++ {
				for ix := x; ix < x+3; ix++ {
					work += g[ix-1][iy-1]
				}
			}

			if work > max {
				max = work
				maxX, maxY = x, y
			}
		}
	}
	// iterations
	// 297 88,209 264,627 793,881

	return
}

func part2(in int) (maxX, maxY, size int) {
	g := initGrid(in)

	max := math.MinInt32
	for s := 1; s <= 300; s++ {
		for y := 1; y <= 300-s; y++ {
			for x := 1; x <= 300-s; x++ {
				var work int
				for iy := y; iy < y+s; iy++ {
					for ix := x; ix < x+s; ix++ {
						work += g[ix-1][iy-1]
					}
				}

				if work > max {
					max = work
					maxX, maxY, size = x, y, s
				}
			}
		}
	}
	// iterations
	// 300 44,850 8,955,050 674,992,500 80,999,999,990

	return
}

func main() {
	i := 6303

	x1, y1 := part1(i)
	log.Printf("X,Y coordinate: %d,%d", x1, y1)

	x2, y2, s := part2(i)
	log.Printf("X,Y,S coordinate: %d,%d,%d", x2, y2, s)
}
