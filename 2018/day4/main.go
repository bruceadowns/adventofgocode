package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type action int

const (
	actionBegin = iota
	actionFallsAsleep
	actionWakesUp
)

type state int

const (
	stateUnknown = iota
	stateBegun
	stateAsleep
	stateAwake
)

type record struct {
	t time.Time
	a action
	g int
}

func (r record) String() string {
	return fmt.Sprintf("[%s] %d %d", r.t, r.a, r.g)
}

type records []record

// sort records by time field

func (r records) Len() int {
	return len(r)
}

func (r records) Less(i, j int) bool {
	return r[i].t.Before(r[j].t)
}

func (r records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func rollState(r records) (res map[int]map[int]int) {
	// map[guard]map[minute]count
	res = make(map[int]map[int]int)

	var guard int
	var startAsleep time.Time
	var st state
	for _, v := range r {
		switch st {
		case stateUnknown:
			switch v.a {
			case actionBegin:
				st = stateBegun
				guard = v.g
			default:
				log.Fatalf("invalid state transition: %d -> %d", st, v.a)
			}
		case stateBegun, stateAwake:
			switch v.a {
			case actionBegin:
				st = stateBegun
				guard = v.g
			case actionFallsAsleep:
				st = stateAsleep
				startAsleep = v.t
			default:
				log.Fatalf("invalid state transition: %d -> %d", st, v.a)
			}
		case stateAsleep:
			switch v.a {
			case actionWakesUp:
				st = stateAwake

				if _, ok := res[guard]; !ok {
					res[guard] = make(map[int]int)
				}
				for i := startAsleep.Minute(); i < v.t.Minute(); i++ {
					res[guard][i]++
				}
			default:
				log.Fatalf("invalid state transition: %d -> %d", st, v.a)
			}
		default:
			log.Fatalf("unknown state: %d", st)
		}
	}

	return
}

func part1(r records) (res int) {
	m := rollState(r)

	// which guard has most minutes asleep
	// for that guard, which of his minutes is he asleep the most

	var maxMinutesAsleep, maxMinutesAsleepGuard, maxMinutesAsleepGuardMinute int
	for k, v := range m {
		var maxMinutesAsleepPerGuard, maxMinutesAsleepGuardPerGuard, maxMinutesAsleepGuardMinutePerGuard int
		for kk, vv := range v {
			maxMinutesAsleepPerGuard += vv

			if vv > maxMinutesAsleepGuardPerGuard {
				maxMinutesAsleepGuardMinutePerGuard = kk
				maxMinutesAsleepGuardPerGuard = vv
			}
		}
		if maxMinutesAsleepPerGuard > maxMinutesAsleep {
			maxMinutesAsleep = maxMinutesAsleepPerGuard
			maxMinutesAsleepGuard = k
			maxMinutesAsleepGuardMinute = maxMinutesAsleepGuardMinutePerGuard
		}
	}

	return maxMinutesAsleepGuard * maxMinutesAsleepGuardMinute
}

func part2(r records) (res int) {
	m := rollState(r)

	// which guard is most frequently asleep on the same minute
	// and for which minute

	var maxMinutesAsleep, maxMinutesAsleepGuard, maxMinutesAsleepGuardMinute int
	for k, v := range m {
		for kk, vv := range v {
			if vv > maxMinutesAsleep {
				maxMinutesAsleep = vv
				maxMinutesAsleepGuard = k
				maxMinutesAsleepGuardMinute = kk
			}
		}
	}

	return maxMinutesAsleepGuard * maxMinutesAsleepGuardMinute
}

func build(s string) (res record) {
	//[1518-11-05 00:03] Guard #99 begins shift
	//[1518-11-05 00:45] falls asleep
	//[1518-11-05 00:55] wakes up

	s = s[1:]

	var err error
	res.t, err = time.Parse("2006-01-02 15:04", s[:16])
	if err != nil {
		log.Fatal(err)
	}

	s = s[18:]

	if s == "falls asleep" {
		res.a = actionFallsAsleep
	} else if s == "wakes up" {
		res.a = actionWakesUp
	} else if strings.HasPrefix(s, "Guard ") {
		res.a = actionBegin
		num, err := fmt.Sscanf(s, "Guard #%d begins shift", &res.g)
		if err != nil {
			log.Fatal(err)
		}
		if num != 1 {
			log.Fatalf("invalid input: %s", s)
		}
	} else {
		log.Fatalf("invalid input: %s", s)
	}

	return
}

func in(r io.Reader) (res records) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		res = append(res, build(scanner.Text()))
	}
	sort.Sort(res)

	return
}

func main() {
	r := in(os.Stdin)
	log.Printf("guard id *times* his max asleep minute: %d", part1(r))
	log.Printf("guard id *times* the minute he is asleep the most: %d", part2(r))
}
