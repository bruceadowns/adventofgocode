package main

import (
	"bufio"
	"log"
	"os"
)

type button struct {
	x, y int
}

func (b button) String() string {
	//log.Printf("(%d,%d)", b.x, b.y)
	if b.x == 0 && b.y == 0 {
		return "1"
	} else if b.x == 1 && b.y == 0 {
		return "2"
	} else if b.x == 2 && b.y == 0 {
		return "3"
	} else if b.x == 0 && b.y == 1 {
		return "4"
	} else if b.x == 1 && b.y == 1 {
		return "5"
	} else if b.x == 2 && b.y == 1 {
		return "6"
	} else if b.x == 0 && b.y == 2 {
		return "7"
	} else if b.x == 1 && b.y == 2 {
		return "8"
	} else if b.x == 2 && b.y == 2 {
		return "9"
	}

	return "invalid logic"
}

func (b *button) move(direction rune) {
	switch direction {
	case 'R':
		if b.x < 2 {
			b.x++
		}
	case 'L':
		if b.x > 0 {
			b.x--
		}
	case 'U':
		if b.y > 0 {
			b.y--
		}
	case 'D':
		if b.y < 2 {
			b.y++
		}
	default:
		log.Fatal("invalid logic")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	current := button{1, 1}
	var solution = make([]button, 0)

	for scanner.Scan() {
		line := scanner.Text()

		for _, direction := range line {
			current.move(direction)
		}

		solution = append(solution, current)
	}

	log.Printf("code is: %s", solution)
}
