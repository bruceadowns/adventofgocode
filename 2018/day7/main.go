package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type input struct {
	src rune
	dep rune
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func mapify(in []input) (res map[rune]map[rune]struct{}) {
	res = make(map[rune]map[rune]struct{})
	for _, v := range in {
		if _, ok := res[v.src]; !ok {
			res[v.src] = make(map[rune]struct{})
		}
		if k, ok := res[v.dep]; ok {
			k[v.src] = struct{}{}
		} else {
			res[v.dep] = make(map[rune]struct{})
			res[v.dep][v.src] = struct{}{}
		}
	}

	return
}

func part1(in []input) string {
	depMap := mapify(in)

	var sb strings.Builder
	for len(depMap) > 0 {
		opts := make([]rune, 0)
		for k, v := range depMap {
			if len(v) == 0 {
				opts = append(opts, k)
			}
		}
		sort.Sort(RuneSlice(opts))

		d := opts[0]
		sb.WriteRune(d)
		delete(depMap, d)
		for _, v := range depMap {
			delete(v, d)
		}
	}

	return sb.String()
}

func part2(in []input) (res int) {
	depMap := mapify(in)

	workers := make(map[rune]int)
	for len(depMap) > 0 {
		opts := make([]rune, 0)
		for k, v := range depMap {
			if len(v) == 0 {
				opts = append(opts, k)
			}
		}
		sort.Sort(RuneSlice(opts))

		for i := 0; i < len(opts) && len(workers) < 6; i++ {
			if _, ok := workers[opts[i]]; !ok {
				workers[opts[i]] = int(60 + opts[i] - 'A' + 1)
			}
		}

		res++
		var done []rune
		for k := range workers {
			workers[k]--
			if workers[k] == 0 {
				done = append(done, k)
			}
		}
		for _, v := range done {
			delete(workers, v)
			delete(depMap, v)
			for _, vv := range depMap {
				delete(vv, v)
			}
		}
	}

	return
}

func in(r io.Reader) (res []input) {
	res = make([]input, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var src, dep rune
		num, err := fmt.Sscanf(
			scanner.Text(),
			"Step %c must be finished before step %c can begin.",
			&src, &dep)
		if err != nil {
			log.Fatal(err)
		}
		if num != 2 {
			log.Fatal("invalid input")
		}

		res = append(res, input{src, dep})
	}

	return
}

func main() {
	i := in(os.Stdin)
	log.Printf("order: %s", part1(i))
	log.Printf("completion time: %d", part2(i))
}
