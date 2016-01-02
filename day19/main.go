package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type rep struct {
	from, to string
	re       *regexp.Regexp
}

func (r rep) String() string {
	return fmt.Sprintf("%s=>%s", r.from, r.to)
}

func (r *rep) init(line string) {
	//Al => ThF
	n, err := fmt.Sscanf(line, "%s => %s", &r.from, &r.to)
	if n != 2 || err != nil {
		log.Fatalf("invalid input")
	}
	r.re = regexp.MustCompile(r.from)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// get replacements and seeds
	var reps, seeds []rep
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			break
		}

		var r rep
		r.init(line)
		if r.from == "e" {
			seeds = append(seeds, r)
		} else {
			reps = append(reps, r)
		}
	}
	//fmt.Println(reps)
	//fmt.Println(seeds)

	// get medicine molecule
	var mole string
	scanner.Scan()
	//CRnCa...
	mole = scanner.Text()
	if len(mole) < 1 {
		log.Fatalf("invalid input")
	}
	if scanner.Scan() {
		log.Fatalf("invalid input")
	}
	//fmt.Println(mole)

	// step 1
	uniqueMoles := make(map[string]struct{})
	for _, v := range reps {
		idx := v.re.FindAllStringIndex(mole, -1)

		for _, w := range idx {
			newmole := mole[:w[0]] + v.to + mole[w[1]:]
			if _, ok := uniqueMoles[newmole]; !ok {
				uniqueMoles[newmole] = struct{}{}
			}
		}
	}

	fmt.Printf("number of unique molecules is %d\n", len(uniqueMoles))
}
