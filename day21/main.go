package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type weapon struct {
	name   string
	cost   int
	damage int
}

type armor struct {
	name       string
	cost       int
	protection int
}

type ring struct {
	name   string
	cost   int
	damage int
	armor  int
}

type player struct {
	points     int
	damage     int
	protection int
	rings      []ring
	cost       int
}

type players []player

func die() {
	log.Fatal("flaw occurred")
}

func (p players) Len() int {
	return len(p)
}

func (p players) Less(i, j int) bool {
	return p[i].cost < p[j].cost
}

func (p players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// https://docs.python.org/2/library/itertools.html
// https://github.com/ntns/goitertools
func combos(iterable []ring, r int) [][]ring {
	pool := iterable
	n := len(pool)

	if r > n || r == 0 {
		return nil
	}

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]ring, r)
	for i, el := range indices {
		result[i] = pool[el]
	}

	results := [][]ring{result}

	for {
		i := r - 1
		for ; i >= 0 && indices[i] == i+n-r; i-- {
		}

		if i < 0 {
			return results
		}

		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}

		result := make([]ring, r)
		for i = 0; i < len(indices); i++ {
			result[i] = pool[indices[i]]
		}

		results = append(results, result)

	}
}

func input(bossHitPoints, bossDamage, bossArmor *int) {
	scanner := bufio.NewScanner(os.Stdin)

	{
		if !scanner.Scan() {
			die()
		}
		line := scanner.Text()
		if len(line) < 1 {
			die()
		}
		n, err := fmt.Sscanf(line, "Hit Points: %d", bossHitPoints)
		if n != 1 || err != nil {
			die()
		}
	}

	{
		if !scanner.Scan() {
			die()
		}
		line := scanner.Text()
		if len(line) < 1 {
			die()
		}
		n, err := fmt.Sscanf(line, "Damage: %d", bossDamage)
		if n != 1 || err != nil {
			die()
		}
	}

	{
		if !scanner.Scan() {
			die()
		}
		line := scanner.Text()
		if len(line) < 1 {
			die()
		}
		n, err := fmt.Sscanf(line, "Armor: %d", bossArmor)
		if n != 1 || err != nil {
			die()
		}
	}

	if scanner.Scan() {
		die()
	}
}

func win(p, b player) bool {
	//fmt.Println()

	for {
		//fmt.Println(p, b)

		// player attack
		{
			damage := p.damage - b.protection
			if damage < 1 {
				damage = 1
			}

			b.points -= damage
			if b.points < 1 {
				return true
			}
		}

		// boss attack
		{
			damage := b.damage - p.protection
			if damage < 1 {
				damage = 1
			}

			p.points -= damage
			if p.points < 1 {
				return false
			}
		}
	}
}

func main() {
	/*
		Hit Points: 100
		Damage: 8
		Armor: 2
	*/
	var bossDamage, bossArmor, bossHitPoints int
	input(&bossHitPoints, &bossDamage, &bossArmor)
	//fmt.Printf("%d %d %d\n", bossHitPoints, bossDamage, bossArmor)

	weapons := []weapon{
		{
			name:   "Dagger",
			cost:   8,
			damage: 4,
		},
		{
			name:   "Shortsword",
			cost:   10,
			damage: 5,
		},
		{
			name:   "Warhammer",
			cost:   25,
			damage: 6,
		},
		{
			name:   "Longsword",
			cost:   40,
			damage: 7,
		},
		{
			name:   "Greataxe",
			cost:   74,
			damage: 8,
		},
	}

	armors := []armor{
		{
			name: "None",
		},
		{
			name:       "Leather",
			cost:       13,
			protection: 1,
		},
		{
			name:       "Chainmail",
			cost:       31,
			protection: 2,
		},
		{
			name:       "Splintmail",
			cost:       53,
			protection: 3,
		},
		{
			name:       "Bandedmail",
			cost:       75,
			protection: 4,
		},
		{
			name:       "Platemail",
			cost:       102,
			protection: 5,
		},
	}

	rings := []ring{
		{
			name: "Damage 0",
		},
		{
			name: "Defense 0",
		},
		{
			name:   "Damage +1",
			cost:   25,
			damage: 1,
		},
		{
			name:  "Defense +1",
			cost:  20,
			armor: 1,
		},
		{
			name:   "Damage +2",
			cost:   50,
			damage: 2,
		},
		{
			name:  "Defense +2",
			cost:  40,
			armor: 2,
		},
		{
			name:   "Damage +3",
			cost:   100,
			damage: 3,
		},
		{
			name:  "Defense +2",
			cost:  80,
			armor: 3,
		},
	}

	var permu players
	for _, w := range weapons {
		for _, a := range armors {
			for _, r := range combos(rings, 2) {
				permu = append(permu,
					player{
						points:     100,
						damage:     w.damage + r[0].damage + r[1].damage,
						protection: a.protection + r[0].armor + r[1].armor,
						rings:      r,
						cost:       w.cost + a.cost + r[0].cost + r[1].cost,
					})
			}
		}
	}

	sort.Sort(permu)

	boss := player{
		points:     bossHitPoints,
		damage:     bossDamage,
		protection: bossArmor,
	}

	var least *player
	for _, v := range permu {
		if win(v, boss) {
			least = &v
			break
		}
	}

	if least == nil {
		die()
	}
	fmt.Printf("least cost to win is %d\n", least.cost)
	//fmt.Printf("%v\n", *least)

	// part 2
	sort.Sort(sort.Reverse(permu))

	var most *player
	for _, v := range permu {
		if !win(v, boss) {
			most = &v
			break
		}
	}

	if most == nil {
		die()
	}
	fmt.Printf("most cost to lose is %d\n", most.cost)
	//fmt.Printf("%v\n", *most)
}
