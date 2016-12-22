package main

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"os"
)

var m = make(map[int]map[rune]int)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//log.Printf("line: %s", line)

		for k, v := range line {
			if m[k] == nil {
				m[k] = make(map[rune]int)
			}
			m[k][v] = m[k][v] + 1
		}

		//log.Printf("m: %v", m)
	}

	message := bytes.Buffer{}
	for i := 0; i < len(m); i++ {
		//log.Printf("mk: %d - %v", i, m[i])

		var r rune
		max := math.MaxInt32
		for k, v := range m[i] {
			if v < max {
				r = k
				max = v
			}
		}

		message.WriteRune(r)
		//log.Printf("mk: %d", r)
	}

	log.Printf("message: %s", message.String())
}
