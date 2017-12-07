package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var flight int
		var basementPosition int

		for k, v := range line {
			switch v {
			case '(':
				flight++
			case ')':
				flight--
			default:
				log.Fatalf("invalid input [%s]", string(v))
			}

			if basementPosition < 1 && flight < 0 {
				basementPosition = k + 1
			}
		}

		fmt.Printf("current floor is %d\n", flight)
		fmt.Printf("entered basement at position %d\n", basementPosition)
	}
}
