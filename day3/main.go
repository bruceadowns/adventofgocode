package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total, valid int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		total++

		fields := strings.Fields(line)
		if len(fields) != 3 {
			log.Fatalf("invalid input: %s", line)
		}

		var a, b, c int
		var err error
		a, err = strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("invalid input: %s", line)
		}
		b, err = strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("invalid input: %s", line)
		}
		c, err = strconv.Atoi(fields[2])
		if err != nil {
			log.Fatalf("invalid input: %s", line)
		}

		if a+b > c && c+a > b && b+c > a {
			valid++
		}
	}

	log.Printf("total: %d valid: %d invalid: %d", total, valid, total-valid)
}
