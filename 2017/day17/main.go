package main

import (
	"fmt"
	"log"
	"os"
)

const (
	part1Iterations = 2017
	part2Iterations = 50000000
)

type spinlock struct {
	pos   int
	locks []int
}

func (sl *spinlock) init() {
	sl.locks = make([]int, 1)
}

func (sl *spinlock) step(steps int, value int) {
	ipos := (sl.pos + steps) % len(sl.locks)
	sl.locks = append(sl.locks[:ipos+1], append([]int{value}, sl.locks[ipos+1:]...)...)
	sl.pos = ipos + 1
}

func (sl *spinlock) stepLite(steps int, value int) {
	ipos := (sl.pos + steps) % len(sl.locks)
	sl.locks = append(sl.locks, 0)
	sl.pos = ipos + 1
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid input")
	}

	var steps int
	n, err := fmt.Sscanf(os.Args[1], "%d", &steps)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("invalid input")
	}

	// part 1
	{
		slock := spinlock{}
		slock.init()

		for i := 1; i < part1Iterations+1; i++ {
			slock.step(steps, i)
		}

		log.Printf("part1 result: %d", slock.locks[slock.pos+1])
	}

	// part 2
	{
		slock := spinlock{}
		slock.init()

		last := 0
		for i := 1; i < part2Iterations+1; i++ {
			slock.stepLite(steps, i)
			if slock.pos == 1 {
				last = i
			}
		}

		log.Printf("part result: %d", last)
	}
}
