package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type direction int

const (
	inc direction = iota
	dec
)

func (d direction) String() (res string) {
	switch d {
	case inc:
		res = "inc"
	case dec:
		res = "dec"
	}

	return
}

type conditionOperation int

const (
	eq conditionOperation = iota
	ne
	lt
	lte
	gt
	gte
)

func (c conditionOperation) String() (res string) {
	switch c {
	case eq:
		res = "=="
	case ne:
		res = "!="
	case lt:
		res = "<"
	case lte:
		res = "<="
	case gt:
		res = ">"
	case gte:
		res = ">="
	}

	return
}

type instruction struct {
	reg string
	dir direction
	amt int

	condReg string
	condOp  conditionOperation
	condAmt int
}

func (inst *instruction) String() string {
	return fmt.Sprintf("%s %s %d if %s %s %d",
		inst.reg, inst.dir, inst.amt, inst.condReg, inst.condOp, inst.condAmt)
}

func (inst *instruction) init(s string) error {
	var dir string
	var condOp string
	n, err := fmt.Sscanf(s,
		"%s %s %d if %s %s %d",
		&inst.reg, &dir, &inst.amt, &inst.condReg, &condOp, &inst.condAmt)
	if n != 6 {
		return fmt.Errorf("invalid input: %d %s", n, s)
	}
	if err != nil {
		return err
	}

	switch dir {
	case "inc":
		inst.dir = inc
	case "dec":
		inst.dir = dec
	default:
		return fmt.Errorf("invalid input: %s [%s]", dir, s)
	}

	switch condOp {
	case "==":
		inst.condOp = eq
	case "!=":
		inst.condOp = ne
	case "<":
		inst.condOp = lt
	case "<=":
		inst.condOp = lte
	case ">":
		inst.condOp = gt
	case ">=":
		inst.condOp = gte
	default:
		return fmt.Errorf("invalid input: %s [%s]", condOp, s)
	}

	return nil
}

func (inst *instruction) eval(condReg int) bool {
	switch inst.condOp {
	case eq:
		return condReg == inst.condAmt
	case ne:
		return condReg != inst.condAmt
	case lt:
		return condReg < inst.condAmt
	case lte:
		return condReg <= inst.condAmt
	case gt:
		return condReg > inst.condAmt
	case gte:
		return condReg >= inst.condAmt
	default:
		log.Fatal("invalid conditional operation")
	}

	return false
}

func main() {
	registers := make(map[string]int)
	maxDuring := math.MinInt32

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		inst := &instruction{}
		if err := inst.init(line); err != nil {
			log.Fatal(err)
		}

		if inst.eval(registers[inst.condReg]) {
			switch inst.dir {
			case inc:
				registers[inst.reg] += inst.amt
			case dec:
				registers[inst.reg] -= inst.amt
			default:
				log.Fatalf("invalid instruction direction: %v", inst.dir)
			}

			if registers[inst.reg] > maxDuring {
				maxDuring = registers[inst.reg]
			}
		}
	}

	maxAfter := math.MinInt32
	for _, v := range registers {
		if v > maxAfter {
			maxAfter = v
		}
	}
	log.Printf("part1 max register after: %d", maxAfter)
	log.Printf("part2 max register during: %d", maxDuring)
}
