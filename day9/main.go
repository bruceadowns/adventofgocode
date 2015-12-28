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
	uniqueLocations := make(map[string]struct{})
	distances := make(map[string]map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// AlphaCentauri to Snowdin = 66
		line := scanner.Text()

		var from, to string
		var distance int
		n, err := fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)
		if n != 3 || err != nil {
			log.Fatalf("invalid input %s\n", line)
		}

		if _, ok := uniqueLocations[from]; !ok {
			uniqueLocations[from] = struct{}{}
		}
		if _, ok := uniqueLocations[to]; !ok {
			uniqueLocations[to] = struct{}{}
		}

		if _, ok := distances[from]; !ok {
			distances[from] = make(map[string]int)
		}
		distances[from][to] = distance

		if _, ok := distances[to]; !ok {
			distances[to] = make(map[string]int)
		}
		distances[to][from] = distance
	}

	var locations []string
	for k := range uniqueLocations {
		locations = append(locations, k)
	}

	minDistance, maxDistance := math.MaxInt32, math.MinInt32
	var minRoute, maxRoute []string
	for _, v := range permutations(locations) {
		var routeDistance int
		for i := 0; i < len(v)-1; i++ {
			routeDistance += distances[v[i]][v[i+1]]
		}
		if routeDistance < minDistance {
			minDistance = routeDistance
			minRoute = v
		}
		if routeDistance > maxDistance {
			maxDistance = routeDistance
			maxRoute = v
		}
	}

	fmt.Printf("min route distance = %d %s\n", minDistance, minRoute)
	fmt.Printf("max route distance = %d %s\n", maxDistance, maxRoute)
}
