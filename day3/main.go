package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func die(msg string) {
	log.Fatalf("invalid input [%s]\n", msg)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		m := make(map[int]map[int]bool)
		var santaX, santaY, roboX, roboY int
		m[0] = make(map[int]bool)
		m[0][0] = true

		for k, r := range line {
			var x, y int

			// change following for one santa
			// if k == k {
			if k%2 == 0 {
				// santa move
				switch r {
				case '>':
					santaX++
				case '<':
					santaX--
				case '^':
					santaY++
				case 'v':
					santaY--
				default:
					log.Fatalf("invalid input [%s]", line)
				}

				x, y = santaX, santaY
			} else {
				// robo move
				switch r {
				case '>':
					roboX++
				case '<':
					roboX--
				case '^':
					roboY++
				case 'v':
					roboY--
				default:
					log.Fatalf("invalid input [%s]", line)
				}

				x, y = roboX, roboY
			}

			if m[x] == nil {
				m[x] = make(map[int]bool)
			}
			m[x][y] = true
		}

		var count int
		for _, mx := range m {
			for _, my := range mx {
				if my {
					count++
				}
			}
		}

		fmt.Printf("number of houses delivered to = %d\n", count)
	}
}
