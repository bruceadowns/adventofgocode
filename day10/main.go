package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// 3113322113
		line := scanner.Text()

		// part 1 = 40
		// part 1 = 50

		for i := 0; i < 50; i++ {
			//fmt.Println(i, line)

			var permu bytes.Buffer
			var prev rune
			var count int

			for k, r := range line {
				if k == 0 || r == prev {
					count++
				} else {
					permu.WriteString(fmt.Sprintf("%d%s", count, string(prev)))
					count = 1
				}
				prev = r
			}
			permu.WriteString(fmt.Sprintf("%d%s", count, string(prev)))
			line = permu.String()
		}

		fmt.Printf("length of 50 lookandsays is %d\n", len(line))
	}
}
