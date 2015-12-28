package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func coord(s string) (x int, y int) {
	sp := strings.Split(s, ",")
	if len(sp) != 2 {
		log.Fatalf("invalid coord %s", s)
	}

	var err error
	x, err = strconv.Atoi(sp[0])
	if err != nil {
		log.Fatalf("invalid coord %s", s)
	}
	y, err = strconv.Atoi(sp[1])
	if err != nil {
		log.Fatalf("invalid coord %s", s)
	}

	return
}

func main() {
	const size = 1000

	// part 1
	//var lights [size][size]bool

	// part 2
	var lights [size][size]int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// turn on 0,0 through 999,999
		// toggle 0,0 through 999,0
		// turn off 499,499 through 500,500

		fields := strings.Fields(line)
		if len(fields) != 4 && len(fields) != 5 {
			log.Fatalf("invalid input %s", line)
		}

		if fields[0] == "turn" && fields[1] == "on" && fields[3] == "through" {
			fromX, fromY := coord(fields[2])
			toX, toY := coord(fields[4])

			for x := fromX; x <= toX; x++ {
				for y := fromY; y <= toY; y++ {
					lights[x][y]++
				}
			}
		} else if fields[0] == "turn" && fields[1] == "off" && fields[3] == "through" {
			fromX, fromY := coord(fields[2])
			toX, toY := coord(fields[4])

			for x := fromX; x <= toX; x++ {
				for y := fromY; y <= toY; y++ {
					if lights[x][y] > 0 {
						lights[x][y]--
					}
				}
			}
		} else if fields[0] == "toggle" && fields[2] == "through" {
			fromX, fromY := coord(fields[1])
			toX, toY := coord(fields[3])

			for x := fromX; x <= toX; x++ {
				for y := fromY; y <= toY; y++ {
					lights[x][y] += 2
				}
			}
		} else {
			log.Fatalf("invalid input %s", line)
		}
	}

	var count int
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			count += lights[x][y]
		}
	}

	// part 1
	//fmt.Printf("%d lights are lit\n", count)

	fmt.Printf("%d light brightness\n", count)
}
