package main

import (
	"bufio"
	"log"
	"os"
)

const (
	outside = iota
	inside
)

type aba struct {
	m [3]rune
}

func (a *aba) push(r rune) {
	if a.m[0] == 0 {
		a.m[0] = r
	} else if a.m[1] == 0 {
		a.m[1] = r
	} else if a.m[2] == 0 {
		a.m[2] = r
	} else {
		a.m[0] = a.m[1]
		a.m[1] = a.m[2]
		a.m[2] = r
	}
}

func (a *aba) clear() {
	a.m[0] = 0
	a.m[1] = 0
	a.m[2] = 0
}

func (a *aba) reverse() aba {
	return aba{m: [3]rune{a.m[1], a.m[0], a.m[1]}}
}

func (a *aba) found() bool {
	if a.m[0] != a.m[1] && a.m[0] == a.m[2] {
		return true
	}
	return false
}

func main() {
	var count int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//log.Printf("%s", line)

		var a aba
		var outsideSequences, insideSequences []aba
		state := outside

		for _, v := range line {
			switch v {
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				a.push(v)

				if a.found() {
					if state == outside {
						outsideSequences = append(outsideSequences, a)
					} else if state == inside {
						insideSequences = append(insideSequences, a)
					}
				}
			case '[':
				state = inside
				a.clear()
			case ']':
				state = outside
				a.clear()
			default:
				log.Fatal("invalid logic")
			}

			//log.Printf("%v", a)
		}

	out:
		for _, o := range outsideSequences {
			//log.Printf("o: %v", o)
			r := o.reverse()
			//log.Printf("r: %v", r)
			for _, i := range insideSequences {
				//log.Printf("i: %v", i)
				if i == r {
					count++
					break out
				}
			}
		}

		//log.Print("")
	}

	log.Printf("number of IPs that support SSL: %d", count)
}
