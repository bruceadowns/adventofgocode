package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
		line := scanner.Text()

		var ing ingredient
		ing.init(line)
		allIngredients = append(allIngredients, &ing)
	}

	// lame permutations algorithm
	var portionPermu [][]int
	for i := 1; i < 100; i++ {
		for j := 1; j < 100-i; j++ {
			for k := 1; k < 100-i-j; k++ {
				l := 100 - k - j - i
				portionPermu = append(portionPermu, []int{i, j, k, l})
			}
		}
	}

	/*
		// test
		var portionPermu [][]int
		for i := 1; i < 100; i++ {
			j := 100 - i
			portionPermu = append(portionPermu, []int{i, j})
		}
	*/

	var maxValue int
	var maxCombo []int
	for _, v := range portionPermu {
		var capacity, durability, flavor, texture, calories int
		for i := 0; i < len(allIngredients); i++ {
			capacity += allIngredients[i].capacity * v[i]
			durability += allIngredients[i].durability * v[i]
			flavor += allIngredients[i].flavor * v[i]
			texture += allIngredients[i].texture * v[i]
			calories += allIngredients[i].calories * v[i]
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
				maxCombo = v
			}
		*/

		// part 2
		if curValue > maxValue && calories == 500 {
			maxValue = curValue
			maxCombo = v
		}
	}

	fmt.Printf("max value is %d %v\n", maxValue, maxCombo)
}
