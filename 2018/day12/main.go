package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

type state struct {
	left  []bool
	right []bool
}

func (s *state) String() string {
	res := strings.Builder{}
	for p := s.first(); p <= s.last(); p++ {
		if s.at(p) {
			res.WriteByte('#')
		} else {
			res.WriteByte('.')
		}
	}
	return res.String()
}

func (s *state) StringBetween(first, last int) string {
	res := strings.Builder{}
	for p := first; p <= last; p++ {
		if s.at(p) {
			res.WriteByte('#')
		} else {
			res.WriteByte('.')
		}
	}
	return res.String()
}

func (s *state) first() int {
	return -len(s.left)
}

func (s *state) last() int {
	return len(s.right) - 1
}

func (s *state) at(position int) bool {
	if position < 0 {
		p := 1 - position
		if p >= len(s.left) {
			return false
		}
		return s.left[p]
	}
	if position >= len(s.right) {
		return false
	}
	return s.right[position]
}

func (s *state) set(position int, plant bool) {
	if position < 0 {
		p := 1 - position
		if plant || p < len(s.left) {
			for p >= len(s.left) {
				s.left = append(s.left, false)
			}
			s.left[p] = plant
		}
		return
	}
	if plant || position < len(s.right) {
		for position >= len(s.right) {
			s.right = append(s.right, false)
		}
		s.right[position] = plant
	}
}

func (s *state) configuration(position int) (plants uint) {
	for i := 0; i < 5; i++ {
		if s.at(position - 2 + i) {
			plants |= 1 << uint(i)
		}
	}
	return plants
}

func (s *state) step(rules [32]bool) (new *state) {
	new = &state{}
	for position := s.first() - 2; position <= s.last()+2; position++ {
		c := s.configuration(position)
		new.set(position, rules[c])
	}
	return new
}

func (s *state) sum() int {
	result := 0
	for pos := s.first(); pos <= s.last(); pos++ {
		if s.at(pos) {
			result += pos
		}
	}
	return result
}

func main() {
	initial, rules := read(input)

	fmt.Printf("Answer for part 1: %d\n", simulate(initial, rules, 20))
	// fmt.Printf("Answer for part 2: %d\n", simulate(input, 50000000000))
}

func simulate(initial *state, rules [32]bool, count int) int {
	fmt.Println(initial.StringBetween(-3, 120))
	var new *state
	for i := 0; i < count; i++ {
		new = initial.step(rules)
		// fmt.Println(new.StringBetween(-3, 120))
		initial = new
	}
	fmt.Println(initial.StringBetween(-3, 120))
	fmt.Printf("From %d to %d\n", new.first(), new.last())
	return new.sum()
}

func read(input string) (initial *state, rules [32]bool) {
	initial = &state{}
	s := bufio.NewScanner(strings.NewReader(input))
	s.Scan()
	var ini string
	n, err := fmt.Sscanf(s.Text(), "initial state: %s", &ini)
	if n != 1 || err != nil {
		log.Fatalf("read initial state: %v", err)
	}
	for i, pl := range ini {
		initial.set(i, pl == '#')
	}

	s.Scan()
	for s.Scan() {
		var conf string
		var res rune
		n, err := fmt.Sscanf(s.Text(), "%s => %c", &conf, &res)
		if n != 2 || err != nil {
			log.Printf("read: %v", err)
			continue
		}
		if len(conf) != 5 {
			log.Printf("read: wrong config length (%d)", len(conf))
			continue
		}
		rule := 0
		for i, pl := range conf {
			if pl == '#' {
				rule |= 1 << uint(i)
			}
		}
		if res == '#' {
			rules[rule] = true
		}
	}
	return initial, rules
}
