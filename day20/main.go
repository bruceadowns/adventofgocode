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
		presentsTarget, err := strconv.Atoi(line)
		for err != nil {
			log.Fatalf("invalid input %s", line)
		}
		fmt.Printf("presents target is %d\n", presentsTarget)

		var house int
		var presentsMax int
		elfDelivery := make(map[int]int)
		for {
			house++

			var presents int
			for elf := house; elf > 0; elf-- {
				if house%elf == 0 {
					elfDelivery[elf]++
					if elfDelivery[elf] < 51 {
						presents += 11 * elf
					}
				}
			}

			if presents > presentsMax {
				presentsMax = presents
				fmt.Printf("House %d got %d presents.\n", house, presents)
			}

			if presents >= presentsTarget {
				break
			}
		}

		fmt.Printf("House %d got %d presents.\n", house, presentsMax)
	}
}
