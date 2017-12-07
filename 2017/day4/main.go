package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func part1() {
	numValid := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		m := make(map[string]struct{})
		valid := true
		for _, pass := range strings.Fields(line) {
			if _, ok := m[pass]; ok {
				valid = false
				break
			}
			m[pass] = struct{}{}
		}

		if valid {
			numValid++
		}
	}

	log.Printf("part1 valid count: %d", numValid)
}

func part2() {
	numValid := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		valid := true
		words := strings.Fields(line)

	nextLine:
		for i := 0; i < len(words)-1; i++ {
			for j := i + 1; j < len(words); j++ {
				if len(words[i]) == len(words[j]) {
					redact := words[i]
					for _, v := range words[j] {
						idx := strings.Index(redact, string(v))
						if idx == -1 {
							break
						}
						redact = redact[:idx] + redact[idx+1:]
					}
					if len(redact) == 0 {
						valid = false
						break nextLine
					}
				}
			}
		}

		if valid {
			numValid++
		}
	}

	log.Printf("part2 valid count: %d", numValid)
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
