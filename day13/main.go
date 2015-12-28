package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// https://docs.python.org/2/library/itertools.html
// https://github.com/ntns/goitertools
func permutations(iterable []string) [][]string {
	r := len(iterable)

	pool := iterable
	n := len(pool)

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]string, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	results := [][]string{result}

	for n > 0 {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				result := make([]string, r)
				for k := 0; k < r; k++ {
					result[k] = pool[indices[k]]
				}

				results = append(results, result)

				break
			}
		}

		if i < 0 {
			return results
		}
	}

	return nil
}

func main() {
	uniquePeople := make(map[string]struct{})
	allUnits := make(map[string]map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Alice would gain 54 happiness units by sitting next to Bob.
		// Alice would lose 81 happiness units by sitting next to Carol.
		line := scanner.Text()

		var from, to, factor string
		var units int
		n, err := fmt.Sscanf(
			line[:len(line)-1],
			"%s would %s %d happiness units by sitting next to %s",
			&from, &factor, &units, &to)
		if n != 4 || err != nil {
			log.Fatalf("invalid input %s", line)
		}
		switch factor {
		case "gain":
		case "lose":
			units = -units
		default:
			log.Fatalf("invalid input %s", line)
		}

		if _, ok := uniquePeople[from]; !ok {
			uniquePeople[from] = struct{}{}
		}
		if _, ok := uniquePeople[to]; !ok {
			uniquePeople[to] = struct{}{}
		}

		if _, ok := allUnits[from]; !ok {
			allUnits[from] = make(map[string]int)
		}
		allUnits[from][to] = units

		//fmt.Printf("%s %d %s\n", from, units, to)
	}

	var people []string
	for k := range uniquePeople {
		people = append(people, k)
	}
	fmt.Println(people)

	maxUnits := math.MinInt32
	var maxTable []string
	for _, v := range permutations(people) {
		var curUnits int
		for i := 0; i < len(v)-1; i++ {
			curUnits += allUnits[v[i]][v[i+1]]
			curUnits += allUnits[v[i+1]][v[i]]
		}
		// bookend
		curUnits += allUnits[v[len(v)-1]][v[0]]
		curUnits += allUnits[v[0]][v[len(v)-1]]

		if curUnits > maxUnits {
			maxUnits = curUnits
			maxTable = v
		}
	}

	fmt.Printf("most change in happiness = %d %s\n", maxUnits, maxTable)
}
