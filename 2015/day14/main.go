package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	resting = iota
	going
)

type reindeer struct {
	name        string
	speed       uint
	secondsGo   uint
	secondsRest uint

	state       uint
	secondsLeft uint
	distance    uint
	points      uint
}

func (rd *reindeer) init(line string) {
	n, err := fmt.Sscanf(
		line,
		"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
		&rd.name, &rd.speed, &rd.secondsGo, &rd.secondsRest)
	if n != 4 || err != nil {
		log.Fatalf("invalid input %s %d %v", line, n, err)
	}
	if rd.secondsGo < 1 || rd.secondsRest < 1 {
		log.Fatalf("invalid input %s", line)
	}

	rd.state = going
	rd.secondsLeft = rd.secondsGo
}

func (rd *reindeer) adv() {
	if rd.state == going {
		rd.distance += rd.speed
	}

	rd.secondsLeft--

	if rd.secondsLeft == 0 {
		if rd.state == resting {
			rd.state = going
			rd.secondsLeft = rd.secondsGo
		} else {
			rd.state = resting
			rd.secondsLeft = rd.secondsRest
		}
	}
}

func main() {
	var allReindeer []*reindeer

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Dancer can fly 7 km/s for 20 seconds, but then must rest for 119 seconds.
		line := scanner.Text()

		var rd reindeer
		rd.init(line)
		allReindeer = append(allReindeer, &rd)
	}

	//for i := 0; i < 1000; i++ {
	for i := 0; i < 2503; i++ {
		var maxDistance uint
		for _, v := range allReindeer {
			v.adv()

			if v.distance > maxDistance {
				maxDistance = v.distance
			}
		}

		// award leader(s) a point
		for _, v := range allReindeer {
			if v.distance == maxDistance {
				v.points++
			}
		}
	}

	var maxDistanceName, maxPointsName string
	var maxDistance, maxPoints uint
	for _, v := range allReindeer {
		fmt.Printf(
			"%s traveled %d km and has %d points\n",
			v.name, v.distance, v.points)

		if v.distance > maxDistance {
			maxDistanceName = v.name
			maxDistance = v.distance
		}

		if v.points > maxPoints {
			maxPointsName = v.name
			maxPoints = v.points
		}
	}

	fmt.Println()
	fmt.Printf("%s traveled the farthest at %d\n", maxDistanceName, maxDistance)
	fmt.Printf("%s has the most points at %d\n", maxPointsName, maxPoints)
}
