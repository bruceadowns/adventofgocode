package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

func dist(x, y, z float64) float64 {
	return (math.Abs(x) + math.Abs(y) + math.Abs(z)) / 2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// cube coordinates
		var x, y, z float64

		var maxCount float64
		for _, step := range strings.Split(line, ",") {
			switch step {
			case "n":
				x++
				y--
			case "ne":
				y--
				z++
			case "se":
				x--
				z++
			case "s":
				x--
				y++
			case "sw":
				y++
				z--
			case "nw":
				x++
				z--
			default:
				log.Fatalf("invalid input: %s", step)
			}

			if maxCount < dist(x, y, z) {
				maxCount = dist(x, y, z)
			}
		}

		log.Printf("part1 count: %0.f", dist(x, y, z))
		log.Printf("part2 max count: %0.f", maxCount)
	}
}
