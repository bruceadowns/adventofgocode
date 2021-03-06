package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	checksum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//log.Print(line)

		min := math.MaxInt32
		max := math.MinInt32

		for _, v := range strings.Fields(line) {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("invalid input: %s", line)
			}

			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}

		checksum += max - min
	}

	log.Printf("checksum: %d", checksum)
	//checksum: 21845
}

func part2() {
	var checksum float64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//log.Print(line)

		sFields := strings.Fields(line)
		nFields := make(sort.Float64Slice, len(sFields))

		for k, v := range sFields {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				log.Fatalf("invalid input: %s", line)
			}
			nFields[k] = f
		}
		sort.Sort(sort.Reverse(nFields))

		var div float64
	out:
		for i := 0; i < len(nFields)-1; i++ {
			for j := i + 1; j < len(nFields); j++ {
				if nFields[i]/nFields[j] == math.Trunc(nFields[i]/nFields[j]) {
					div += nFields[i] / nFields[j]
					break out
				}
			}
		}
		if div == 0 {
			log.Fatalf("invalid input: %s", line)
		}

		checksum += div
	}

	log.Printf("checksum: %.0f", checksum)
	//checksum: 191
}

func main() {
	bPart1 := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		bPart1 = false
	}

	if bPart1 {
		part1()
	} else {
		part2()
	}
}
