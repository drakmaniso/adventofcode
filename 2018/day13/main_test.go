package main

import "testing"

func TestPart1(t *testing.T) {
	tracks, carts := read(example)
	x, y := part1(tracks, carts)
	if x != 7 || y != 3 {
		t.Errorf("wrong position (%d,%d instead of %d,%d)", x, y, 7, 3)
	}
}
