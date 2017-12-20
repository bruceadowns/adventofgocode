package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	part1Iterations = 1000
)

type coord struct {
	x, y, z float64
}

type particle struct {
	p         coord
	v         coord
	a         coord
	destroyed bool
}

func (part *particle) String() string {
	return fmt.Sprintf("p=<%.0f,%.0f,%.0f> v=<%.0f,%.0f,%.0f> a=<%.0f,%.0f,%.0f>",
		part.p.x, part.p.y, part.p.z,
		part.v.x, part.v.y, part.v.z,
		part.a.x, part.a.y, part.a.z)
}

func (part *particle) hash() string {
	return fmt.Sprintf("%.0f%.0f%.0f", part.p.x, part.p.y, part.p.z)
}

type buffer []*particle

func main() {
	buf := make(buffer, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var pX, pY, pZ, vX, vY, vZ, aX, aY, aZ float64
		n, err := fmt.Sscanf(line, "p=<%f,%f,%f>, v=<%f,%f,%f>, a=<%f,%f,%f>",
			&pX, &pY, &pZ, &vX, &vY, &vZ, &aX, &aY, &aZ)
		if err != nil {
			log.Fatal("invalid input")
		}
		if n != 9 {
			log.Fatal("invalid input")
		}

		buf = append(buf, &particle{
			p: coord{x: pX, y: pY, z: pZ},
			v: coord{x: vX, y: vY, z: vZ},
			a: coord{x: aX, y: aY, z: aZ}})
	}

	for i := 0; i < part1Iterations; i++ {
		for _, part := range buf {
			part.v.x += part.a.x
			part.v.y += part.a.y
			part.v.z += part.a.z
			part.p.x += part.v.x
			part.p.y += part.v.y
			part.p.z += part.v.z
		}

		uniq := make(map[string][]int)
		for k, v := range buf {
			h := v.hash()
			uniq[h] = append(uniq[h], k)
		}
		for _, vu := range uniq {
			if len(vu) > 1 {
				for _, v := range vu {
					buf[v].destroyed = true
				}
			}
		}
	}

	minValue := math.MaxFloat64
	closestParticle := -1
	survivalCount := 0
	for k, v := range buf {
		magnitude := math.Abs(v.p.x) + math.Abs(v.p.y) + math.Abs(v.p.z)
		if minValue > magnitude {
			minValue = magnitude
			closestParticle = k
		}

		if !v.destroyed {
			survivalCount++
		}
	}
	if closestParticle == -1 || survivalCount == 0 {
		log.Fatalf("invalid logic")
	}

	log.Printf("part1 %d will stay closest", closestParticle)
	log.Printf("part2 %d left after collisions", survivalCount)
}
