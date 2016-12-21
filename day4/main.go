package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	one = iota
	two
	three
	four
)

type letter struct {
	ch    rune
	count int
}

type letters []letter

func (l letters) Len() int {
	return len(l)
}

func (l letters) Less(i, j int) bool {
	if l[i].count < l[j].count {
		return true
	} else if l[i].count == l[j].count {
		if l[i].ch > l[j].ch {
			return true
		}
	}

	return false
}

func (l letters) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	sectorSum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("%s", line)

		letterMap := make(map[rune]int)
		number := bytes.Buffer{}
		checkSum := bytes.Buffer{}

		state := one
		for _, r := range line {
			switch state {
			case one:
				if r >= 'a' && r <= 'z' {
					letterMap[r] = letterMap[r] + 1
				} else if r >= '0' && r <= '9' {
					state = two
					number.WriteRune(r)
				} else if r == '-' {
					//ignore
				} else {
					log.Fatal("invalid logic")
				}

			case two:
				if r >= '0' && r <= '9' {
					number.WriteRune(r)
				} else if r == '[' {
					state = three
				} else {
					log.Fatal("invalid logic")
				}

			case three:
				if r >= 'a' && r <= 'z' {
					checkSum.WriteRune(r)
				} else if r == ']' {
					state = four
				} else {
					log.Fatal("invalid logic")
				}

			default:
				log.Print("invalid logic")
			}
		}

		// sort
		var list letters
		for k, v := range letterMap {
			list = append(list, letter{k, v})
		}
		sort.Sort(sort.Reverse(list))

		checkCalc := bytes.Buffer{}
		idx := 0
		for _, v := range list {
			checkCalc.WriteRune(v.ch)
			idx++

			if idx == 5 {
				break
			}
		}

		//log.Printf("%v", letterMap)
		sectorID, err := strconv.Atoi(number.String())
		if err != nil {
			log.Fatal("invalid logic")
		}
		log.Printf("sector id: %d", sectorID)
		log.Printf("check sum: %s", checkSum.String())
		log.Printf("check sum (calc): %s", checkCalc.String())
		//log.Printf("%d", bytes.Compare(checkSum.Bytes(), checkCalc.Bytes()))
		log.Print("")

		if 0 == bytes.Compare(checkSum.Bytes(), checkCalc.Bytes()) {
			sectorSum += sectorID
		}
	}

	log.Printf("sector sum: %d", sectorSum)
}
