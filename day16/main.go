package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type aunt struct {
	number int

	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

func num(s string) int {
	s1 := s
	s2 := s[:len(s)-1]

	n1, err := strconv.Atoi(s1)
	if err == nil {
		return n1
	}

	n2, err := strconv.Atoi(s2)
	if err == nil {
		return n2
	}

	log.Fatalf("invalid input %s", s)
	return 0
}

func (a *aunt) init(line string) {
	/*
		Sue 1: children: 1, cars: 8, vizslas: 7
		Sue 2: akitas: 10, perfumes: 10, children: 5
		Sue 3: cars: 5, pomeranians: 4, vizslas: 1
		Sue 4: goldfish: 5, children: 8, perfumes: 3
	*/

	fields := strings.Fields(line)
	if len(fields) < 4 {
		log.Fatalf("invalid input %s", line)
	}
	if fields[0] != "Sue" {
		log.Fatalf("invalid input %s", line)
	}

	a.number = num(fields[1])
	a.children, a.cats, a.samoyeds, a.pomeranians, a.akitas, a.vizslas, a.goldfish, a.trees, a.cars, a.perfumes =
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1

	for i := 2; i < len(fields); i++ {
		switch fields[i] {
		case "children:":
			a.children = num(fields[i+1])
		case "cats:":
			a.cats = num(fields[i+1])
		case "samoyeds:":
			a.samoyeds = num(fields[i+1])
		case "pomeranians:":
			a.pomeranians = num(fields[i+1])
		case "akitas:":
			a.akitas = num(fields[i+1])
		case "vizslas:":
			a.vizslas = num(fields[i+1])
		case "goldfish:":
			a.goldfish = num(fields[i+1])
		case "trees:":
			a.trees = num(fields[i+1])
		case "cars:":
			a.cars = num(fields[i+1])
		case "perfumes:":
			a.perfumes = num(fields[i+1])
		default:
			log.Fatalf("invalid input %s", line)
		}

		i++
	}
}

func rem(n int, a []*aunt) []*aunt {
	for i := 0; i < len(a); i++ {
		if a[i].number == n {
			return append(a[:i], a[i+1:]...)
		}
	}

	log.Fatalf("invalid condition %d", n)
	return nil
}

func main() {
	var idx int
	all := make([]*aunt, 500)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var a aunt
		a.init(line)
		all[idx] = &a
		idx++
	}
	if len(all) != 500 {
		log.Fatal("invalid input")
	}

	trim := make([]*aunt, len(all))
	copy(trim, all)

	for _, v := range all {
		/*
		   children: 3
		   cats: 7
		   samoyeds: 2
		   pomeranians: 3
		   akitas: 0
		   vizslas: 0
		   goldfish: 5
		   trees: 3
		   cars: 2
		   perfumes: 1
		*/

		// part 1
		// if v.cats > -1 && v.cats != 7
		// if v.trees > -1 && v.trees != 3
		// if v.pomeranians > -1 && v.pomeranians != 3
		// if v.goldfish > -1 && v.goldfish != 5

		// part 2
		// if v.cats > -1 && v.cats <= 7
		// if v.trees > -1 && v.trees <= 3
		// if v.pomeranians > -1 && v.pomeranians >= 3
		// if v.goldfish > -1 && v.goldfish >= 5

		var b bool
		if v.children > -1 && v.children != 3 {
			b = true
		} else if v.cats > -1 && v.cats <= 7 {
			b = true
		} else if v.samoyeds > -1 && v.samoyeds != 2 {
			b = true
		} else if v.pomeranians > -1 && v.pomeranians >= 3 {
			b = true
		} else if v.akitas > -1 && v.akitas != 0 {
			b = true
		} else if v.vizslas > -1 && v.vizslas != 0 {
			b = true
		} else if v.goldfish > -1 && v.goldfish >= 5 {
			b = true
		} else if v.trees > -1 && v.trees <= 3 {
			b = true
		} else if v.cars > -1 && v.cars != 2 {
			b = true
		} else if v.perfumes > -1 && v.perfumes != 1 {
			b = true
		}

		if b {
			trim = rem(v.number, trim)
		}
	}

	if len(trim) != 1 {
		log.Fatal("invalid condition")
	}
	fmt.Println(trim[0].number)
}
