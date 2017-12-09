package main

import (
	"bufio"
	"log"
	"os"
)

const (
	beginGroupBlock   = '{'
	endGroupBlock     = '}'
	beginGarbageBlock = '<'
	endGarbageBlock   = '>'
	negator           = '!'
)

const (
	def = iota
	garbage
	ignoreDefault
	ignoreGarbage
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var groupLayer, groupSum int
		var garbageCount int

		state := def
		for _, v := range line {
			switch state {
			case def:
				switch v {
				case beginGroupBlock:
					groupLayer++
				case endGroupBlock:
					groupSum += groupLayer
					groupLayer--
				case beginGarbageBlock:
					state = garbage
				case negator:
					state = ignoreDefault
				}
			case garbage:
				switch v {
				case endGarbageBlock:
					state = def
				case negator:
					state = ignoreGarbage
				default:
					garbageCount++
				}
			case ignoreDefault:
				state = def
			case ignoreGarbage:
				state = garbage
			}
		}
		if state != def || groupLayer != 0 {
			log.Fatalf("invalid logic: %d %d", state, groupLayer)
		}

		log.Printf("part1 group sum: %d", groupSum)
		log.Printf("part2 garbage count: %d", garbageCount)
	}
}
