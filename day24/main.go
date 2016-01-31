package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type sleigh struct {
	passenger []int
	left      []int
	right     []int
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
	res = make([]int, len(in1))
	copy(res, in1)

	for _, v1 := range in2 {
		for _, v2 := range v1 {
			for k, v3 := range res {
				if v2 == v3 {
					res = append(res[:k], res[k+1:]...)
					break
				}
			}
		}
	}

	return
}

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

	return (p == l) && (l == r)
}

func (s sleigh) qe() (res int) {
	res = 1

	for _, v := range s.passenger {
		res *= v
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

	// find all combinations
	var allSleighs []sleigh
	for i := 1; i < len(in)/3; i++ {
		for _, groupCombo1 := range combo(in, i) {
			group2 := subtract(in, groupCombo1)
			for j := 1; j < len(group2)/3; j++ {
				for _, groupCombo2 := range combo(group2, j) {
					group3 := subtract(in, groupCombo1, groupCombo2)
					allSleighs = append(allSleighs, sleigh{
						passenger: groupCombo1,
						left:      groupCombo2,
						right:     group3,
					})
					//fmt.Printf("groupCombo1=%v groupCombo2=%v group3=%v\n", groupCombo1, groupCombo2, group3)
				}
			}
		}
	}
	fmt.Printf("count of all combinations of sleighs %d\n", len(allSleighs))

	// find all balanced
	var allSleighsBalanced []sleigh
	for _, v := range allSleighs {
		if v.balanced() {
			allSleighsBalanced = append(allSleighsBalanced, v)
		}
	}
	fmt.Printf("balanced sleigh count %d\n", len(allSleighsBalanced))

	if len(allSleighsBalanced) < 1 {
		die("no sleighs are balanced")
	}

	// find fewest passenger packages
	fewestPassengerPackages := math.MaxInt32
	for _, v := range allSleighsBalanced {
		if len(v.passenger) < fewestPassengerPackages {
			fewestPassengerPackages = len(v.passenger)
		}
	}
	fmt.Printf("fewest passenger packages value %d\n", fewestPassengerPackages)

	// coalesce fewest passenger packages
	var allFewestPassengerPackages []sleigh
	for _, v := range allSleighsBalanced {
		if len(v.passenger) == fewestPassengerPackages {
			allFewestPassengerPackages = append(allFewestPassengerPackages, v)
		}
	}
	fmt.Printf("fewest passenger packages sleigh count %d\n", len(allFewestPassengerPackages))

	var targetSleighs []sleigh
	smallestQuantumEntanglement := math.MaxInt32
	for _, v := range allFewestPassengerPackages {
		if v.qe() < smallestQuantumEntanglement {
			smallestQuantumEntanglement = v.qe()
			targetSleighs = append(targetSleighs, v)
		}
	}
	if len(targetSleighs) != 1 {
		die("multiple target sleighs")
	}

	fmt.Printf("smallest quantum entanglement value %d\n", smallestQuantumEntanglement)
	fmt.Printf("smallest quantum entanglement sleigh %v\n", targetSleighs[0])
}
