package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// adapted from https://github.com/ntns/goitertools
func combo(iterable []uint64, r int) [][]uint64 {
	pool := iterable
	n := len(pool)

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]uint64, r)
	for i, el := range indices {
		result[i] = pool[el]
	}

	results := [][]uint64{result}

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

		result := make([]uint64, r)
		for i = 0; i < len(indices); i++ {
			result[i] = pool[indices[i]]
		}

		results = append(results, result)
	}
}

func main() {
	var all []uint64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		n, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			log.Fatalf("invalid input %s", line)
		}
		all = append(all, n)
	}
	//fmt.Println(all)

	//const target uint64 = 25
	const target uint64 = 150

	// in.txt generates 1048575 combinations

	var total, totalMin uint
	minCombo := math.MaxUint32
	for i := 0; i < len(all); i++ {
		for _, v := range combo(all, i) {
			//fmt.Println(v)

			var liters uint64
			for _, i := range v {
				liters += i
			}

			if liters == target {
				//fmt.Printf("met target %d %v\n", target, v)
				total++

				if len(v) <= minCombo {
					//fmt.Printf("met min target %d %v\n", target, v)
					totalMin++
					minCombo = len(v)
				}
			}
		}
	}

	fmt.Printf("total combos that total %d is %d\n", target, total)
	fmt.Printf("total minimum combos [%d] that total %d is %d\n", minCombo, target, totalMin)
}
