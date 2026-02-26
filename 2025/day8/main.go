package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Point struct {
	x, y, z int64
}

type Input []Point

func In(r io.Reader) (res Input) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var p Point
		fmt.Sscanf(line, "%d,%d,%d", &p.x, &p.y, &p.z)
		res = append(res, p)
	}

	return
}

type pair struct {
	i, j  int
	dist2 int64
}

type uf struct {
	parent []int
	size   []int
}

func newUF(n int) *uf {
	u := &uf{parent: make([]int, n), size: make([]int, n)}
	for i := range u.parent {
		u.parent[i] = i
		u.size[i] = 1
	}
	return u
}

func (u *uf) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *uf) union(a, b int) {
	ra, rb := u.find(a), u.find(b)
	if ra == rb {
		return
	}
	if u.size[ra] < u.size[rb] {
		ra, rb = rb, ra
	}
	u.parent[rb] = ra
	u.size[ra] += u.size[rb]
}

func Part1(in Input) (res int) {
	n := len(in)
	pairs := make([]pair, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := in[i].x - in[j].x
			dy := in[i].y - in[j].y
			dz := in[i].z - in[j].z
			pairs = append(pairs, pair{i, j, dx*dx + dy*dy + dz*dz})
		}
	}
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].dist2 < pairs[b].dist2
	})

	u := newUF(n)
	limit := 1000
	if limit > len(pairs) {
		limit = len(pairs)
	}
	for _, p := range pairs[:limit] {
		u.union(p.i, p.j)
	}

	sizes := make(map[int]int)
	for i := 0; i < n; i++ {
		sizes[u.find(i)]++
	}

	top := make([]int, 0, len(sizes))
	for _, s := range sizes {
		top = append(top, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(top)))

	res = 1
	for i := 0; i < 3 && i < len(top); i++ {
		res *= top[i]
	}

	return
}

func Part2(in Input) (res int) {
	return
}

func main() {
	i := In(os.Stdin)
	log.Printf("part1: %d", Part1(i))
	log.Printf("part2: %d", Part2(i))
}
