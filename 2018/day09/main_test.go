package main

import "testing"

var examples = []struct {
	input                   string
	playerCount, lastMarble int
	highscore               int
}{
	{"9 players; last marble is worth 25 points", 9, 25, 32},
	{"10 players; last marble is worth 1618 points", 10, 1618, 8317},
	{"13 players; last marble is worth 7999 points", 13, 7999, 146373},
	{"17 players; last marble is worth 1104", 17, 1104, 2764},
	{"21 players; last marble is worth 6111 points", 21, 6111, 54718},
	{"30 players; last marble is worth 5807 points", 30, 5807, 37305},
}

func TestRead(t *testing.T) {
	for i := range examples {
		p, m := read(examples[i].input)
		if p != examples[i].playerCount {
			t.Errorf("wrong player count (%d instead of %d)", p, examples[i].playerCount)
		}
		if m != examples[i].lastMarble {
			t.Errorf("wrong last marble (%d instead of %d)", p, examples[i].lastMarble)
		}
	}
}

func TestPart1(t *testing.T) {
	for i := range examples {
		highscore := part1(examples[i].input)
		if highscore != examples[i].highscore {
			t.Errorf("wrong player count (%d instead of %d)", highscore, examples[i].highscore)
		}
	}
}
