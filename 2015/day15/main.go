package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	totalScore int = 100
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

var allIngredients []*ingredient

func (ing *ingredient) init(line string) {
	n, err := fmt.Sscanf(
		line,
		"%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
		&ing.name, &ing.capacity, &ing.durability, &ing.flavor, &ing.texture, &ing.calories)
	if n != 6 || err != nil {
		log.Fatalf("invalid input %s %d %v", line, n, err)
	}
	if ing.name[len(ing.name)-1] != ':' {
		log.Fatalf("invalid input %s", line)
	}
	ing.name = ing.name[:len(ing.name)-1]
}

func permu(total, num int, in []int) (res [][]int) {
	rem := total
	for i := 0; i < len(in); i++ {
		rem -= in[i]
	}

	if num == 1 {
		return append(res, append(in, rem))
	}

	for i := 0; i < rem+1; i++ {
		recurseRes := permu(total, num-1, append(in, i))
		res = append(res, recurseRes...)
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
		line := scanner.Text()

		var ing ingredient
		ing.init(line)
		allIngredients = append(allIngredients, &ing)
	}

	var maxValue int
	var maxCombo []int

	/*
		loop through portion permutations
		[0 0 0 100]
		[0 0 1 99]
		[0 0 2 98]
		[0 0 3 97]
		...

		for each portion permutation
			for all ingredients calculate {capacity,durability,flavor,texture,calories}
			determine max score
	*/

	for _, portions := range permu(totalScore, len(allIngredients), []int{}) {
		//fmt.Println(portions)
		var capacity, durability, flavor, texture, calories int
		for i := 0; i < len(allIngredients); i++ {
			capacity += allIngredients[i].capacity * portions[i]
			durability += allIngredients[i].durability * portions[i]
			flavor += allIngredients[i].flavor * portions[i]
			texture += allIngredients[i].texture * portions[i]
			calories += allIngredients[i].calories * portions[i]
		}
		if capacity < 0 {
			capacity = 0
		}
		if durability < 0 {
			durability = 0
		}
		if flavor < 0 {
			flavor = 0
		}
		if texture < 0 {
			texture = 0
		}

		curValue := capacity * durability * flavor * texture

		/*
			// part 1
			if curValue > maxValue {
				maxValue = curValue
				maxCombo = portions
			}
		*/

		// part 2
		if curValue > maxValue && calories == 500 {
			maxValue = curValue
			maxCombo = portions
		}
	}

	fmt.Printf("max value is %d %v\n", maxValue, maxCombo)
}
