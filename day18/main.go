package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func print(grid [][]bool) {
	var buf bytes.Buffer
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			switch grid[i][j] {
			case true:
				buf.WriteRune('#')
			case false:
				buf.WriteRune('.')
			default:
				log.Fatal("invalid condition")
			}
		}
		buf.WriteRune('\n')
	}
	fmt.Printf("%s\n", buf.String())
}

func corners(grid [][]bool) {
	grid[1][1] = true
	grid[1][len(grid)-2] = true
	grid[len(grid)-2][1] = true
	grid[len(grid)-2][len(grid)-2] = true
}

func main() {
	// get input
	var lines []string
	var size int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		size++
	}

	// validate input
	for _, v := range lines {
		if len(v) != size {
			log.Fatal("invalid input")
		}
	}

	// make grid + outer shell
	grid := make([][]bool, size+2)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, size+2)
	}

	// populate grid
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			switch lines[i][j] {
			case '.':
			case '#':
				grid[i+1][j+1] = true
			default:
				log.Fatal("invalid input")
			}
		}
	}
	//print(grid)

	// part 2, turn on corners
	corners(grid)
	//print(grid)

	//const steps int = 4
	//const steps int = 5
	const steps int = 100

	for s := 0; s < steps; s++ {
		// copy grid arrays
		work := make([][]bool, len(grid))
		copy(work, grid)
		for i := 0; i < len(grid); i++ {
			work[i] = make([]bool, len(grid[i]))
			copy(work[i], grid[i])
		}

		for i := 1; i < len(grid)-1; i++ {
			for j := 1; j < len(grid[i])-1; j++ {
				// count neighbors
				var on uint
				for row := -1; row < 2; row++ {
					for col := -1; col < 2; col++ {
						if row == 0 && col == 0 {
							continue
						}
						if grid[i+row][j+col] {
							on++
						}
					}
				}

				// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
				// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.

				if grid[i][j] && on != 2 && on != 3 {
					work[i][j] = false
				} else if !grid[i][j] && on == 3 {
					work[i][j] = true
				}
			}
		}
		// part 2, keep corners on
		corners(work)

		grid = work
		//print(grid)
	}

	var total int
	for _, i := range grid {
		for _, j := range i {
			if j {
				total++
			}
		}
	}
	fmt.Printf("total lights on %d\n", total)
}
