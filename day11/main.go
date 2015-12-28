package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// recursively increment alphabetically
func inc(in string, pos int) (out string) {
	r := in[pos] + 1
	out = fmt.Sprintf("%s%s%s", in[:pos], string(r), in[pos+1:])

	if r > 'z' {
		r = 'a'
		out = fmt.Sprintf("%s%s%s", in[:pos], string(r), in[pos+1:])

		if pos > 0 {
			out = inc(out, pos-1)
		}
	}

	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// vzbxkghb
		line := scanner.Text()
		if len(line) != 8 {
			log.Fatalf("invalid input %s\n", line)
		}
		for _, v := range line {
			if v < 'a' || v > 'z' {
				log.Fatalf("invalid input %s\n", line)
			}
		}

		/*
			  exactly eight lowercase letters
				Passwords must include one increasing straight of at least three letters
				  , like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
				Passwords may not contain the letters i, o, or l
				  , as these letters can be mistaken for other characters and are therefore confusing.
				Passwords must contain at least two different
				  , non-overlapping pairs of letters, like aa, bb, or zz.
		*/

		fmt.Printf("current password is %s\n", line)
		for {
			line = inc(line, 7)
			//fmt.Println(line)

			// Passwords must include one increasing straight of at least three letters
			{
				var cond bool
				var prev rune
				var count int
				for _, r := range line {
					if r == prev+1 {
						count++
					} else {
						count = 1
					}

					if count > 2 {
						cond = true
						break
					}

					prev = r
				}
				/*
					// or this may be clearer
					for i := 0; i < len(line)-2; i++ {
						if line[i] == line[i+1]-1 && line[i+1]-1 == line[i+2]-2 {
							cond = true
							break
						}
					}
				*/
				if !cond {
					continue
				}
			}

			// Passwords may not contain the letters i, o, or l
			{
				var cond bool
				for _, r := range line {
					switch r {
					case 'i', 'o', 'l':
						cond = true
						break
					}
				}
				if cond {
					continue
				}
				/*
					// or this may be clearer
					if strings.Contains(line, "i") || strings.Contains(line, "o") || strings.Contains(line, "l") {
						continue
					}
				*/
			}

			// Passwords must contain at least two different
			//	, non-overlapping pairs of letters, like aa, bb, or zz.
			{
				var cond int
				for i := 0; i < len(line)-1; i++ {
					if line[i] == line[i+1] {
						cond++
						i++
					}
				}
				if cond < 2 {
					continue
				}
			}

			fmt.Printf("next password is %s\n", line)
			break
		}
	}
}
