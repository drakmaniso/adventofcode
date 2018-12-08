package main

import (
	"strings"
	"testing"
)

var example = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

func TestRead(t *testing.T) {
	in := read(strings.NewReader(example))
	out := []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}
	eq := len(in) == len(out)
	for i := range in {
		if in[i] != out[i] {
			eq = false
			break
		}
	}
	if !eq {
		t.Errorf("incorrect input read (%v instead of %v)", in, out)
	}
}

func TestPart1(t *testing.T) {
	in := read(strings.NewReader(example))
	_, answer := part1(in, 0)
	if answer != 138 {
		t.Errorf("wronmg answer (%d instead of %d)", answer, 138)
	}
}

// func TestParseNode(t *testing.T) {
// 	in := read(strings.NewReader(example))
// 	_, nodes := parseNode(in, 0, nil)
// 	out := []node{node{children:[]int{1, 2}, metadata:[]int{1, 1, 2}}, node{children:[]int(nil), metadata:[]int{10, 11,12}}, node{children:[]int{3}, metadata:[]int{2}}, node{children:[]int(nil), metadata:[]int{99}}}
// 	fmt.Printf("%#v\n", nodes)
// 	t.Errorf("wrong parsing")
// }

func TestPart2(t *testing.T) {
	in := read(strings.NewReader(example))
	out := part2(in)
	if out != 66 {
		t.Errorf("wrong answer (%d instead of %d)", out, 66)
	}
}
