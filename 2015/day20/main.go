package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// 33100000
		line := scanner.Text()
		target, err := strconv.Atoi(line)
		for err != nil {
			log.Fatalf("invalid input %s", line)
		}
		fmt.Printf("presents target is %d\n", target)

		{
			houses := make([]int, target/10)
			for elf := 1; elf < len(houses); elf++ {
				for house := elf; house < len(houses); house += elf {
					houses[house] += elf * 10
				}
			}

			for k, v := range houses {
				if v > target {
					fmt.Printf("House %d got %d presents using first formula.\n", k, v)
					break
				}
			}
		}

		{
			houses := make([]int, target/10)
			for elf := 1; elf < len(houses); elf++ {
				var num int
				for house := elf; house < len(houses); house += elf {
					houses[house] += elf * 11
					num++
					if num > 50 {
						break
					}
				}
			}

			for k, v := range houses {
				if v > target {
					fmt.Printf("House %d got %d presents using second formula.\n", k, v)
					break
				}
			}
		}
	}
}
