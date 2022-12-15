package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
)

type tree struct {
	height    int
	visible   bool
	leftView  int
	rightView int
	upView    int
	downView  int
}

func (t tree) view() (res int) {
	return t.leftView * t.rightView * t.upView * t.downView
}

func buildTree(i int32) (res tree) {
	return tree{int(i - '0'), false, 0, 0, 0, 0}
}

type Input [][]tree

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var row []tree
		for _, v := range line {
			row = append(row, buildTree(v))
		}

		res = append(res, row)
	}

	return
}

func Part1(in Input) (res int) {
	row := 0
	for col := 0; col < len(in[row]); col++ {
		in[row][col].visible = true
	}
	row = len(in) - 1
	for col := 0; col < len(in[row]); col++ {
		in[row][col].visible = true
	}
	col := 0
	for row := 0; row < len(in); row++ {
		in[row][col].visible = true
	}
	col = len(in[0]) - 1
	for row := 0; row < len(in); row++ {
		in[row][col].visible = true
	}

	for row := 0; row < len(in); row++ {
		for col := 0; col < len(in[row]); col++ {
			if !in[row][col].visible {
				// l->r
				visible := true
				for i := col - 1; i > -1; i-- {
					if in[row][i].height >= in[row][col].height {
						visible = false
						break
					}
				}
				if visible {
					in[row][col].visible = true
					continue
				}
				// r->l
				visible = true
				for i := col + 1; i < len(in[row]); i++ {
					if in[row][i].height >= in[row][col].height {
						visible = false
						break
					}
				}
				if visible {
					in[row][col].visible = true
					continue
				}
				// t->d
				visible = true
				for i := row - 1; i > -1; i-- {
					if in[i][col].height >= in[row][col].height {
						visible = false
						break
					}
				}
				if visible {
					in[row][col].visible = true
					continue
				}
				// b->u
				visible = true
				for i := row + 1; i < len(in); i++ {
					if in[i][col].height >= in[row][col].height {
						visible = false
						break
					}
				}
				if visible {
					in[row][col].visible = true
					continue
				}
			}
		}
	}

	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			if in[i][j].visible {
				res++
			}
		}
	}

	return
}

func Part2(in Input) (res int) {
	for row := 0; row < len(in); row++ {
		for col := 0; col < len(in[row]); col++ {
			// to l
			for i := col - 1; i > -1; i-- {
				in[row][col].leftView++
				if in[row][i].height >= in[row][col].height {
					break
				}

			}
			// to right
			for i := col + 1; i < len(in[row]); i++ {
				in[row][col].rightView++
				if in[row][i].height >= in[row][col].height {
					break
				}
			}
			// to up
			for i := row - 1; i > -1; i-- {
				in[row][col].upView++
				if in[i][col].height >= in[row][col].height {
					break
				}
			}
			// to down
			for i := row + 1; i < len(in); i++ {
				in[row][col].downView++
				if in[i][col].height >= in[row][col].height {
					break
				}
			}
		}
	}

	res = math.MinInt
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			totalView := in[i][j].view()
			if totalView > res {
				res = totalView
			}
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
