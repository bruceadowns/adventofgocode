package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type direction int

const (
	right direction = iota
	left
)

type tape map[int]bool

type machine struct {
	t      tape
	cursor int
}

func (m *machine) init() {
	m.t = make(map[int]bool)
}

func (m machine) checksum() (res int) {
	for _, v := range m.t {
		if v {
			res++
		}
	}

	return
}

type executeFn func(*machine) byte
type stateTable map[byte]executeFn

func genExecuteFn(onWrite bool, onDirection direction, onNext byte,
	offWrite bool, offDirection direction, offNext byte) executeFn {
	return func(m *machine) (res byte) {
		if m.t[m.cursor] {
			m.t[m.cursor] = onWrite
			switch onDirection {
			case right:
				m.cursor++
			case left:
				m.cursor--
			}
			res = onNext
		} else {
			m.t[m.cursor] = offWrite
			switch offDirection {
			case right:
				m.cursor++
			case left:
				m.cursor--
			}
			res = offNext
		}

		return
	}
}

func input(scanner *bufio.Scanner) (res stateTable) {
	res = make(stateTable)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}

		var stateName byte
		if n, err := fmt.Sscanf(scanner.Text(), "In state %c:", &stateName); err != nil || n != 1 {
			log.Fatal("invalid input")
		}

		var offValue int
		if !scanner.Scan() {
			log.Fatal("invalid input")
		}
		if n, err := fmt.Sscanf(scanner.Text(), "  If the current value is %d:", &offValue); err != nil || n != 1 || offValue != 0 {
			log.Fatal("invalid input")
		}

		var offWriteI int
		if !scanner.Scan() {
			log.Fatal("invalid input")
		}
		if n, err := fmt.Sscanf(scanner.Text(), "    - Write the value %d.", &offWriteI); err != nil || n != 1 {
			log.Fatal("invalid input")
		}
		var offWrite bool
		if offWriteI == 0 {
			offWrite = false
		} else if offWriteI == 1 {
			offWrite = true
		} else {
			log.Fatal("invalid input")
		}

		var offDirS string
		if !scanner.Scan() {
			log.Fatal("invalid input")
		}
		if n, err := fmt.Sscanf(scanner.Text(), "    - Move one slot to the %s", &offDirS); err != nil || n != 1 {
			log.Fatalf("invalid input: %d [%s]", n, err)
		}
		var offDir direction
		if offDirS == "right." {
			offDir = right
		} else if offDirS == "left." {
			offDir = left
		} else {
			log.Fatal("invalid input")
		}

		var offNext byte
		if !scanner.Scan() {
			log.Fatal("invalid input")
		}
		if n, err := fmt.Sscanf(scanner.Text(), "    - Continue with state %c.", &offNext); err != nil || n != 1 {
			log.Fatalf("invalid input: %d [%s]", n, err)
		}

		var onValue int
		if !scanner.Scan() {
			log.Fatal("invalid input")
		}
		if n, err := fmt.Sscanf(scanner.Text(), "  If the current value is %d:", &onValue); err != nil || n != 1 || onValue != 1 {
			log.Fatal("invalid input")
		}

		var onWriteI int
		if !scanner.Scan() {
			log.Fatal("invalid input")
		}
		if n, err := fmt.Sscanf(scanner.Text(), "    - Write the value %d.", &onWriteI); err != nil || n != 1 {
			log.Fatal("invalid input")
		}
		var onWrite bool
		if onWriteI == 0 {
			onWrite = false
		} else if onWriteI == 1 {
			onWrite = true
		} else {
			log.Fatal("invalid input")
		}

		var onDirS string
		if !scanner.Scan() {
			log.Fatal("invalid input")
		}
		if n, err := fmt.Sscanf(scanner.Text(), "    - Move one slot to the %s", &onDirS); err != nil || n != 1 {
			log.Fatalf("invalid input: %d [%s]", n, err)
		}
		var onDir direction
		if onDirS == "right." {
			onDir = right
		} else if onDirS == "left." {
			onDir = left
		} else {
			log.Fatal("invalid input")
		}

		var onNext byte
		if !scanner.Scan() {
			log.Fatal("invalid input")
		}
		if n, err := fmt.Sscanf(scanner.Text(), "    - Continue with state %c.", &onNext); err != nil || n != 1 {
			log.Fatalf("invalid input: %d [%s]", n, err)
		}

		res[stateName] = genExecuteFn(onWrite, onDir, onNext,
			offWrite, offDir, offNext)
	}

	return
}

func main() {
	var state byte
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}
	if n, err := fmt.Sscanf(scanner.Text(), "Begin in state %c.", &state); err != nil || n != 1 {
		log.Fatal("invalid input")
	}

	var steps int
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}
	if n, err := fmt.Sscanf(scanner.Text(), "Perform a diagnostic checksum after %d steps.", &steps); err != nil || n != 1 {
		log.Fatal("invalid input")
	}

	states := input(scanner)

	m := &machine{}
	m.init()

	for i := 0; i < steps; i++ {
		state = states[state](m)
	}

	log.Printf("part1 checksum %d after %d steps", m.checksum(), steps)
	log.Printf("part2 cash in %d stars", 50)
}
