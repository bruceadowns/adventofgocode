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
	var rc [3][3]int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		index := total % 3
		total++
		//log.Printf("total: %d index: %d", total, index)

		fields := strings.Fields(line)
		if len(fields) != 3 {
			log.Fatalf("invalid input: %s", line)
		}

		var err error
		rc[0][index], err = strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("invalid input: %s", line)
		}
		rc[1][index], err = strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("invalid input: %s", line)
		}
		rc[2][index], err = strconv.Atoi(fields[2])
		if err != nil {
			log.Fatalf("invalid input: %s", line)
		}

		if index == 2 {
			if rc[0][0]+rc[0][1] > rc[0][2] && rc[0][2]+rc[0][0] > rc[0][1] && rc[0][1]+rc[0][2] > rc[0][0] {
				valid++
			}
			if rc[1][0]+rc[1][1] > rc[1][2] && rc[1][2]+rc[1][0] > rc[1][1] && rc[1][1]+rc[1][2] > rc[1][0] {
				valid++
			}
			if rc[2][0]+rc[2][1] > rc[2][2] && rc[2][2]+rc[2][0] > rc[2][1] && rc[2][1]+rc[2][2] > rc[2][0] {
				valid++
			}
		}
	}

	log.Printf("total: %d valid: %d invalid: %d", total, valid, total-valid)
}
