package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution(days int, pool map[int]int) (res int) {
	for i := 0; i < days; i++ {
		var poolNew = make(map[int]int)
		for k, v := range pool {
			k--
			if k < 0 {
				poolNew[6] += v
				poolNew[8] += v
			} else {
				poolNew[k] += v
			}
		}
		pool = poolNew
	}

	for _, v := range pool {
		res += v
	}
	return
}

func In(r io.Reader) (res map[int]int) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		log.Fatal("invalid input")
	}

	res = make(map[int]int)
	line := scanner.Text()
	for _, v := range strings.Split(line, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		res[n]++
	}

	return
}

func main() {
	days := 80
	if len(os.Args) == 2 {
		var err error
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}

	i := In(os.Stdin)
	log.Printf("part1: %d", Solution(days, i))
}
