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
		x, y, p := part1(ex.serial, 3)
		if x != ex.x {
			t.Errorf("wrong x (%d instead of %d", x, ex.x)
		}
		if y != ex.y {
			t.Errorf("wrong y (%d instead of %d", y, ex.y)
		}
		if p != ex.power {
			t.Errorf("wrong power (%d instead of %d", p, ex.power)
		}
	}
}
