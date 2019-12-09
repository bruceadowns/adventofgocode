package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
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

type machine struct {
	mem   []int
	ptr   int
	input []int
}

func build(in []int, phase int) (res machine) {
	mem := make([]int, len(in))
	copy(mem, in)

	res = machine{mem: mem}
	res.input = append(res.input, phase)

	return
}

func (m *machine) push(i int) {
	m.input = append(m.input, i)
}

func (m *machine) done() bool {
	return m.mem == nil
}

func (m *machine) intcode() (res int) {
outer:
	for {
		opcode, modes := op(m.mem[m.ptr])
		m.ptr++

		switch opcode {
		case 1:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++
			param3 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == 0 {
				value1 = m.mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = m.mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			m.mem[param3] = value1 + value2
		case 2:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++
			param3 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == 0 {
				value1 = m.mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = m.mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			m.mem[param3] = value1 * value2
		case 3:
			param1 := m.mem[m.ptr]
			m.ptr++

			input := m.input[0]
			m.input = m.input[1:]
			m.mem[param1] = input
		case 4:
			param1 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == 0 {
				value1 = m.mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			res = value1
			break outer
		case 5:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == 0 {
				value1 = m.mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = m.mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			if value1 != 0 {
				m.ptr = value2
			}
		case 6:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == 0 {
				value1 = m.mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = m.mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			if value1 == 0 {
				m.ptr = value2
			}
		case 7:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++
			param3 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == 0 {
				value1 = m.mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = m.mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			if value1 < value2 {
				m.mem[param3] = 1
			} else {
				m.mem[param3] = 0
			}
		case 8:
			param1 := m.mem[m.ptr]
			m.ptr++
			param2 := m.mem[m.ptr]
			m.ptr++
			param3 := m.mem[m.ptr]
			m.ptr++

			var value1 int
			if modes[0] == 0 {
				value1 = m.mem[param1]
			} else if modes[0] == 1 {
				value1 = param1
			}

			var value2 int
			if modes[1] == 0 {
				value2 = m.mem[param2]
			} else if modes[1] == 1 {
				value2 = param2
			}

			if value1 == value2 {
				m.mem[param3] = 1
			} else {
				m.mem[param3] = 0
			}
		case 99:
			m.mem = nil
			break outer
		default:
			log.Fatal("invalid input")
		}
	}

	return
}

// adapted from github.com/Legogris/goitertools
func permu(iterable []int) [][]int {
	r := len(iterable)

	pool := iterable
	n := len(pool)

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]int, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	results := [][]int{result}

	for n > 0 {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				result := make([]int, r)
				for k := 0; k < r; k++ {
					result[k] = pool[indices[k]]
				}

				results = append(results, result)

				break
			}
		}

		if i < 0 {
			return results
		}
	}

	return nil
}

func part1(in []int) (res int) {
	res = math.MinInt32

	c := permu([]int{0, 1, 2, 3, 4})
	for _, v := range c {
		var thrust int
		for i := 0; i < 5; i++ {
			m := build(in, v[i])
			m.push(thrust)
			thrust = m.intcode()
		}

		if thrust > res {
			res = thrust
		}
	}

	return
}

func part2(in []int) (res int) {
	res = math.MinInt32

	c := permu([]int{5, 6, 7, 8, 9})
	for _, v := range c {
		var machines [5]machine
		for i := 0; i < 5; i++ {
			machines[i] = build(in, v[i])
		}
		machines[0].push(0)

		var thrust int
		for {
			for i := 0; i < 5; i++ {
				output := machines[i].intcode()
				if !machines[i].done() {
					thrust = output
				}

				next := i + 1
				if next > 4 {
					next = 0
				}
				machines[next].push(thrust)
			}

			if machines[4].done() {
				break
			}
		}

		if thrust > res {
			res = thrust
		}
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
	log.Printf("part1: %d", part1(in))
	log.Printf("part2: %d", part2(in))
}
