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

type abba struct {
	m [4]rune
}

func (a *abba) push(r rune) {
	if a.m[0] == 0 {
		a.m[0] = r
	} else if a.m[1] == 0 {
		a.m[1] = r
	} else if a.m[2] == 0 {
		a.m[2] = r
	} else if a.m[3] == 0 {
		a.m[3] = r
	} else {
		a.m[0] = a.m[1]
		a.m[1] = a.m[2]
		a.m[2] = a.m[3]
		a.m[3] = r
	}
}

func (a *abba) clear() {
	a.m[0] = 0
	a.m[1] = 0
	a.m[2] = 0
	a.m[3] = 0
}

func (a *abba) found() bool {
	if a.m[0] != a.m[1] && a.m[0] == a.m[3] && a.m[1] == a.m[2] {
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

		var a abba
		state := outside
		var foundOutside, foundInside bool

		for _, v := range line {
			switch v {
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				a.push(v)

				if a.found() {
					if state == outside {
						foundOutside = true
					} else if state == inside {
						foundInside = true
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

		if foundOutside {
			//log.Print("found outside")
			if foundInside {
				//log.Print("found inside")
			} else {
				count++
				//log.Print("found outside and not inside")
				//log.Print("not found inside")
			}
		} else {
			//log.Print("not found outside")
			if foundInside {
				//log.Print("found inside")
			} else {
				//log.Print("not found inside")
			}
		}
	}

	log.Printf("number of IPs that support TLS: %d", count)
}
