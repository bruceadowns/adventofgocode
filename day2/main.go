package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func die(msg string) {
	log.Fatalf("invalid input [%s]", msg)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var totalPaper int
	var totalRibbon int

	for scanner.Scan() {
		line := scanner.Text()

		dimensions := strings.Split(line, "x")
		if len(dimensions) != 3 {
			die(line)
		}

		l, err := strconv.Atoi(dimensions[0])
		if err != nil {
			die(line)
		}

		w, err := strconv.Atoi(dimensions[1])
		if err != nil {
			die(line)
		}

		h, err := strconv.Atoi(dimensions[2])
		if err != nil {
			die(line)
		}

		surfaceArea := 2*l*w + 2*w*h + 2*h*l
		minPaper := l * w
		if w*h < minPaper {
			minPaper = w * h
		}
		if h*l < minPaper {
			minPaper = h * l
		}
		totalPaper += surfaceArea + minPaper

		minPerimeter := l + l + w + w
		if w+w+h+h < minPerimeter {
			minPerimeter = w + w + h + h
		}
		if h+h+l+l < minPerimeter {
			minPerimeter = h + h + l + l
		}
		totalRibbon += minPerimeter + l*w*h
	}

	fmt.Printf("total paper %d\n", totalPaper)
	fmt.Printf("total ribbon %d\n", totalRibbon)
}
