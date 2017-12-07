package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("input: %s", line)
		if len(line) < 1 {
			continue
		}

		sum := 0

		i := 0
		for ; i < len(line)-1; i++ {
			if line[i] < '1' || line[i] > '9' {
				log.Fatal("invalid input: ", line)
			}

			//next := i + 1
			next := (i + len(line)/2) % len(line)
			if line[i] == line[next] {
				sum += int(line[i]) - '0'
			}
		}

		//next := 0
		next := (i + len(line)/2) % len(line)
		if line[i] == line[next] {
			sum += int(line[i]) - '0'
		}

		log.Printf("sum: %d", sum)
	}
}
