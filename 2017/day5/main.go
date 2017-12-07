package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	bPart1 := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		bPart1 = false
	}

	input := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}

	steps := 0
	curr := 0
	for {
		steps++
		next := curr + input[curr]

		if bPart1 || input[curr] < 3 {
			input[curr]++
		} else {
			input[curr]--
		}

		if next < 0 || next > len(input)-1 {
			break
		}

		curr = next
	}

	log.Printf("steps: %d", steps)
}
