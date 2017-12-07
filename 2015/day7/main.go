package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	wire   = "->"
	not    = "NOT"
	and    = "AND"
	or     = "OR"
	lshift = "LSHIFT"
	rshift = "RSHIFT"
)

var board map[string]uint16

func die(msg string) {
	log.Fatalf("invalid input %s", msg)
}

func val(field string) (uint16, bool) {
	if res, err := strconv.Atoi(field); err == nil {
		return uint16(res), true
	}

	res, ok := board[field]
	return res, ok
}

func main() {
	var unresolved []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		unresolved = append(unresolved, scanner.Text())
	}

	board = make(map[string]uint16)

	for len(unresolved) > 0 {
		working := unresolved
		unresolved = nil

		for _, line := range working {
			fields := strings.Fields(line)
			if len(fields) == 3 {
				// 456 -> y
				// lx -> a

				if fields[1] != wire {
					die(line)
				}

				if v, ok := val(fields[0]); ok {
					board[fields[2]] = v
				} else {
					unresolved = append(unresolved, line)
				}
			} else if len(fields) == 4 {
				// NOT x -> h
				if fields[0] != not || fields[2] != wire {
					die(line)
				}

				if v, ok := val(fields[1]); ok {
					board[fields[3]] = ^v
				} else {
					unresolved = append(unresolved, line)
				}
			} else if len(fields) == 5 {
				// x AND y -> d
				// x OR y -> e
				// x LSHIFT 2 -> f
				// y RSHIFT 2 -> g
				if fields[1] != and && fields[1] != or && fields[1] != lshift && fields[1] != rshift {
					die(line)
				}
				if fields[3] != wire {
					die(line)
				}

				switch fields[1] {
				case and:
					// x AND y -> d
					if v, ok := val(fields[0]); ok {
						if w, ok := val(fields[2]); ok {
							board[fields[4]] = v & w
						} else {
							unresolved = append(unresolved, line)
						}
					} else {
						unresolved = append(unresolved, line)
					}
				case or:
					// x OR y -> e
					if v, ok := val(fields[0]); ok {
						if w, ok := val(fields[2]); ok {
							board[fields[4]] = v | w
						} else {
							unresolved = append(unresolved, line)
						}
					} else {
						unresolved = append(unresolved, line)
					}
				case lshift:
					// x LSHIFT 2 -> f
					shift, err := strconv.ParseUint(fields[2], 10, 8)
					if err != nil {
						die(line)
					}

					if v, ok := val(fields[0]); ok {
						board[fields[4]] = v << shift
					} else {
						unresolved = append(unresolved, line)
					}
				case rshift:
					// y RSHIFT 2 -> g
					shift, err := strconv.ParseUint(fields[2], 10, 8)
					if err != nil {
						die(line)
					}

					if v, ok := val(fields[0]); ok {
						board[fields[4]] = v >> shift
					} else {
						unresolved = append(unresolved, line)
					}
				}
			} else {
				die(line)
			}
		}
	}

	for k, v := range board {
		fmt.Println(k, v)
	}
}
