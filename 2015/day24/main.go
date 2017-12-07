package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type sleigh struct {
	passenger []int
	left      []int
	right     []int
}

type sleighs []sleigh

func (s sleigh) balanced() bool {
	var p, l, r int

	for _, v := range s.passenger {
		p += v
	}

	for _, v := range s.left {
		l += v
	}

	for _, v := range s.right {
		r += v
	}

	return p == l && l == r
}

func (s sleigh) qe() (res int) {
	res = 1

	for _, v := range s.passenger {
		res *= v
	}

	return
}

func (s sleighs) Len() int {
	return len(s)
}

func (s sleighs) Less(i, j int) bool {
	if len(s[i].passenger) < len(s[j].passenger) {
		return true
	} else if len(s[i].passenger) == len(s[j].passenger) && s[i].qe() < s[j].qe() {
		return true
	}

	return false
}

func (s sleighs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func die(m string) {
	log.Fatalf("error occurred [%s]", m)
}

func input() (res []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if n, err := strconv.Atoi(line); err == nil {
			res = append(res, n)
		} else {
			die(line)
		}
	}

	return
}

// adapted from https://github.com/ntns/goitertools
func combo(iterable []int, r int) [][]int {
	pool := iterable
	n := len(pool)

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]int, r)
	for i, el := range indices {
		result[i] = pool[el]
	}

	results := [][]int{result}

	for {
		i := r - 1
		for ; i >= 0 && indices[i] == i+n-r; i-- {
		}

		if i < 0 {
			return results
		}

		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}

		result := make([]int, r)
		for i = 0; i < len(indices); i++ {
			result[i] = pool[indices[i]]
		}

		results = append(results, result)
	}
}

func subtract(in1 []int, in2 ...[]int) (res []int) {
	res = []int{}

	for _, v1 := range in1 {
		found := false
	OUTER:
		for _, v2 := range in2 {
			for _, v3 := range v2 {
				if v1 == v3 {
					found = true
					break OUTER
				}
			}
		}

		if !found {
			res = append(res, v1)
		}
	}

	return
}

func main() {
	/*
	 generate three groups from input
	 where each group's sum is equal
	 where the first group has the fewest packages (legroom)
	 where the group with the fewest possibles
	  has the smallest product (quantum entanglement)
	*/

	in := input()
	fmt.Printf("in %d %v\n", len(in), in)

	// find all combinations for passenger
	// 6,474,541 combinations for 29 items
	var allSleighs sleighs
	for i := 1; i < len(in)/3; i++ {
		for _, v := range combo(in, i) {
			//fmt.Println(v)
			allSleighs = append(allSleighs, sleigh{
				passenger: v,
			})
		}
	}
	fmt.Printf("number of sleigh combinations is %d\n", len(allSleighs))

	// sort ascending via
	// 1) passenger package count
	// 2) qe
	sort.Sort(allSleighs)
	fmt.Printf("number of sorted sleigh combinations is %d\n", len(allSleighs))

	// loop through all sleighs
	// make combinations for left and right
	// find first where sums are equal
	var count int
exit:
	for _, s := range allSleighs {
		left := subtract(in, s.passenger)
		for i := 1; i < len(in)/3; i++ {
			for _, leftCombo := range combo(left, i) {
				s.left = leftCombo
				s.right = subtract(in, s.passenger, leftCombo)

				count++
				if count%10000000 == 0 {
					fmt.Printf("%d %v\n", count, s)
				}

				if s.balanced() {
					fmt.Println(s)
					break exit
				}
			}
		}
	}
}
