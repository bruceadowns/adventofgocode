package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"math"
	"os"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type row []byte

func (r row) String() string {
	sb := bytes.Buffer{}
	for _, v := range r {
		sb.WriteByte(v)
	}

	return sb.String()
}

func initRow(l int) (r row) {
	r = make(row, l)
	for i := 0; i < len(r); i++ {
		r[i] = ' '
	}

	return
}

func input(r io.Reader) (diagram []row) {
	// get lines and max width
	maxLen := math.MinInt32
	lines := make([]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if maxLen < len(line) {
			maxLen = len(line)
		}

		lines = append(lines, line)
	}
	if len(lines) == 0 || maxLen == 0 {
		log.Fatal("invalid input")
	}

	// pad diagram on each of 4 sides
	diagram = make([]row, len(lines)+2)
	diagram[0] = initRow(maxLen + 2)
	diagram[len(diagram)-1] = initRow(maxLen + 2)

	for i := 0; i < len(lines); i++ {
		r := initRow(maxLen + 2)
		for j := 0; j < len(lines[i]); j++ {
			r[j+1] = lines[i][j]
		}
		diagram[i+1] = r
	}

	return
}

func main() {
	diagram := input(os.Stdin)
	if len(diagram) == 0 {
		log.Fatal("invalid input")
	}

	// find start in first row
	x := 1
	y := -1
	for k, v := range diagram[x] {
		if v == '|' {
			y = k
			break
		}
	}
	if y == -1 {
		log.Fatal("invalid input")
	}

	dir := down
	letters := make(row, 0)
	steps := 0

out:
	for {
		// move directionally
		switch dir {
		case up:
			x--
		case down:
			x++
		case left:
			y--
		case right:
			y++
		default:
			log.Fatal("invalid logic")
		}
		if x < 1 || x > len(diagram)-2 || y < 1 || y > len(diagram[x])-2 {
			log.Fatal("invalid logic")
		}

		steps++

		switch node := diagram[x][y]; node {
		case '+':
			// change direction
			if dir == down || dir == up {
				if diagram[x][y-1] != ' ' {
					dir = left
				} else if diagram[x][y+1] != ' ' {
					dir = right
				} else {
					log.Fatal("invalid input")
				}
			} else if dir == left || dir == right {
				if diagram[x-1][y] != ' ' {
					dir = up
				} else if diagram[x+1][y] != ' ' {
					dir = down
				} else {
					log.Fatal("invalid input")
				}
			} else {
				log.Fatal("invalid logic")
			}
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			// capture letter
			letters = append(letters, node)
		case ' ':
			// empty node is end
			break out
		default:
			// keep moving
		}
	}

	log.Printf("part1 letters: %s", letters)
	log.Printf("part2 total steps: %d", steps)
}
