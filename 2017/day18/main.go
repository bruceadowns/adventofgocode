package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type operation struct {
	optype   string
	operands []string
}

type program struct {
	curr int
	inst []operation
	last int
	sent int
}

type registers map[string]int

func (r registers) get(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}

	return r[s]
}

func (p *program) execute(r registers) (err error) {
	switch inst := p.inst[p.curr]; inst.optype {
	case "set":
		r[inst.operands[0]] = r.get(inst.operands[1])
		p.curr++
	case "add":
		r[inst.operands[0]] += r.get(inst.operands[1])
		p.curr++
	case "mul":
		r[inst.operands[0]] *= r.get(inst.operands[1])
		p.curr++
	case "mod":
		r[inst.operands[0]] %= r.get(inst.operands[1])
		p.curr++
	case "snd":
		p.last = r.get(inst.operands[0])
		p.curr++
	case "rcv":
		currValue := r.get(inst.operands[0])
		if currValue != 0 {
			err = fmt.Errorf("part1 recover sound: %d", p.last)
		}
		p.curr++
	case "jgz":
		currValue := r.get(inst.operands[0])
		if currValue > 0 {
			p.curr += r.get(inst.operands[1])
		} else {
			p.curr++
		}
	default:
		err = fmt.Errorf("invalid input")
	}

	if p.curr < 0 || p.curr > len(p.inst)-1 {
		err = fmt.Errorf("invalid logic outside")
	}

	return
}

func (p *program) executeMT(r registers, chSend chan<- int, chReceive <-chan int) (err error) {
	switch inst := p.inst[p.curr]; inst.optype {
	case "set":
		r[inst.operands[0]] = r.get(inst.operands[1])
		p.curr++
	case "add":
		r[inst.operands[0]] += r.get(inst.operands[1])
		p.curr++
	case "mul":
		r[inst.operands[0]] *= r.get(inst.operands[1])
		p.curr++
	case "mod":
		r[inst.operands[0]] %= r.get(inst.operands[1])
		p.curr++
	case "snd":
		p.last = r.get(inst.operands[0])
		chSend <- p.last
		p.sent++
		p.curr++
	case "rcv":
		select {
		case p.last = <-chReceive:
			r[inst.operands[0]] = p.last
		case <-time.After(time.Millisecond):
			err = fmt.Errorf("timeout")
		}
		p.curr++
	case "jgz":
		currValue := r.get(inst.operands[0])
		if currValue > 0 {
			p.curr += r.get(inst.operands[1])
		} else {
			p.curr++
		}
	default:
		err = fmt.Errorf("invalid input")
	}

	if p.curr < 0 || p.curr > len(p.inst)-1 {
		err = fmt.Errorf("invalid logic outside")
	}

	return
}

func input(r io.Reader) (res program) {
	res.inst = make([]operation, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 5 {
			log.Fatal("invalid input")
		}

		op := operation{
			optype:   line[0:3],
			operands: strings.Split(line[4:], " ")}

		res.inst = append(res.inst, op)
	}

	return
}

func part1() {
	prog := input(os.Stdin)
	reg := make(registers, 0)

	var err error
	for err == nil {
		err = prog.execute(reg)
	}
	log.Print(err)
}

func part2() {
	progOne := input(os.Stdin)
	progTwo := program{inst: progOne.inst}

	regOne := make(registers, 0)
	regOne["p"] = 0

	regTwo := make(registers, 0)
	regTwo["p"] = 1

	chOne := make(chan int, 64)
	chTwo := make(chan int, 64)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		var err error
		for err == nil {
			err = progOne.executeMT(regOne, chOne, chTwo)
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		var err error
		for err == nil {
			err = progTwo.executeMT(regTwo, chTwo, chOne)
		}

		wg.Done()
	}()

	wg.Wait()
	log.Printf("part2 program two sent: %d", progTwo.sent)
}

func main() {
	bPart1 := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		bPart1 = false
	}

	if bPart1 {
		part1()
	} else {
		part2()
	}
}
