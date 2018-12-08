package main

import "testing"

func TestPart1(t *testing.T) {
	input := []coordinates{
		{1, 1},
		{1, 6},
		{8, 3},
		{3, 4},
		{5, 5},
		{8, 9},
	}
	r, n := part1(input)
	if r != 'E' {
		t.Errorf("Wrong coordinates pair selected: %c instead of %c", r, 'E')
	}
	if n != 17 {
		t.Errorf("Wrong area: %d instead of %d", n, 17)
	}
}

func TestPart2(t *testing.T){
	input := []coordinates{
		{1, 1},
		{1, 6},
		{8, 3},
		{3, 4},
		{5, 5},
		{8, 9},
	}
	s := part2(input, 32)
	if s != 16 {
		t.Errorf("Wrong size: %d instead of %d", s, 16)
	}
}
