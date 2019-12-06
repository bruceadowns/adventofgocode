package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

const (
	wall = iota
	open
	goblin
	elf
)

type node int
type row []node
type area []row
type visited [][]bool
type coord struct {
	x, y int
}

func in(r io.Reader) (res area) {
	var l int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if l > 0 && l != len(line) {
			log.Fatal("invalid input")
		}
		l = len(line)

		row := make(row, 0)
		for _, v := range line {
			switch v {
			case '#':
				row = append(row, wall)
			case '.':
				row = append(row, open)
			case 'G':
				row = append(row, goblin)
			case 'E':
				row = append(row, elf)
			default:
				log.Fatal("invalid input")
			}
		}
		res = append(res, row)
	}

	return
}

/*
Targets:      In range:     Reachable:    Nearest:      Chosen:
#######       #######       #######       #######       #######
#E..G.#       #E.?G?#       #E.@G.#       #E.!G.#       #E.+G.#
#...#.#  -->  #.?.#?#  -->  #.@.#.#  -->  #.!.#.#  -->  #...#.#
#.G.#G#       #?G?#G#       #@G@#G#       #!G.#G#       #.G.#G#
#######       #######       #######       #######       #######

In range:     Nearest:      Chosen:       Distance:     Step:
#######       #######       #######       #######       #######
#.E...#       #.E...#       #.E...#       #4E212#       #..E..#
#...?.#  -->  #...!.#  -->  #...+.#  -->  #32101#  -->  #.....#
#..?G?#       #..!G.#       #...G.#       #432G2#       #...G.#
#######       #######       #######       #######       #######

#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######

#######
#E..G.#
#...#.#
#.G.#G#
#######
*/

func targets(in area, v visited, c coord, k node) (res []coord) {
	v[c.y][c.x] = true

	if !v[c.y][c.x-1] {
		switch in[c.y][c.x-1] {
		case k:
			res = append(res, coord{c.x - 1, c.y})
		case open:
			res = append(res, targets(in, v, coord{c.x - 1, c.y}, k)...)
		}
	}
	if !v[c.y-1][c.x] {
		switch in[c.y-1][c.x] {
		case k:
			res = append(res, coord{c.x, c.y - 1})
		case open:
			res = append(res, targets(in, v, coord{c.x, c.y - 1}, k)...)
		}
	}
	if !v[c.y][c.x+1] {
		switch in[c.y][c.x+1] {
		case k:
			res = append(res, coord{c.x + 1, c.y})
		case open:
			res = append(res, targets(in, v, coord{c.x + 1, c.y}, k)...)
		}
	}
	if !v[c.y+1][c.x] {
		switch in[c.y+1][c.x] {
		case k:
			res = append(res, coord{c.x, c.y + 1})
		case open:
			res = append(res, targets(in, v, coord{c.x, c.y + 1}, k)...)
		}
	}

	return
}

func buildVisited(in area) (res visited) {
	for _, v := range in {
		res = append(res, make([]bool, len(v)))
	}

	return
}

func part1(in area) (res int) {
	var round int
	for {
		var totalTargets int
		for y := 1; y < len(in)-1; y++ {
			for x := 1; x < len(in[y])-1; x++ {
				var inRange []coord

				switch in[y][x] {
				case goblin:
					inRange = targets(in, buildVisited(in), coord{x, y}, elf)
				case elf:
					inRange = targets(in, buildVisited(in), coord{x, y}, goblin)
				case wall, open:
				default:
					log.Fatal("invalid logic")
				}

				totalTargets += len(inRange)

				//2019/01/02 13:51:28 1,1 => 4,1
				//2019/01/02 13:51:28 1,1 => 2,3
				//2019/01/02 13:51:28
				//2019/01/02 13:51:28 4,1 => 1,1
				//2019/01/02 13:51:28
				//2019/01/02 13:51:28 2,3 => 1,1

				m := make(map[coord]struct{})
				for _, vv := range inRange {
					m[vv] = struct{}{}
				}

				for kk := range m {
					log.Printf("%d,%d => %d,%d", x, y, kk.x, kk.y)
				}
				log.Print()
			}
		}

		if totalTargets == 0 {
			break
		}

		round++
		break
	}

	return
}

//func part2(in area) (res int) {
//	for k, v := range in {
//		log.Printf("%d: %s", k, v)
//	}
//
//	return
//}

func main() {
	buf := `
#######
#E....#
#.#####
#.#EG.#
#.#G.E#
#.....#
#######
`
	i := in(bytes.NewBufferString(buf))
	//i := in(os.Stdin)
	log.Printf("part1: %d", part1(i))
	//log.Printf("part2: %d", part2(i))
}
