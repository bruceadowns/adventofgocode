package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

type rep struct {
	from, to     string
	reFrom, reTo *regexp.Regexp
}

func (r rep) String() string {
	return fmt.Sprintf("%s=>%s", r.from, r.to)
}

type reps []rep

func (r reps) Len() int {
	return len(r)
}

func (r reps) Less(i, j int) bool {
	if len(r[i].from) < len(r[j].from) {
		return true
	}

	if len(r[i].to) < len(r[j].to) {
		return true
	}

	return false
}

func (r reps) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *rep) init(line string) {
	//Al => ThF
	n, err := fmt.Sscanf(line, "%s => %s", &r.from, &r.to)
	if n != 2 || err != nil {
		log.Fatalf("invalid input")
	}

	r.reFrom = regexp.MustCompile(r.from)
	r.reTo = regexp.MustCompile(r.to)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// get replacements and seeds
	var rs reps
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			break
		}

		var r rep
		r.init(line)
		rs = append(rs, r)
	}
	//fmt.Println(rs)

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
	for _, v := range rs {
		idx := v.reFrom.FindAllStringIndex(mole, -1)

		for _, w := range idx {
			newmole := mole[:w[0]] + v.to + mole[w[1]:]
			if _, ok := uniqueMoles[newmole]; !ok {
				uniqueMoles[newmole] = struct{}{}
			}
		}
	}
	fmt.Printf("number of unique molecules is %d\n", len(uniqueMoles))

	// sort by len(from+to) descending
	sort.Sort(sort.Reverse(rs))

	var count int
	newmole := mole
	for newmole != "e" {
		for _, v := range rs {
			for {
				indexes := v.reTo.FindStringIndex(newmole)
				if indexes == nil {
					break
				}

				newmole = newmole[:indexes[0]] + v.from + newmole[indexes[1]:]
				//fmt.Println(newmole)
				count++
			}
		}
	}
	fmt.Printf("fewest number of steps is %d\n", count)
}
