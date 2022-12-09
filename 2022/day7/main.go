package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	cdRoot = iota
	cdUp
	cdDir
	list
	dir
	file
)

type File struct {
	size int
	name string
}

type Dir struct {
	name   string
	size   int
	parent *Dir
	dirs   map[string]*Dir
	files  map[string]*File
}

func (d *Dir) updateBranch(size int) {
	for i := d; i != nil; i = i.parent {
		i.size += size
	}
}

func (d *Dir) sizes() (res []int) {
	res = append(res, d.size)
	for _, v := range d.dirs {
		res = append(res, v.sizes()...)
	}

	return
}

func buildDir(name string, parent *Dir) (res *Dir) {
	res = &Dir{name: name, parent: parent}
	res.dirs = make(map[string]*Dir)
	res.files = make(map[string]*File)

	return
}

func buildFile(size, name string) *File {
	iSize, err := strconv.Atoi(size)
	if err != nil {
		log.Fatal(err)
	}

	return &File{name: name, size: iSize}
}

func In(r io.Reader) (res *Dir) {
	scanner := bufio.NewScanner(r)
	res = buildDir("", nil)

	var cur *Dir
	var inListing bool
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		var op = cdRoot
		if fields[0] == "$" {
			if fields[1] == "cd" {
				if fields[2] == "/" {
					op = cdRoot
				} else if fields[2] == ".." {
					op = cdUp
				} else {
					op = cdDir
				}
			} else if fields[1] == "ls" {
				op = list
			} else {
				log.Fatal("invalid input")
			}
		} else {
			if fields[0] == "dir" {
				op = dir
			} else {
				op = file
			}
		}

		switch op {
		case cdRoot:
			inListing = false
			cur = res
		case cdUp:
			inListing = false
			cur = cur.parent
		case cdDir:
			inListing = false
			cur = cur.dirs[fields[2]]
		case list:
			inListing = true
		case dir:
			if !inListing {
				log.Fatal("invalid input")
			}
			if _, exists := cur.dirs[fields[1]]; exists {
				log.Fatal("invalid input")
			}

			cur.dirs[fields[1]] = buildDir(fields[1], cur)
		case file:
			if !inListing {
				log.Fatal("invalid input")
			}
			if _, exists := cur.files[fields[1]]; exists {
				log.Fatal("invalid input")
			}

			cur.files[fields[1]] = buildFile(fields[0], fields[1])
			cur.updateBranch(cur.files[fields[1]].size)
		}
	}

	return
}

func Part1(in *Dir) (res int) {
	sizes := in.sizes()
	for _, v := range sizes {
		if v <= 100000 {
			res += v
		}
	}

	return
}

func Part2(in *Dir) (res int) {
	used := in.size
	unused := 70000000 - used
	free := 30000000 - unused

	res = math.MaxInt
	sizes := in.sizes()
	for _, v := range sizes {
		if v >= free && v < res {
			res = v
		}
	}

	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
