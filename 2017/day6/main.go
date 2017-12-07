package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type bank []int

func (b bank) max() (int, int) {
	maxK := 0
	maxV := math.MinInt32
	for k, v := range b {
		if v > maxV {
			maxK = k
			maxV = v
		}
	}

	return maxK, maxV
}

func (b bank) hash() string {
	sb := bytes.Buffer{}
	for _, v := range b {
		sb.WriteString(fmt.Sprintf("%d,", v))
	}

	return sb.String()
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		b := make(bank, 0)
		for _, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
			b = append(b, n)
		}

		count := 0
		history := make(map[string]struct{})
		for {
			count++

			idx, r := b.max()
			b[idx] -= r
			for i := 0; i < r; i++ {
				idx++
				if idx > len(b)-1 {
					idx = 0
				}

				b[idx]++
			}

			h := b.hash()
			if _, ok := history[h]; ok {
				break
			}

			history[h] = struct{}{}
		}

		log.Printf("part1 count: %d", count)
	}
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		b := make(bank, 0)
		for _, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
			b = append(b, n)
		}

		count := 0
		history := make(map[string]int)
		startCount := false
		for {
			if startCount {
				count++
			}

			idx, r := b.max()
			b[idx] -= r
			for i := 0; i < r; i++ {
				idx++
				if idx > len(b)-1 {
					idx = 0
				}

				b[idx]++
			}

			h := b.hash()
			if v, ok := history[h]; ok {
				if v > 1 {
					break
				}

				startCount = true
			}

			history[h] += 1
		}

		log.Printf("part2 count: %d", count)
	}
}

func main() {
	bPart1 := false
	if len(os.Args) == 2 && os.Args[1] == "1" {
		bPart1 = true
	}

	if bPart1 {
		part1()
	} else {
		part2()
	}
}
