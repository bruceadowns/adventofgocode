package main

import (
	"log"
	"math"
	"strconv"
)

func part1(in int) (res int) {
	idxElf1 := 0
	idxElf2 := 1
	scoreboard := []int{3, 7}
	for len(scoreboard) < in+10 {
		newRecipe := scoreboard[idxElf1] + scoreboard[idxElf2]
		if newRecipe%10 == newRecipe {
			scoreboard = append(scoreboard, newRecipe)
		} else {
			scoreboard = append(scoreboard, newRecipe/10%10)
			scoreboard = append(scoreboard, newRecipe%10)
		}

		idxElf1 = (idxElf1 + 1 + scoreboard[idxElf1]) % len(scoreboard)
		idxElf2 = (idxElf2 + 1 + scoreboard[idxElf2]) % len(scoreboard)
	}

	for i := 0; i < 10; i++ {
		res += scoreboard[in+i] * int(math.Pow10(9-i))
	}

	return
}

func part2(in int) int {
	var compare []int
	for i := 0; i < len(strconv.Itoa(in)); i++ {
		compare = append(compare, in/int(math.Pow10(i))%10)
	}

	idxElf1 := 0
	idxElf2 := 1
	scoreboard := []int{3, 7}
	for {
		var two bool
		newRecipe := scoreboard[idxElf1] + scoreboard[idxElf2]
		if newRecipe%10 == newRecipe {
			scoreboard = append(scoreboard, newRecipe)
		} else {
			scoreboard = append(scoreboard, newRecipe/10%10)
			scoreboard = append(scoreboard, newRecipe%10)
			two = true
		}

		idxElf1 = (idxElf1 + 1 + scoreboard[idxElf1]) % len(scoreboard)
		idxElf2 = (idxElf2 + 1 + scoreboard[idxElf2]) % len(scoreboard)

		if len(scoreboard) > len(compare) {
			match := true
			for i := 0; i < len(compare); i++ {
				if compare[i] != scoreboard[len(scoreboard)-1-i] {
					match = false
					break
				}
			}
			if match {
				return len(scoreboard) - len(compare)
			}

			if two {
				match = true
				for i := 0; i < len(compare); i++ {
					if compare[i] != scoreboard[len(scoreboard)-2-i] {
						match = false
						break
					}
				}
				if match {
					return len(scoreboard) - len(compare) - 1
				}
			}
		}
	}
}

func main() {
	//i := []int{5, 9, 18, 2018, 293801}
	i := []int{293801}
	for _, v := range i {
		log.Printf("scores of the ten recipes after %d: %010d", v, part1(v))
		log.Printf("number of recipes to the left of sequence %d: %d", v, part2(v))
	}
}
