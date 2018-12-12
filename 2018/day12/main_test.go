package main

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	initial, rules := read(example)

	if !reflect.DeepEqual(*initial, state{left: []bool(nil), right: []bool{true, false, false, true, false, true, false, false, true, true, false, false, false, false, false, false, true, true, true, false, false, false, true, true, true}}) {
		t.Errorf("wrong initial state\n")
	}

	if !reflect.DeepEqual(rules, [32]bool{false, false, true, false, true, false, true, true, false, false, true, true, false, false, false, true, false, false, false, false, false, true, false, true, true, false, true, true, false, true, true, false}) {
		t.Errorf("wrong rules")
	}
}

func TestPart1(t *testing.T) {
	answer := simulate(example, 20)
	if answer != 325 {
		t.Errorf("wrong answer (%d instead of 325)", answer)
	}
}
