package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func part1(sl []string) (res int) {
	var twos, threes int
	for _, v := range sl {
		m := make(map[int32]int)
		for _, vv := range v {
			m[vv]++
		}

		var two, three bool
		for _, vv := range m {
			if !two && vv == 2 {
				two = true
				twos++
			}

			if !three && vv == 3 {
				three = true
				threes++
			}

			if two && three {
				// short circuit
				break
			}
		}
	}

	return twos * threes
}

func part2(sl []string) (res string) {
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			// known to be same length
			curr := sl[i]
			next := sl[j]
			l := len(curr)

			var diff, idx int
			for k := 0; k < l; k++ {
				if curr[k] != next[k] {
					diff++
					idx = k
				}

				if diff > 1 {
					break
				}
			}

			if diff == 1 {
				return fmt.Sprintf("%s%s", curr[0:idx], curr[idx+1:])
			}
		}
	}

	return
}

func in(r io.Reader) (res []string) {
	var l int
	m := make(map[string]struct{})
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		// lines expected to be same length without duplicates
		if l == 0 {
			l = len(line)
		} else if l != len(line) {
			log.Fatalf("mismatch line len: %d:%s", l, line)
		}
		if _, ok := m[line]; ok {
			log.Fatalf("dup line: %s", line)
		}
		m[line] = struct{}{}

		res = append(res, line)
	}

	return
}

func main() {
	sl := in(os.Stdin)
	log.Printf("checksum: %d", part1(sl))
	log.Printf("common chars: %s", part2(sl))
	part2(sl)
}
