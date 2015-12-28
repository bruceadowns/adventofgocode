package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nice int

	for scanner.Scan() {
		line := scanner.Text()

		var vowel int
		for _, r := range line {
			switch r {
			case 'a', 'e', 'i', 'o', 'u':
				vowel++
			}
		}
		if vowel < 3 {
			continue
		}

		var prevR rune
		var twice bool
		for _, r := range line {
			if prevR == r {
				twice = true
				break
			}
			prevR = r
		}
		if !twice {
			continue
		}

		if strings.Contains(line, "ab") {
			continue
		}
		if strings.Contains(line, "cd") {
			continue
		}
		if strings.Contains(line, "pq") {
			continue
		}
		if strings.Contains(line, "xy") {
			continue
		}

		/*
		   part 2
		   		var double bool
		   	OUTER:
		   		for i := 0; i < len(line)-2; i++ {
		   			for j := i + 2; j < len(line)-1; j++ {
		   				if line[i] == line[j] && line[i+1] == line[j+1] {
		   					double = true
		   					break OUTER
		   				}
		   			}
		   		}
		   		if !double {
		   			continue
		   		}

		   		var repeats bool
		   		for i := 0; i < len(line)-2; i++ {
		   			if line[i] == line[i+2] {
		   				repeats = true
		   				break
		   			}
		   		}
		   		if !repeats {
		   			continue
		   		}
		*/

		nice++
	}

	fmt.Printf("nice strings %d\n", nice)
}
