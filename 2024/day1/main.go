package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
)

type InputType struct {
	one int
	two int
}

type Input []InputType

func parse(s string) InputType {
	var one, two int
	num, err := fmt.Sscanf(s, "%d %d", &one, &two)
	if err != nil {
		log.Fatal(err)
	}
	if num != 2 {
		log.Fatal("invalid input")
	}
	return InputType{one, two}
}

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, parse(line))
	}

	return
}

func Part1(in Input) int {
	var res float64

	// split to 2 lists
	var listOne, listTwo []int
	for _, v := range in {
		listOne = append(listOne, v.one)
		listTwo = append(listTwo, v.two)
	}

	// sort lists
	sort.Ints(listOne)
	sort.Ints(listTwo)

	// count differences
	for i := 0; i < len(in); i++ {
		res += math.Abs(float64(listOne[i] - listTwo[i]))
	}

	return int(res)
}

func Part2(in Input) (res int) {
	// split to 2 lists
	var listOne, listTwo []int
	for _, v := range in {
		listOne = append(listOne, v.one)
		listTwo = append(listTwo, v.two)
	}

	// loop through listOne
	for _, v := range listOne {
		var count int
		// count number of v's in listTwo
		for _, vv := range listTwo {
			if v == vv {
				count++
			}
		}
		// add num times count
		res += v * count
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
