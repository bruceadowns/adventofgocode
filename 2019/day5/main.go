package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func op(i int) (o int, m []int) {
	s := fmt.Sprintf("%05d", i)
	if a, err := strconv.Atoi(s[len(s)-2:]); err == nil {
		o = a
	} else {
		log.Fatal("invalid input")
	}

	if s[2] == '0' {
		m = append(m, 0)
	} else {
		m = append(m, 1)
	}

	if s[1] == '0' {
		m = append(m, 0)
	} else {
		m = append(m, 1)
	}

	if s[0] == '0' {
		m = append(m, 0)
	} else {
		m = append(m, 1)
	}

	return
}

func part1(in []int, p1 int) {
	mem := make([]int, len(in))
	copy(mem, in)

	var iPtr int
	for {
		opcode, modes := op(mem[iPtr])
		iPtr++

		switch opcode {
		case 1:
			param1 := mem[iPtr]
			iPtr++
			param2 := mem[iPtr]
			iPtr++
			param3 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			mem[param3] = value1 + value2
		case 2:
			param1 := mem[iPtr]
			iPtr++
			param2 := mem[iPtr]
			iPtr++
			param3 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			mem[param3] = value1 * value2
		case 3:
			param1 := mem[iPtr]
			iPtr++

			mem[param1] = p1
		case 4:
			param1 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			if value1 != 0 {
				log.Printf("part1: %d", value1)
			}
		case 99:
			return
		default:
			log.Fatal("invalid input")
		}
	}
}

func part2(in []int, p1 int) {
	mem := make([]int, len(in))
	copy(mem, in)

	var iPtr int
	for {
		opcode, modes := op(mem[iPtr])
		iPtr++

		switch opcode {
		case 1:
			param1 := mem[iPtr]
			iPtr++
			param2 := mem[iPtr]
			iPtr++
			param3 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			mem[param3] = value1 + value2
		case 2:
			param1 := mem[iPtr]
			iPtr++
			param2 := mem[iPtr]
			iPtr++
			param3 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			mem[param3] = value1 * value2
		case 3:
			param1 := mem[iPtr]
			iPtr++

			mem[param1] = p1
		case 4:
			param1 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			log.Printf("part2: %d", value1)
		case 5:
			param1 := mem[iPtr]
			iPtr++
			param2 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			if value1 != 0 {
				iPtr = value2
			}
		case 6:
			param1 := mem[iPtr]
			iPtr++
			param2 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			if value1 == 0 {
				iPtr = value2
			}
		case 7:
			param1 := mem[iPtr]
			iPtr++
			param2 := mem[iPtr]
			iPtr++
			param3 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			if value1 < value2 {
				mem[param3] = 1
			} else {
				mem[param3] = 0
			}
		case 8:
			param1 := mem[iPtr]
			iPtr++
			param2 := mem[iPtr]
			iPtr++
			param3 := mem[iPtr]
			iPtr++

			var value1 int
			if modes[0] == 0 {
				value1 = mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			if value1 == value2 {
				mem[param3] = 1
			} else {
				mem[param3] = 0
			}
		case 99:
			return
		default:
			log.Fatal("invalid input")
		}
	}
}

func getin(r io.Reader) (res []int) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), ",") {
		if n, err := strconv.Atoi(v); err == nil {
			res = append(res, n)
		} else {
			log.Fatal("invalid input")
		}
	}

	return
}

func main() {
	in := getin(os.Stdin)
	part1(in, 1)
	part2(in, 5)
}
