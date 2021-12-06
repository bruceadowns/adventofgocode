package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Draw []int
type Cell struct {
	value int
	on    bool
}
type Row [5]Cell
type Board [5]Row

func (board *Board) isWinner() bool {
	for i := 0; i < 5; i++ {
		if board[i][0].on && board[i][1].on && board[i][2].on && board[i][3].on && board[i][4].on {
			return true
		}

		if board[0][i].on && board[1][i].on && board[2][i].on && board[3][i].on && board[4][i].on {
			return true
		}
	}

	return false
}

func (board *Board) setDraw(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if n == board[i][j].value {
				board[i][j].on = true
				return
			}
		}
	}
}

func (board *Board) sumUnmarked() (res int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board[i][j].on {
				res += board[i][j].value
			}
		}
	}

	return
}

func Part1(draw Draw, boards []*Board) int {
	for _, v := range draw {
		for _, vv := range boards {
			vv.setDraw(v)
			if vv.isWinner() {
				return vv.sumUnmarked() * v
			}
		}
	}

	return 0
}

func Part2(draw Draw, boards []*Board) int {
	for _, v := range draw {
		if len(boards) == 1 {
			boards[0].setDraw(v)
			if boards[0].isWinner() {
				return boards[0].sumUnmarked() * v
			}
		} else {
			var boardsCurr []*Board
			for _, vv := range boards {
				vv.setDraw(v)
				if !vv.isWinner() {
					boardsCurr = append(boardsCurr, vv)
				}
			}
			boards = boardsCurr
		}
	}

	return 0
}

func makeBoard(in [5]string) (res Board) {
	for i := 0; i < 5; i++ {
		split := strings.Fields(in[i])
		if len(split) != 5 {
			log.Fatal("invalid input")
		}

		var row Row
		for j := 0; j < 5; j++ {
			if n, err := strconv.Atoi(split[j]); err == nil {
				row[j] = Cell{n, false}
			} else {
				log.Fatal("invalid input")
			}
		}

		res[i] = row
	}

	return
}

func In(r io.Reader) (draw Draw, boards []*Board) {
	scanner := bufio.NewScanner(r)

	if !scanner.Scan() {
		log.Fatal("invalid input")
	}
	line := scanner.Text()
	for _, v := range strings.Split(line, ",") {
		if n, err := strconv.Atoi(v); err == nil {
			draw = append(draw, n)
		} else {
			log.Fatal(err)
		}
	}

	var lines []string
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			continue
		}

		lines = append(lines, line)
	}
	if len(lines)%5 != 0 {
		log.Fatal("invalid input")
	}

	var partition [5]string
	for k, v := range lines {
		partition[k%5] = v
		if (k+1)%5 == 0 {
			board := makeBoard(partition)
			boards = append(boards, &board)
		}
	}

	return
}

func main() {
	draw, boards := In(os.Stdin)
	log.Printf("part1: %d", Part1(draw, boards))
	log.Printf("part2: %d", Part2(draw, boards))
}
