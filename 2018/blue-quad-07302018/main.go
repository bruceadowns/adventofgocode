package main

import (
	"bufio"
	"log"
	"os"
)

func combos(sl []uint8) []map[uint8]uint8 {
	if len(sl) == 0 {
		return nil
	}

	res := make([]map[uint8]uint8, 0)

	children := combos(sl[1:])
	if children == nil {
		for i := 0; i < 10; i++ {
			m := make(map[uint8]uint8)
			m[sl[0]] = uint8(i)
			res = append(res, m)
		}
	} else {
		for i := 0; i < 10; i++ {
			for _, v := range children {
				m := make(map[uint8]uint8)
				for kk, vv := range v {
					m[kk] = vv
				}
				m[sl[0]] = uint8(i)
				res = append(res, m)
			}
		}
	}

	return res
}

func main() {
	idx := 0
	var eqs [36]uint8
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		eqs[idx+0] = line[0]
		eqs[idx+1] = line[1]
		eqs[idx+2] = line[3]
		eqs[idx+3] = line[4]
		eqs[idx+4] = line[6]
		eqs[idx+5] = line[7]
		idx += 6
	}

	lMap := make(map[uint8]struct{})
	for i := 0; i < len(eqs); i++ {
		lMap[eqs[i]] = struct{}{}
	}

	lSlice := make([]uint8, 0)
	for k := range lMap {
		lSlice = append(lSlice, k)
	}

	chCombos := func() chan map[uint8]uint8 {
		res := make(chan map[uint8]uint8)

		go func() {
			for _, v := range combos(lSlice) {
				res <- v
			}

			log.Print("combos gen")
			close(res)
		}()

		return res
	}()

	idx = 0
nextPermute:
	for combo := range chCombos {
		idx++
		if idx%1000000 == 0 {
			log.Printf("idx: %d", idx)
		}

		peqs := eqs

		for _, v := range lSlice {
			for i := 0; i < len(peqs); i++ {
				if peqs[i] == v {
					peqs[i] = combo[v]
				}
			}
		}

		for i := 0; i < 6; i++ {
			diff := peqs[i+0]*10 + peqs[i+1]
			lhs := peqs[i+2]*10 + peqs[i+3]
			rhs := peqs[i+4]*10 + peqs[i+5]

			if diff != lhs-rhs {
				continue nextPermute
			}
		}

		log.Print(combo)
		break
	}
}
