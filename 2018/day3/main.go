package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type coord struct {
	x, y int
}

type claim struct {
	n    int
	l, r int
	w, h int
}

func mapify(c []claim) (res map[coord][]int) {
	res = make(map[coord][]int)
	for _, v := range c {
		for x := v.l; x < v.l+v.w; x++ {
			for y := v.r; y < v.r+v.h; y++ {
				res[coord{x, y}] = append(res[coord{x, y}], v.n)
			}
		}
	}

	return
}

func part1(c []claim) (res int) {
	m := mapify(c)
	for _, v := range m {
		if len(v) > 1 {
			res++
		}
	}

	return
}

func part2(c []claim) (res int) {
	m := mapify(c)

outer:
	for _, v := range c {
		for x := v.l; x < v.l+v.w; x++ {
			for y := v.r; y < v.r+v.h; y++ {
				if len(m[coord{x, y}]) > 1 {
					continue outer
				}
			}
		}

		return v.n
	}

	return
}

func build(s string) (res claim) {
	//#1373 @ 369,713: 20x29
	num, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d",
		&res.n, &res.l, &res.r, &res.w, &res.h)
	if err != nil {
		log.Fatal(err)
	}
	if num != 5 {
		log.Fatalf("invalid line actual %d expect 5", num)
	}

	return
}

func in(r io.Reader) (res []claim) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		res = append(res, build(scanner.Text()))
	}

	return
}

func main() {
	c := in(os.Stdin)
	log.Printf("2 or more claims: %d", part1(c))
	log.Printf("claim id that does not overlap: %d", part2(c))
}
