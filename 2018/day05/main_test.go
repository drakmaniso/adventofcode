package main

import "testing"

func TestPart1(t *testing.T) {
	examples := []struct {
		in  []byte
		out []byte
	}{
		{in: []byte("aA"), out: []byte("")},
		{in: []byte("abBA"), out: []byte("")},
		{in: []byte("abAB"), out: []byte("abAB")},
		{in: []byte("aabAAB"), out: []byte("aabAAB")},
		{in: []byte("dabAcCaCBAcCcaDA"), out: []byte("dabCBAcaDA")},
	}
	for i, e := range examples {
		a := part1(e.in)
		diff := len(a) != len(e.out)
		for i := range a {
			if a[i] != e.out[i] {
				diff = true
				break
			}
		}
		if diff {
			t.Errorf("Fail for example %d %#v: got %#v instead of %#v",
				i, e.in, part1(e.in), e.out)
		}
	}
}

func TestStrip(t *testing.T) {
	examples := []struct {
		in  []byte
		strip byte
		out []byte
	}{
		{in: []byte("dabAcCaCBAcCcaDA"), strip:'a', out: []byte("dbcCCBcCcD")},
		{in: []byte("dabAcCaCBAcCcaDA"), strip:'b', out: []byte("daAcCaCAcCcaDA")},
		{in: []byte("dabAcCaCBAcCcaDA"), strip:'c', out: []byte("dabAaBAaDA")},
		{in: []byte("dabAcCaCBAcCcaDA"), strip:'d', out: []byte("abAcCaCBAcCcaA")},
	}
	for i, e := range examples {
		a := strip(e.in, e.strip)
		diff := len(a) != len(e.out)
		for i := range a {
			if a[i] != e.out[i] {
				diff = true
				break
			}
		}
		if diff {
			t.Errorf("Fail for example %d %#v: got %#v instead of %#v",
				i, string(e.in), string(part1(e.in)), string(e.out))
		}
	}
}

func TestPart2(t *testing.T) {
	u, l := part2([]byte("dabAcCaCBAcCcaDA"))
	if u != 'c' {
		t.Errorf("Fail: wrong unit %c instead of %c", u, 'c')
	}
	if l != 4 {
		t.Errorf("Fail: wrong length %d instead of %d", l, 4)
	}
}
