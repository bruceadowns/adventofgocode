package main

import (
	"log"
	"os"
	"strconv"
)

func calcManhattanDistanceInSpiral(input int) int {
	floor := 1
	for layer := 2; ; layer++ {
		sideLen := (layer - 1) * 2
		for side := 0; side < 4; side++ {
			if floor+sideLen >= input {
				sideDiff := floor + sideLen - input
				sideMiddle := layer - 1
				offset := sideMiddle - sideDiff
				if offset < 0 {
					offset *= -1
				}
				return offset + layer - 1
			} else {
				floor += sideLen
			}
		}
	}
	return -1
}

type matrix [][]int

func resolve(m matrix, x int, y int) (sum int) {
	l := len(m[0])

	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i > l-1 {
			continue
		}

		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j > l-1 {
				continue
			}

			sum += m[i][j]
		}
	}

	m[x][y] = sum
	return
}

func expand(oldMatrix matrix) matrix {
	newSize := len(oldMatrix) + 2
	newMatrix := make(matrix, newSize)
	for i := 0; i < newSize; i++ {
		newMatrix[i] = make([]int, newSize)
	}

	for i := 1; i < newSize-1; i++ {
		for j := 1; j < newSize-1; j++ {
			newMatrix[i][j] = oldMatrix[i-1][j-1]
		}
	}
	return newMatrix
}

func firstGreaterInSpiral(input int) int {
	m := make(matrix, 1)
	m[0] = make([]int, 1)
	m[0][0] = 1

	for i := 3; ; i += 2 {
		m = expand(m)

		x := i - 1
		y := 1
		for ; y < i-1; y++ {
			if resolve(m, x, y) > input {
				return m[x][y]
			}
		}
		for ; x > 0; x-- {
			if resolve(m, x, y) > input {
				return m[x][y]
			}
		}
		for ; y > 0; y-- {
			if resolve(m, x, y) > input {
				return m[x][y]
			}
		}
		for ; x < i-1; x++ {
			if resolve(m, x, y) > input {
				return m[x][y]
			}
		}
		for ; y < 1; y++ {
			if resolve(m, x, y) > input {
				return m[x][y]
			}
		}
	}

	return -1
}

func main() {
	input := 0
	part1 := true

	if len(os.Args) == 2 || len(os.Args) == 3 {
		var err error
		input, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		if input < 2 {
			log.Fatalf("expect numeric 2 or greater [%d]", input)
		}
		log.Printf("input: %d", input)

		if len(os.Args) == 3 && os.Args[2] == "2" {
			part1 = false
		}
	} else {
		log.Fatal("expect [num] [1|2]")
	}

	if part1 {
		log.Printf("manhattan distance: %d", calcManhattanDistanceInSpiral(input))
	} else {
		log.Printf("first greater: %d", firstGreaterInSpiral(input))
	}
}
