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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var groupLayer, groupSum int
		var inGarbageBlock, negatorOn bool
		var garbageCount int

		for _, v := range line {
			if negatorOn {
				negatorOn = false
				continue
			}

			switch v {
			case beginGroupBlock:
				if inGarbageBlock {
					garbageCount++
				} else {
					groupLayer++
				}
			case endGroupBlock:
				if inGarbageBlock {
					garbageCount++
				} else {
					groupSum += groupLayer
					groupLayer--
				}
			case beginGarbageBlock:
				if inGarbageBlock {
					garbageCount++
				} else {
					inGarbageBlock = true
				}
			case endGarbageBlock:
				inGarbageBlock = false
			case negator:
				negatorOn = true
			default:
				if inGarbageBlock {
					garbageCount++
				}
			}
		}
		if inGarbageBlock || negatorOn || groupLayer != 0 {
			log.Fatal("invalid logic")
		}

		log.Printf("part1 group sum: %d", groupSum)
		log.Printf("part2 garbage count: %d", garbageCount)
	}
}
