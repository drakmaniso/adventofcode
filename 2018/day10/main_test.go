package main

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	expPos := []Coord{Coord{X: 9, Y: 1}, Coord{X: 7, Y: 0}, Coord{X: 3, Y: -2}, Coord{X: 6, Y: 10}, Coord{X: 2, Y: -4}, Coord{X: -6, Y: 10}, Coord{X: 1, Y: 8}, Coord{X: 1, Y: 7}, Coord{X: -3, Y: 11}, Coord{X: 7, Y: 6}, Coord{X: -2, Y: 3}, Coord{X: -4, Y: 3}, Coord{X: 10, Y: -3}, Coord{X: 5, Y: 11}, Coord{X: 4, Y: 7}, Coord{X: 8, Y: -2}, Coord{X: 15, Y: 0}, Coord{X: 1, Y: 6}, Coord{X: 8, Y: 9}, Coord{X: 3, Y: 3}, Coord{X: 0, Y: 5}, Coord{X: -2, Y: 2}, Coord{X: 5, Y: -2}, Coord{X: 1, Y: 4}, Coord{X: -2, Y: 7}, Coord{X: 3, Y: 6}, Coord{X: 5, Y: 0}, Coord{X: -6, Y: 0}, Coord{X: 5, Y: 9}, Coord{X: 14, Y: 7}, Coord{X: -3, Y: 6}}
	expVel := []Coord{Coord{X: 0, Y: 2}, Coord{X: -1, Y: 0}, Coord{X: -1, Y: 1}, Coord{X: -2, Y: -1}, Coord{X: 2, Y: 2}, Coord{X: 2, Y: -2}, Coord{X: 1, Y: -1}, Coord{X: 1, Y: 0}, Coord{X: 1, Y: -2}, Coord{X: -1, Y: -1}, Coord{X: 1, Y: 0}, Coord{X: 2, Y: 0}, Coord{X: -1, Y: 1}, Coord{X: 1, Y: -2}, Coord{X: 0, Y: -1}, Coord{X: 0, Y: 1}, Coord{X: -2, Y: 0}, Coord{X: 1, Y: 0}, Coord{X: 0, Y: -1}, Coord{X: -1, Y: 1}, Coord{X: 0, Y: -1}, Coord{X: 2, Y: 0}, Coord{X: 1, Y: 2}, Coord{X: 2, Y: 1}, Coord{X: 2, Y: -2}, Coord{X: -1, Y: -1}, Coord{X: 1, Y: 0}, Coord{X: 2, Y: 0}, Coord{X: 1, Y: -2}, Coord{X: -2, Y: 0}, Coord{X: 2, Y: -1}}
	positions, velocities := read(example)
	if !reflect.DeepEqual(positions, expPos) {
		t.Errorf("wrong positions parsed")
	}
	if !reflect.DeepEqual(velocities, expVel) {
		t.Errorf("wrong velocities parsed")
	}
}

func TestPart1(t *testing.T) {
	positions, velocities := read(example)
	answer, time := part1(positions, velocities)
	if answer != "HI" {
		t.Errorf("wrong answer")
	}
	if time != 3 {
		t.Errorf("wrong time")
	}
}
