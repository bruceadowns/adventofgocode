package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type register string

type opcode string

type operation interface {
	// in is instruction index
	// out is next instruction index
	execute(in int) (out int)
}

type instruction struct {
	oc opcode
	op operation
}

type hlfOp struct {
	r register
}

func (o hlfOp) execute(i int) int {
	registers[o.r] /= 2
	return i + 1
}

type tplOp struct {
	r register
}

func (o tplOp) execute(i int) int {
	registers[o.r] *= 3
	return i + 1
}

type incOp struct {
	r register
}

func (o incOp) execute(i int) int {
	registers[o.r]++
	return i + 1
}

type jmpOp struct {
	o int
}

func (o jmpOp) execute(i int) int {
	return i + o.o
}

type jieOp struct {
	r register
	o int
}

func (o jieOp) execute(i int) int {
	if registers[o.r]%2 == 0 {
		return i + o.o
	}
	return i + 1
}

type jioOp struct {
	r register
	o int
}

func (o jioOp) execute(i int) int {
	if registers[o.r] == 1 {
		return i + o.o
	}
	return i + 1
}

func die(m string) {
	log.Fatalf("error occurred [%s]", m)
}

func readTape() (program []instruction) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			die(line)
		}
		fields := strings.Fields(line)
		if len(fields) < 1 {
			die(line)
		}

		switch fields[0] {
		case "hlf":
			if len(fields) != 2 {
				die(line)
			}

			r := register(fields[1])
			if _, ok := registers[r]; !ok {
				registers[r] = 0
			}

			program = append(program, instruction{
				oc: "hlf",
				op: hlfOp{
					r: r,
				},
			})
		case "tpl":
			if len(fields) != 2 {
				die(line)
			}

			r := register(fields[1])
			if _, ok := registers[r]; !ok {
				registers[r] = 0
			}

			program = append(program, instruction{
				oc: "tpl",
				op: tplOp{
					r: r,
				},
			})
		case "inc":
			if len(fields) != 2 {
				die(line)
			}

			r := register(fields[1])
			if _, ok := registers[r]; !ok {
				registers[r] = 0
			}

			program = append(program, instruction{
				oc: "inc",
				op: incOp{
					r: r,
				},
			})
		case "jmp":
			if len(fields) != 2 {
				die(line)
			}

			o, err := strconv.Atoi(fields[1])
			if err != nil {
				die(err.Error())
			}

			program = append(program, instruction{
				oc: "jmp",
				op: jmpOp{
					o: o,
				},
			})
		case "jie":
			if len(fields) != 3 {
				die(line)
			}

			f := fields[1]
			if f[len(f)-1] != ',' {
				die(line)
			}

			r := register(f[:len(f)-1])
			if _, ok := registers[r]; !ok {
				registers[r] = 0
			}

			o, err := strconv.Atoi(fields[2])
			if err != nil {
				die(err.Error())
			}

			program = append(program, instruction{
				oc: "jie",
				op: jieOp{
					r: r,
					o: o,
				},
			})
		case "jio":
			if len(fields) != 3 {
				die(line)
			}

			f := fields[1]
			if f[len(f)-1] != ',' {
				die(line)
			}

			r := register(f[:len(f)-1])
			if _, ok := registers[r]; !ok {
				registers[r] = 0
			}

			o, err := strconv.Atoi(fields[2])
			if err != nil {
				die(err.Error())
			}

			program = append(program, instruction{
				oc: "jio",
				op: jioOp{
					r: r,
					o: o,
				},
			})
		default:
			die("invalid opcode")
		}
	}

	return
}

func executeProgram(program []instruction) {
	for cur := 0; cur >= 0 && cur <= len(program)-1; {
		cur = program[cur].op.execute(cur)
	}
}

var registers = make(map[register]int)

func main() {
	program := readTape()

	// part 1
	{
		executeProgram(program)

		fmt.Println("part 1")
		for k, v := range registers {
			fmt.Printf("register %s is %d\n", k, v)
		}
	}

	// part 2
	{
		// reset registers
		for k := range registers {
			if k == "a" {
				registers[k] = 1
			} else {
				registers[k] = 0
			}
		}

		executeProgram(program)

		fmt.Println("part 2")
		for k, v := range registers {
			fmt.Printf("register %s is %d\n", k, v)
		}
	}
}
