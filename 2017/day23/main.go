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

type operation struct {
	optype   string
	operands []string
}

type program struct {
	curr int
	inst []operation
	mul  int
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
	case "sub":
		r[inst.operands[0]] -= r.get(inst.operands[1])
		p.curr++
	case "mul":
		r[inst.operands[0]] *= r.get(inst.operands[1])
		p.mul++
		p.curr++
	case "jnz":
		currValue := r.get(inst.operands[0])
		if currValue != 0 {
			p.curr += r.get(inst.operands[1])
		} else {
			p.curr++
		}
	default:
		err = fmt.Errorf("invalid input [%s]", inst)
	}

	if p.curr < 0 || p.curr > len(p.inst)-1 {
		err = fmt.Errorf("complete")
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

	log.Printf("part1 mul count: %d", prog.mul)
}

func part2Interpreted() {
	prog := input(os.Stdin)
	reg := make(registers, 0)
	reg["a"] = 1

	var err error
	for err == nil {
		err = prog.execute(reg)
	}

	regH, ok := reg["h"]
	if !ok {
		log.Fatal("invalid logic")
	}
	log.Printf("part2 register h: %s", regH)
}

func part2AsIs() {
	/*
		1 set b 65
		2 set c b
		3 jnz a 2
		4 jnz 1 5
		5 mul b 100
		6 sub b -100000
		7 set c b
		8 sub c -17000
		9 set f 1
		10 set d 2
		11 set e 2
		12 set g d
		13 mul g e
		14 sub g b
		15 jnz g 2
		16 set f 0
		17 sub e -1
		18 set g e
		19 sub g b
		20 jnz g -8
		21 sub d -1
		22 set g d
		23 sub g b
		24 jnz g -13
		25 jnz f 2
		26 sub h -1
		27 set g b
		28 sub g c
		29 jnz g 2
		30 jnz 1 3
		31 sub b -17
		32 jnz 1 -23
	*/
}

func part2AsIndented() {
	/*
		1 set b 65
		2 set c b

		3 jnz a 2
		4 jnz 1 5
		5 mul b 100
		6 sub b -100000
		7 set c b
		8 sub c -17000

			9 set f 1
			10 set d 2
				11 set e 2
						12 set g d
						13 mul g e
						14 sub g b
						15 jnz g 2
						16 set f 0
						17 sub e -1
						18 set g e
						19 sub g b
						20 jnz g -8
					21 sub d -1
					22 set g d
					23 sub g b
					24 jnz g -13

			25 jnz f 2
			26 sub h -1

			27 set g b
			28 sub g c

			29 jnz g 2
			30 jnz 1 3
			31 sub b -17
			32 jnz 1 -23
	*/
}

func part2Translated() {
	mul := 0
	h := 0

	//a := 0
	a := 1

	//1 set b 65
	b := 65
	//2 set c b
	c := b

	//3 jnz a 2
	if a != 0 {
		//4 jnz 1 5
		//5 mul b 100
		b *= 100
		mul++
		//6 sub b -100000
		b -= -100000
		//7 set c b
		c = b
		//8 sub c -17000
		c -= -17000
	}
	log.Printf("a: %d", a)
	log.Printf("b: %d", b)
	log.Printf("c: %d", c)

	for {
		g := 0

		//9 set f 1
		f := 1
		//10. set d 2
		d := 2

		for {
			//11 set e 2
			e := 2

			for {
				//12 set g d
				g = d
				//13 mul g e
				g *= e
				mul++
				//log.Printf("mul: %d", mul)
				//14 sub g b
				g -= b

				//15 jnz g 2
				if g == 0 {
					//16 set f 0
					f = 0
				}

				//17 sub e -1
				e -= -1
				//18 set g e
				g = e
				//19 sub g b
				g -= b

				//20 jnz g -8
				if g == 0 {
					break
				}
			}

			//21 sub d -1
			d -= -1
			//22 set g d
			g = d
			//23 sub g b
			g -= b
			//24 jnz g -13
			if g == 0 {
				break
			}
		}

		//25 jnz f 2
		if f == 0 {
			//26 sub h -1
			h -= -1
			log.Printf("h: %d", h)
		}

		//27 set g b
		g = b
		//28 sub g c
		g -= c
		log.Printf("g: %d", g)

		//29 jnz g 2
		if g == 0 {
			//30 jnz 1 3
			return
		}

		//31 sub b -17
		b -= -17
		//32 jnz 1 -23
	}
}

func part2TranslatedOp1() {
	h := 0
	b := 106500
	c := 123500

	for {
		g := 0
		f := 1
		d := 2

		for {
			e := 2

			for {
				g = d
				g *= e
				g -= b

				if g == 0 {
					f = 0
				}

				e++
				g = e
				g -= b

				if g == 0 {
					break
				}
			}

			d++
			g = d
			g -= b

			if g == 0 {
				break
			}
		}

		if f == 0 {
			h++
			log.Printf("h: %d", h)
		}

		g = b
		g -= c

		if g == 0 {
			return
		}

		b += 17
	}
}

func part2TranslatedOp2() {
	count := 0
	b := 106500
	c := 123500
	for i := b; i <= c; i += 17 {
		f := 1

		for d := 2; d < i; d++ {
			for e := 2; e < i; e++ {
				if d*e == i {
					f = 0
				}
			}
		}

		if f == 0 {
			count++
		}
	}

	log.Printf("part2 result h %d", count)
}

func part2TranslatedOp3() {
	count := 0
	for b := 106500; b < 123500+1; b += 17 {
		nonPrime := func() bool {
			for d := 2; d < b; d++ {
				for e := 2; e < b; e++ {
					if d*e == b {
						return true
					}
				}
			}
			return false
		}()

		if nonPrime {
			count++
		}
	}

	log.Printf("part2 result h: %d", count)
}

func part2TranslatedOp4() {
	count := 0
	for b := 106500; b <= 123500; b += 17 {
		nonPrime := func() bool {
			for d := 2; d < b/2; d++ {
				if b%d == 0 {
					return true
				}
			}
			return false
		}()

		if nonPrime {
			count++
		}
	}

	log.Printf("part2 result h: %d", count)
}

func main() {
	bPart1 := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		bPart1 = false
	}

	if bPart1 {
		part1()
	} else {
		//part2Interpreted()
		//part2AsIs()
		//part2AsIndented()
		//part2Translated()
		//part2TranslatedOp1()
		//part2TranslatedOp2()
		//part2TranslatedOp3()
		part2TranslatedOp4()
	}
}
