package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type rule struct {
	prev [5]bool
	next bool
}

type input struct {
	initialState string
	rules        []rule
}

func in(r io.Reader) (res input) {
	scanner := bufio.NewScanner(r)

	if !scanner.Scan() {
		log.Fatal("invalid input")
	}
	if num, err := fmt.Sscanf(
		scanner.Text(),
		"initial state: %s",
		&res.initialState); err != nil {
		log.Fatal(err)
	} else if num != 1 {
		log.Fatal("invalid input")
	}
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}

	for scanner.Scan() {
		var prev, next string
		if num, err := fmt.Sscanf(
			scanner.Text(),
			"%s => %s",
			&prev, &next); err != nil {
			log.Fatal(err)
		} else if num != 2 {
			log.Fatal("invalid input")
		}
		if len(prev) != 5 || len(next) != 1 {
			log.Fatal("invalid input")
		}

		if next == "#" {
			r := rule{next: true}

			for k, v := range prev {
				switch v {
				case '#':
					r.prev[k] = true
				case '.':
				default:
					log.Fatal("invalid input")
				}
			}

			res.rules = append(res.rules, r)
		}
	}

	return
}

type pot struct {
	idx int
	on  bool
}

type tunnel []pot

func grow(in input, gen int) (res int) {
	// initialize tunnel with 4 pot padding
	var t tunnel
	for i := -4; i < 0; i++ {
		t = append(t, pot{i, false})
	}
	for k, v := range in.initialState {
		switch v {
		case '#':
			t = append(t, pot{k, true})
		case '.':
			t = append(t, pot{k, false})
		default:
			log.Fatal("invalid logic")
		}
	}
	nexti := t[len(t)-1].idx + 1
	for i := nexti; i < nexti+4; i++ {
		t = append(t, pot{i, false})
	}

	for g := 0; g < gen; g++ {
		// working tunnel is off
		wt := make(tunnel, 0)
		for _, v := range t {
			v.on = false
			wt = append(wt, v)
		}

		// match and turn on
		for i := 2; i < len(t)-2; i++ {
			for _, v := range in.rules {
				if t[i-2].on == v.prev[0] &&
					t[i-1].on == v.prev[1] &&
					t[i].on == v.prev[2] &&
					t[i+1].on == v.prev[3] &&
					t[i+2].on == v.prev[4] {
					wt[i].on = true
					break
				}
			}
		}

		// expand as necessary
		if wt[4].on {
			wt = append([]pot{
				{wt[0].idx - 1, false}},
				wt...)
		}
		if wt[len(wt)-4].on {
			wt = append(wt, pot{wt[len(wt)-1].idx + 1, false})
		}

		t = wt
	}

	// count indexes that are on
	for i := 0; i < len(t); i++ {
		if t[i].on {
			res += t[i].idx
		}
	}

	return
}

func main() {
	i := in(os.Stdin)
	log.Printf("20 generations: %d", grow(i, 20))
	log.Printf("100 generations: %d", grow(i, 100))

	// for test input
	// generation 100 is 1354
	// and each gen adds 20
	// so
	// part 2 is 1354 + ((50,000,000,000-100)*40) = 1,999,999,997,354

	// for real input
	// generation 100 is 5684
	// and each generation adds 40
	// so
	// part 2 is 5684 + ((50000000000-100)*40) = 2000000001684
}
