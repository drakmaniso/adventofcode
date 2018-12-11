package main

import "testing"

func TestPower(t *testing.T) {
	for _, v := range subexamples {
		p := power(v.x, v.y, v.serial)
		if p != v.power {
			t.Errorf("wrong power (%d instead of %d)", p, v.power)
		}
	}
}

func TestPart1(t *testing.T) {
	for _, ex := range examples {
		x, y, p := part1(ex.serial)
		if x != ex.x || y != ex.y {
			t.Errorf("wrong cluster (%d,%d instead of %d,%d", x, y, ex.x, ex.y)
		}
		if p != ex.power {
			t.Errorf("wrong power (%d instead of %d", p, ex.power)
		}
	}
}

func TestPart2(t *testing.T) {
	for _, ex := range examples2 {
		x, y, s, p := part2(ex.serial)
		if x!= ex.x || y != ex.y || s != ex.size {
			t.Errorf("wrong cluster (%d,%d,%d instead of %d,%d,%d)", x, y, s, ex.x, ex.y, ex.size)
		}
		if p != ex.power {
			t.Errorf("wrong power (%d instead of %d)", p, ex.power)
		}
	}
}
