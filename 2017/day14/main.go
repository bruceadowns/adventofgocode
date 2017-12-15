package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	gridSize = 128
	keySize  = 256
)

type knothash struct {
	pos  int
	list []int
}

func builder(keyLen int) (res knothash) {
	res.list = make([]int, keyLen)
	for i := 0; i < len(res.list); i++ {
		res.list[i] = i
	}

	return
}

func (kl *knothash) advance(i int) {
	kl.pos = (kl.pos + i) % len(kl.list)
}

func (kl knothash) calcEnd(i int) int {
	return (kl.pos + i - 1) % len(kl.list)
}

func (kl *knothash) knot(i int) error {
	if i < 0 || i > len(kl.list) {
		return fmt.Errorf("invalid logic i: %d len: %d", i, len(kl.list))
	}
	if i == 0 || i == 1 {
		return nil
	}

	end := kl.calcEnd(i)
	if kl.pos < end {
		slist := kl.list[kl.pos : end+1]
		for i := 0; i < len(slist)/2; i++ {
			slist[i], slist[len(slist)-i-1] = slist[len(slist)-i-1], slist[i]
		}
	} else if kl.pos > end {
		slist := append(kl.list[kl.pos:], kl.list[:end+1]...)
		for i := 0; i < len(slist)/2; i++ {
			slist[i], slist[len(slist)-i-1] = slist[len(slist)-i-1], slist[i]
		}

		idx := 0
		for i := kl.pos; i < len(kl.list); i++ {
			kl.list[i] = slist[idx]
			idx++
		}
		for i := 0; i < end+1; i++ {
			kl.list[i] = slist[idx]
			idx++
		}
	} else {
		return fmt.Errorf("invalid logic pos: %d end: %d", kl.pos, end)
	}

	return nil
}

type square struct {
	used    bool
	visited bool
}

type diskRow []square

type diskGrid []diskRow

// recursive DFS visitor
func visitRegion(d diskGrid, x int, y int) {
	if !d[x][y].used || d[x][y].visited {
		log.Fatalf("invalid logic square: %s", d[x][y])
	}

	d[x][y].visited = true

	// north
	if y > 0 && d[x][y-1].used && !d[x][y-1].visited {
		visitRegion(d, x, y-1)
	}
	// south
	if y < len(d[x])-1 && d[x][y+1].used && !d[x][y+1].visited {
		visitRegion(d, x, y+1)
	}
	// west
	if x > 0 && d[x-1][y].used && !d[x-1][y].visited {
		visitRegion(d, x-1, y)
	}
	// east
	if x < len(d)-1 && d[x+1][y].used && !d[x+1][y].visited {
		visitRegion(d, x+1, y)
	}

	return
}

func main() {
	// input is key prefix
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	keyPrefix := scanner.Text()
	if len(keyPrefix) < 1 {
		log.Fatal("empty input")
	}

	// generate keys
	keys := make([]*bytes.Buffer, 0)
	for i := 0; i < gridSize; i++ {
		b := &bytes.Buffer{}
		b.WriteString(keyPrefix + "-" + strconv.Itoa(i))
		keys = append(keys, b)
	}

	// make disk grid
	dGrid := make(diskGrid, gridSize)
	for x := 0; x < gridSize; x++ {
		dGrid[x] = make(diskRow, gridSize)
	}

	// generate knot hash per key
	for keyRow, keyValue := range keys {
		// key is stream of ascii bytes
		nLens := make([]int, 0)
		for _, v := range keyValue.Bytes() {
			nLens = append(nLens, int(v))
		}
		nLens = append(nLens, []int{17, 31, 73, 47, 23}...)

		// build knot hash
		khash := builder(keySize)
		skip := 0
		for i := 0; i < 64; i++ {
			for _, nLen := range nLens {
				if err := khash.knot(nLen); err != nil {
					log.Fatal(err)
				}

				khash.advance(nLen + skip)
				skip++
			}
		}

		// calc and encode dense hash
		hashBin := bytes.Buffer{}
		for i := 0; i < 16; i++ {
			var denseHash int
			for j := 0; j < 16; j++ {
				denseHash ^= khash.list[i*16+j]
			}

			hashBin.WriteString(fmt.Sprintf("%08b", denseHash))
		}

		// populate disk grid from binary dense hash
		for k, v := range hashBin.Bytes() {
			if v == '1' {
				dGrid[keyRow][k].used = true
			}
		}
	}

	// calc used sum and region count
	sum := 0
	regions := 0
	for x := range dGrid {
		for y := range dGrid[x] {
			if dGrid[x][y].used {
				sum++

				if !dGrid[x][y].visited {
					regions++
					visitRegion(dGrid, x, y)
				}
			}
		}
	}

	log.Printf("part1 number of squares: %d", sum)
	log.Printf("part2 number of regions: %d", regions)
}
