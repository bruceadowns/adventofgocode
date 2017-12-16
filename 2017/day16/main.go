package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	part2IterCount = 1000000000
)

type programs []rune

func (p programs) String() string {
	sb := bytes.Buffer{}
	for _, v := range p {
		sb.WriteRune(v)
	}
	return sb.String()
}

type danceMove func(programs) programs

func genSpinMove(s string) danceMove {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("error in spin: %s", err)
	}

	return func(p programs) programs {
		return append(p[len(p)-n:], p[:len(p)-n]...)
	}
}

func genExchangeMove(s string) danceMove {
	var pos1, pos2 int
	n, err := fmt.Sscanf(s, "%d/%d", &pos1, &pos2)
	if err != nil {
		log.Fatalf("error in exchange: %s", err)
	}
	if n != 2 {
		log.Fatal("invalid input in exchange")
	}

	return func(p programs) programs {
		p[pos1], p[pos2] = p[pos2], p[pos1]
		return p
	}
}

func genPartnerMove(s string) danceMove {
	if len(s) != 3 {
		log.Fatalf("invalid input in partner len %d", len(s))
	}
	if s[1] != '/' {
		log.Fatalf("invalid input in partner %s", s)
	}
	name1 := rune(s[0])
	name2 := rune(s[2])

	return func(p programs) programs {
		pos1 := -1
		for k, v := range p {
			if v == name1 {
				pos1 = k
				break
			}
		}

		pos2 := -1
		for k, v := range p {
			if v == name2 {
				pos2 = k
				break
			}
		}

		p[pos1], p[pos2] = p[pos2], p[pos1]
		return p
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}

	inputMoves := strings.Split(scanner.Text(), ",")
	if len(inputMoves) < 1 {
		log.Fatalf("invalid input")
	}

	dMoves := make([]danceMove, 0)
	for _, v := range inputMoves {
		if len(v) < 1 {
			log.Fatalf("invalid input")
		}

		switch v[0] {
		case 's':
			dMoves = append(dMoves, genSpinMove(v[1:]))
		case 'x':
			dMoves = append(dMoves, genExchangeMove(v[1:]))
		case 'p':
			dMoves = append(dMoves, genPartnerMove(v[1:]))
		default:
			log.Fatal("invalid input")
		}
	}

	cycles := make(map[string]int)

	ps := make(programs, 16)
	for i := 0; i < len(ps); i++ {
		ps[i] = 'a' + rune(i)
	}
	cycles[fmt.Sprintf("%s", ps)] = 0

	for i := 1; i < part2IterCount; i++ {
		for _, dMove := range dMoves {
			ps = dMove(ps)
		}

		if i == 1 {
			log.Printf("part1 result: %s", ps)
		}

		if _, ok := cycles[fmt.Sprintf("%s", ps)]; ok {
			break
		}
		cycles[fmt.Sprintf("%s", ps)] = i
	}

	mod := part2IterCount % len(cycles)
	result := ""
	for k, v := range cycles {
		if v == mod {
			result = k
			break
		}
	}
	log.Printf("part2 result: %s", result)
}
