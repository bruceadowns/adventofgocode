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

type pixel bool

type grid [][]pixel

func initGrid(size int) (res grid) {
	res = make(grid, size)
	for i := 0; i < size; i++ {
		res[i] = make([]pixel, size)
	}

	return
}

func seedGrid() (res grid) {
	// .#.
	// ..#
	// ###

	res = initGrid(3)
	res[0][1] = true
	res[1][2] = true
	res[2][0] = true
	res[2][1] = true
	res[2][2] = true

	return
}

func buildGrid(s string) (res grid) {
	// ../.#
	// ##./#../...
	// #..#/..../..../#..#

	rows := strings.Split(s, "/")
	size := len(rows)

	res = initGrid(size)
	for i := 0; i < size; i++ {
		if len(rows[i]) != size {
			log.Fatal("invalid input")
		}

		for j := 0; j < size; j++ {
			switch rows[i][j] {
			case '#':
				res[i][j] = true
			case '.':
			default:
				log.Fatal("invalid input")
			}
		}
	}

	return
}

func (g grid) String() string {
	sep := false
	sb := bytes.Buffer{}
	for _, v := range g {
		if sep {
			sb.WriteByte('/')
		}

		for _, vv := range v {
			if vv {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}

		sep = true
	}

	return sb.String()
}

func (g grid) copy() (res grid) {
	size := len(g)
	res = initGrid(size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			res[x][y] = g[x][y]
		}
	}

	return
}

func (g grid) rotate() (res grid) {
	size := len(g)
	res = g.copy()

	for i := 0; i < size-1; i++ {
		res[i][size-1] = g[0][i]
		res[size-1][size-i-1] = g[i][size-1]
		res[size-i-1][0] = g[size-1][size-i-1]
		res[0][i] = g[size-i-1][0]
	}

	return
}

func (g grid) flip() (res grid) {
	size := len(g)
	res = g.copy()

	for i := 0; i < size; i++ {
		res[i][0] = g[i][size-1]
		res[i][size-1] = g[i][0]
	}

	return
}

func (g grid) divide(subsize int) (res []grid) {
	size := len(g)
	if size%subsize != 0 {
		log.Fatal("invalid logic")
	}

	for x := 0; x < size; x += subsize {
		for y := 0; y < size; y += subsize {
			subg := initGrid(subsize)

			for subx := 0; subx < subsize; subx++ {
				for suby := 0; suby < subsize; suby++ {
					subg[subx][suby] = g[x+subx][y+suby]
				}
			}

			res = append(res, subg)
		}
	}

	return
}

func (g grid) paste(startX int, startY int, subg grid) grid {
	subsize := len(subg)

	for x := 0; x < subsize; x++ {
		for y := 0; y < subsize; y++ {
			g[x+startX][y+startY] = subg[x][y]
		}
	}

	return g
}

func (g grid) count() (res int) {
	for _, v := range g {
		for _, vv := range v {
			if vv {
				res++
			}
		}
	}

	return
}

func main() {
	iterCount := 5
	if len(os.Args) == 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("invalid input")
		}
		iterCount = n
	}

	rules := make(map[string]grid)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var src, dest string
		fmt.Sscanf(line, "%s => %s", &src, &dest)

		srcRule := buildGrid(src)
		destRule := buildGrid(dest)

		for i := 0; i < 2; i++ {
			for i := 0; i < 4; i++ {
				rules[srcRule.String()] = destRule
				srcRule = srcRule.rotate()
			}

			srcRule = srcRule.flip()
		}
	}

	image := seedGrid()

	for i := 0; i < iterCount; i++ {
		var subsize int
		size := len(image)
		if size%2 == 0 {
			subsize = 2
		} else if size%3 == 0 {
			subsize = 3
		} else {
			log.Fatal("invalid input")
		}
		subimages := image.divide(subsize)
		enhimage := initGrid(size / subsize * (subsize + 1))
		enhsize := len(enhimage)

		idx := 0
		for x := 0; x < enhsize; x += subsize + 1 {
			for y := 0; y < enhsize; y += subsize + 1 {
				repl, ok := rules[subimages[idx].String()]
				if !ok {
					log.Fatal("invalid input")
				}

				enhimage = enhimage.paste(x, y, repl)

				idx++
			}
		}

		image = enhimage
	}

	log.Printf("%d pixels stay after %d iterations", image.count(), iterCount)
}
