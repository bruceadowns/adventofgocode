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

func op(i int) (o int, m []uint8) {
	s := fmt.Sprintf("%05d", i)

	if a, err := strconv.Atoi(s[len(s)-2:]); err == nil {
		o = a
	} else {
		log.Fatal("invalid input")
	}

	m = []uint8{s[2], s[1], s[0]}

	return
}

type machine struct {
	mem   []int
	ptr   int
	input []int
	rbase int
}

func build(in []int) machine {
	mem := make([]int, len(in))
	copy(mem, in)

	return machine{mem: mem}
}

func (m *machine) push(i int) {
	m.input = append(m.input, i)
}

func (m *machine) done() bool {
	return m.mem == nil
}

func (m *machine) read(l int) int {
	for i := len(m.mem); i <= l; i++ {
		m.mem = append(m.mem, 0)
	}

	return m.mem[l]
}

func (m *machine) write(l, v int) {
	for i := len(m.mem); i <= l; i++ {
		m.mem = append(m.mem, 0)
	}

	m.mem[l] = v
}

const (
	ADD  = 1
	MULT = 2
	SET  = 3
	OUT  = 4
	JT   = 5
	JF   = 6
	LT   = 7
	EQ   = 8
	ARB  = 9
	HALT = 99
)

func (m *machine) intcode() (res int) {
outer:
	for {
		opcode, modes := op(m.mem[m.ptr])
		m.ptr++

		switch opcode {
		case ADD:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++
			param3 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = m.read(param1)
			} else if modes[0] == '1' {
				value1 = param1
			} else if modes[0] == '2' {
				value1 = m.read(param1 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value2 int
			if modes[1] == '0' {
				value2 = m.read(param2)
			} else if modes[1] == '1' {
				value2 = param2
			} else if modes[1] == '2' {
				value2 = m.read(param2 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value3 int
			if modes[2] == '0' {
				value3 = param3
			} else if modes[2] == '1' {
				log.Fatal("invalid input")
			} else if modes[2] == '2' {
				value3 = param3 + m.rbase
			} else {
				log.Fatal("invalid input")
			}

			m.write(value3, value1+value2)
		case MULT:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++
			param3 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = m.read(param1)
			} else if modes[0] == '1' {
				value1 = param1
			} else if modes[0] == '2' {
				value1 = m.read(param1 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value2 int
			if modes[1] == '0' {
				value2 = m.read(param2)
			} else if modes[1] == '1' {
				value2 = param2
			} else if modes[1] == '2' {
				value2 = m.read(param2 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value3 int
			if modes[2] == '0' {
				value3 = param3
			} else if modes[2] == '1' {
				log.Fatal("invalid input")
			} else if modes[2] == '2' {
				value3 = param3 + m.rbase
			} else {
				log.Fatal("invalid input")
			}

			m.write(value3, value1*value2)
		case SET:
			param1 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = param1
			} else if modes[0] == '1' {
				log.Fatal("invalid input")
			} else if modes[0] == '2' {
				value1 = param1 + m.rbase
			} else {
				log.Fatal("invalid input")
			}

			input := m.input[0]
			m.input = m.input[1:]
			m.write(value1, input)
		case OUT:
			param1 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = m.read(param1)
			} else if modes[0] == '1' {
				value1 = param1
			} else if modes[0] == '2' {
				value1 = m.read(param1 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			res = value1
			break outer
		case JT:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = m.read(param1)
			} else if modes[0] == '1' {
				value1 = param1
			} else if modes[0] == '2' {
				value1 = m.read(param1 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value2 int
			if modes[1] == '0' {
				value2 = m.read(param2)
			} else if modes[1] == '1' {
				value2 = param2
			} else if modes[1] == '2' {
				value2 = m.read(param2 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			if value1 != 0 {
				m.ptr = value2
			}
		case JF:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = m.read(param1)
			} else if modes[0] == '1' {
				value1 = param1
			} else if modes[0] == '2' {
				value1 = m.read(param1 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value2 int
			if modes[1] == '0' {
				value2 = m.read(param2)
			} else if modes[1] == '1' {
				value2 = param2
			} else if modes[1] == '2' {
				value2 = m.read(param2 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			if value1 == 0 {
				m.ptr = value2
			}
		case LT:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++
			param3 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = m.read(param1)
			} else if modes[0] == '1' {
				value1 = param1
			} else if modes[0] == '2' {
				value1 = m.read(param1 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value2 int
			if modes[1] == '0' {
				value2 = m.read(param2)
			} else if modes[1] == '1' {
				value2 = param2
			} else if modes[1] == '2' {
				value2 = m.read(param2 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value3 int
			if modes[2] == '0' {
				value3 = param3
			} else if modes[2] == '1' {
				log.Fatal("invalid input")
			} else if modes[2] == '2' {
				value3 = param3 + m.rbase
			} else {
				log.Fatal("invalid input")
			}

			if value1 < value2 {
				m.write(value3, 1)
			} else {
				m.write(value3, 0)
			}
		case EQ:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++
			param3 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = m.read(param1)
			} else if modes[0] == '1' {
				value1 = param1
			} else if modes[0] == '2' {
				value1 = m.read(param1 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value2 int
			if modes[1] == '0' {
				value2 = m.read(param2)
			} else if modes[1] == '1' {
				value2 = param2
			} else if modes[1] == '2' {
				value2 = m.read(param2 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			var value3 int
			if modes[2] == '0' {
				value3 = param3
			} else if modes[2] == '1' {
				log.Fatal("invalid input")
			} else if modes[2] == '2' {
				value3 = param3 + m.rbase
			} else {
				log.Fatal("invalid input")
			}

			if value1 == value2 {
				m.write(value3, 1)
			} else {
				m.write(value3, 0)
			}
		case ARB:
			param1 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == '0' {
				value1 = m.read(param1)
			} else if modes[0] == '1' {
				value1 = param1
			} else if modes[0] == '2' {
				value1 = m.read(param1 + m.rbase)
			} else {
				log.Fatal("invalid input")
			}

			m.rbase += value1
		case HALT:
			m.mem = nil
			break outer
		default:
			log.Fatal("invalid input")
		}
	}

	return
}

func part1(in []int, input int) (res int) {
	m := build(in)
	m.push(input)

	for {
		i := m.intcode()
		log.Printf("intcode: %d", i)

		if m.done() {
			break
		}

		res = i
	}

	return
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
	log.Printf("part1: %d", part1(in, 1))
	log.Printf("part2: %d", part1(in, 2))
}
