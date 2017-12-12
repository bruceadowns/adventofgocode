package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type knotlist struct {
	pos  int
	list []int
}

func builder(l int) knotlist {
	ret := knotlist{}
	ret.list = make([]int, l)
	for i := 0; i < len(ret.list); i++ {
		ret.list[i] = i
	}

	return ret
}

func (kl knotlist) String() string {
	sb := bytes.Buffer{}
	space := ""

	for k, v := range kl.list {
		if k == kl.pos {
			sb.WriteString(fmt.Sprintf("%s[%d]", space, v))
		} else {
			sb.WriteString(fmt.Sprintf("%s%d", space, v))
		}

		space = " "
	}

	return sb.String()
}

func (kl *knotlist) advance(i int) {
	kl.pos = (kl.pos + i) % len(kl.list)
}

func (kl knotlist) calcEnd(i int) int {
	return (kl.pos + i - 1) % len(kl.list)
}

func (kl *knotlist) knot(i int) error {
	if i < 0 || i > len(kl.list) {
		return fmt.Errorf("invalid logic i: %d len: %d", i, len(kl.list))
	}
	if i == 0 || i == 1 {
		return nil
	}

	end := kl.calcEnd(i)
	if kl.pos < end {
		slist := kl.list[kl.pos : end+1]
		for i := 0; i < len(slist)/2; i++ {
			slist[i], slist[len(slist)-i-1] = slist[len(slist)-i-1], slist[i]
		}

		// re-writing the slice, may or may not be required
		idx := 0
		for i := kl.pos; i < end+1; i++ {
			kl.list[i] = slist[idx]
			idx++
		}
	} else if kl.pos > end {
		slist := append(kl.list[kl.pos:], kl.list[:end+1]...)
		for i := 0; i < len(slist)/2; i++ {
			slist[i], slist[len(slist)-i-1] = slist[len(slist)-i-1], slist[i]
		}

		idx := 0
		for i := kl.pos; i < len(kl.list); i++ {
			kl.list[i] = slist[idx]
			idx++
		}
		for i := 0; i < end+1; i++ {
			kl.list[i] = slist[idx]
			idx++
		}
	} else {
		return fmt.Errorf("invalid logic pos: %d end: %d", kl.pos, end)
	}

	return nil
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	// array count
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}
	var l int
	n, err := fmt.Sscanf(scanner.Text(), "%d", &l)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatalf("invalid input: %d", n)
	}
	kl := builder(l)

	// array of lengths
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}
	nLens := make([]int, 0)
	for _, v := range strings.Split(scanner.Text(), ",") {
		nLen, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			log.Fatal(err)
		}
		nLens = append(nLens, nLen)
	}

	// knot and advance
	for step, nLen := range nLens {
		if err := kl.knot(nLen); err != nil {
			log.Fatal(err)
		}
		kl.advance(nLen + step)
	}

	log.Printf("part1 %d * %d = %d", kl.list[0], kl.list[1], kl.list[0]*kl.list[1])
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	// array count
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}
	var l int
	n, err := fmt.Sscanf(scanner.Text(), "%d", &l)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatalf("invalid input: %d", n)
	}
	if l != 256 {
		log.Fatalf("invalid input: %d", l)
	}
	kl := builder(l)

	// byte array of lengths
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}
	nLens := make([]int, 0)
	for _, v := range scanner.Text() {
		nLens = append(nLens, int(v))
	}
	nLens = append(nLens, []int{17, 31, 73, 47, 23}...)

	// knot and advance 64x
	skip := 0
	for i := 0; i < 64; i++ {
		for _, nLen := range nLens {
			if err := kl.knot(nLen); err != nil {
				log.Fatal(err)
			}

			kl.advance(nLen + skip)
			skip++
		}
	}

	// calc and encode dense hash
	hash := bytes.Buffer{}
	for i := 0; i < 16; i++ {
		var denseHash int
		for j := 0; j < 16; j++ {
			denseHash ^= kl.list[i*16+j]
		}

		hash.WriteString(fmt.Sprintf("%02x", denseHash))
	}
	log.Printf("part2 hash: %s", hash.String())
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
