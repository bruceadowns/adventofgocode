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
	programSize    = 16
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

func (p programs) copy() (res programs) {
	res = make(programs, programSize)
	copy(res, p)

	return
}

func (p programs) equals(q programs) bool {
	for i := 0; i < len(p); i++ {
		if p[i] != q[i] {
			return false
		}
	}

	return true
}

type danceMoveFn func(programs) programs

func genSpinMoveFn(s string) danceMoveFn {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("error in spin: %s", err)
	}

	return func(p programs) programs {
		return append(p[len(p)-n:], p[:len(p)-n]...)
	}
}

func genExchangeMoveFn(s string) danceMoveFn {
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

func genPartnerMoveFn(s string) danceMoveFn {
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

	dMoves := make([]danceMoveFn, 0)
	for _, v := range inputMoves {
		if len(v) < 1 {
			log.Fatalf("invalid input")
		}

		switch v[0] {
		case 's':
			dMoves = append(dMoves, genSpinMoveFn(v[1:]))
		case 'x':
			dMoves = append(dMoves, genExchangeMoveFn(v[1:]))
		case 'p':
			dMoves = append(dMoves, genPartnerMoveFn(v[1:]))
		default:
			log.Fatal("invalid input")
		}
	}

	pgSeed := make(programs, programSize)
	for i := 0; i < len(pgSeed); i++ {
		pgSeed[i] = 'a' + rune(i)
	}

	pgrams := pgSeed.copy()
	cycles := make([]programs, 0)
	cycles = append(cycles, pgrams)

	for i := 1; i < part2IterCount; i++ {
		for _, dMove := range dMoves {
			pgrams = dMove(pgrams)
		}

		if i == 1 {
			log.Printf("part1 result: %s", pgrams)
		}

		if pgrams.equals(pgSeed) {
			break
		}

		cycles = append(cycles, pgrams.copy())
	}

	mod := part2IterCount % len(cycles)
	log.Printf("part2 result: %s", cycles[mod])
}
