package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type entry struct {
	date   time.Time
	action action
	guard  int
}

type action int

const (
	beginShift action = iota
	fallAsleep
	wakeUp
)

func main() {
	entries := read()

	var sleepyguard int
	asleep := map[int]int{}
	var guard, from int
	for _, e := range entries {
		switch e.action {
		case beginShift:
			guard = e.guard
		case fallAsleep:
			from = e.date.Minute()
		case wakeUp:
			t := e.date.Minute() - from
			asleep[guard] += t
			if asleep[guard] > asleep[sleepyguard] {
				sleepyguard = guard
			}
		}
	}

	minutes := [60]int{}
	guard = -1
	var sleepyminute int
	for _, e := range entries {
		if e.action == beginShift {
			guard = e.guard
			continue
		}
		if guard != sleepyguard {
			continue
		}
		switch e.action {
		case fallAsleep:
			from = e.date.Minute()
		case wakeUp:
			to := e.date.Minute()
			for i := from; i < to; i++ {
				minutes[i]++
				if minutes[i] > minutes[sleepyminute] {
					sleepyminute = i
				}
			}
		}
	}

	fmt.Printf("Answer: guard %d * minute %d = %d\n",
		sleepyguard, sleepyminute, sleepyguard*sleepyminute)
}

func read() []entry {
	entries := []entry{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		txt := s.Text()
		e := entry{guard: -1}

		var y, m, d, hr, mn int
		n, err := fmt.Sscanf(txt, "[%d-%d-%d %d:%d]", &y, &m, &d, &hr, &mn)
		if n < 5 || err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			continue
		}
		e.date = time.Date(y, time.Month(m), d, hr, mn, 0, 0, time.UTC)

		i := strings.Index(txt, "] ")
		if i == -1 {
			fmt.Fprintf(os.Stderr, "ERROR: unable to parse message\n")
			continue
		}
		txt = txt[i+2:]
		n, err = fmt.Sscanf(txt, "Guard #%d begins shift", &e.guard)
		switch {
		case n == 1:
			e.action = beginShift
		case txt == "falls asleep":
			e.action = fallAsleep
		case txt == "wakes up":
			e.action = wakeUp
		default:
			fmt.Fprintf(os.Stderr, "ERROR: unknown action\n")
			continue
		}
		entries = append(entries, e)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].date.Before(entries[j].date)
	})
	return entries
}
