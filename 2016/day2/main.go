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
	if b.x == 2 && b.y == 0 {
		return "1"
	} else if b.x == 1 && b.y == 1 {
		return "2"
	} else if b.x == 2 && b.y == 1 {
		return "3"
	} else if b.x == 3 && b.y == 1 {
		return "4"
	} else if b.x == 0 && b.y == 2 {
		return "5"
	} else if b.x == 1 && b.y == 2 {
		return "6"
	} else if b.x == 2 && b.y == 2 {
		return "7"
	} else if b.x == 3 && b.y == 2 {
		return "8"
	} else if b.x == 4 && b.y == 2 {
		return "9"
	} else if b.x == 1 && b.y == 3 {
		return "A"
	} else if b.x == 2 && b.y == 3 {
		return "B"
	} else if b.x == 3 && b.y == 3 {
		return "C"
	} else if b.x == 2 && b.y == 4 {
		return "D"
	}

	return "invalid logic"
}

func (b *button) move(direction rune) {
	switch direction {
	case 'R':
		if (b.y == 1 || b.y == 3) && b.x < 3 {
			b.x++
		} else if b.y == 2 && b.x < 4 {
			b.x++
		}
	case 'L':
		if (b.y == 1 || b.y == 3) && b.x > 1 {
			b.x--
		} else if b.y == 2 && b.x > 0 {
			b.x--
		}
	case 'D':
		if (b.x == 1 || b.x == 3) && b.y < 3 {
			b.y++
		} else if b.x == 2 && b.y < 4 {
			b.y++
		}
	case 'U':
		if (b.x == 1 || b.x == 3) && b.y > 1 {
			b.y--
		} else if b.x == 2 && b.y > 0 {
			b.y--
		}
	default:
		log.Fatal("invalid logic")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	current := button{0, 2}
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
