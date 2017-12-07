package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	one = iota
	two
	three
	four
)

const (
	secretRoom = "northpole object storage"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("%s", line)

		letterMap := make(map[rune]int)
		roomName := bytes.Buffer{}
		number := bytes.Buffer{}
		checkSum := bytes.Buffer{}

		state := one
		for _, r := range line {
			switch state {
			case one:
				if r >= 'a' && r <= 'z' {
					letterMap[r] = letterMap[r] + 1
					roomName.WriteRune(r)
				} else if r >= '0' && r <= '9' {
					state = two
					number.WriteRune(r)
				} else if r == '-' {
					roomName.WriteRune(r)
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

		if 0 == bytes.Compare(checkSum.Bytes(), checkCalc.Bytes()) {
			roomNameDecrypted := bytes.Buffer{}
			for _, v := range roomName.String() {
				var n rune
				if v == '-' {
					n = ' '
				} else {
					n = v + rune(sectorID%26)
					if n > 'z' {
						n = ('a' - 1) + (n - 'z')
					}
				}
				roomNameDecrypted.WriteRune(n)
			}
			room := strings.TrimSpace(roomNameDecrypted.String())
			log.Printf("room name (decrypted): %s", room)

			if 0 == strings.Compare(secretRoom, room) {
				log.Printf("secret room: %d", sectorID)
				return
			}
		}
	}
}
