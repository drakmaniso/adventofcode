package main

import "testing"

func stepOf(r rune) step {
	return step(r-'A')
}

var example = [][2]step{
	{stepOf('C'), stepOf('A')},
	{stepOf('C'), stepOf('F')},
	{stepOf('A'), stepOf('B')},
	{stepOf('A'), stepOf('D')},
	{stepOf('B'), stepOf('E')},
	{stepOf('D'), stepOf('E')},
	{stepOf('F'), stepOf('E')},
}

func TestPart1(t *testing.T) {
	answer := part1(example)
	if answer != "CABDFE" {
		t.Errorf("wrong answer (%#v instead of %#v)", answer, "CABDFE")
	}
}

func TestPart2(t *testing.T) {
	answer := part2(example, 2, 0)
	if answer != 15 {
		t.Errorf("wrong answer (%d instead of %d)", answer, 15)
	}
}
