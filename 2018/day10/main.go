package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	positions, velocities := read(input)
	answer, time := part1(positions, velocities)
	fmt.Printf("Answer for part1: %s, after %d seconds\n", answer, time)
}

func part1(positions, velocities []Coord) (answer string, time int) {
	pos := make([]Coord, 0, len(positions))
	pos = append(pos, positions...)
	for {
		fmt.Printf("\nAfter %d seconds:\n", time)
		s := StringOf(pos)
		if s != "" {
			fmt.Println(s)
			fmt.Print("Enter answer (RETURN to continue simulation): ")
			var answer string
			fmt.Scanf("%s", &answer)
			if answer != "" {
				return answer, time
			}
		}
		for i := range pos {
			pos[i] = pos[i].Plus(velocities[i])
		}
		time++
	}
}

func read(input string) (positions, velocities []Coord) {
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		var p, v Coord
		n, err := fmt.Sscanf(s.Text(), "position=<%d, %d> velocity=<%d, %d>",
			&p.X, &p.Y, &v.X, &v.Y)
		if n != 4 || err != nil {
			log.Printf("WARNING: unable to parse line: %#v: %v", s.Text(), err)
			continue
		}
		positions = append(positions, p)
		velocities = append(velocities, v)
	}
	return positions, velocities
}

// Coord is a pair of 2D cartesian coordinates.
type Coord struct {
	X, Y int
}

// Plus returns the sum of two coordinates.
func (c Coord) Plus(other Coord) Coord {
	return Coord{c.X + other.X, c.Y + other.Y}
}

// Bounds returns the boundries of a list of points.
func Bounds(points []Coord) (min, max Coord) {
	min = points[0]
	for _, p := range points {
		if p.X < min.X {
			min.X = p.X
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	return min, max
}

// StringOf returns a representation of the points.
func StringOf(points []Coord) string {
	min, max := Bounds(points)
	width, height := max.X-min.X+2, max.Y-min.Y+1
	if width > 200 || height > 50 {
		return ""
	}
	s := make([]byte, width*height)
	for i := range s {
		if i%width == width-1 {
			s[i] = '\n'
		} else {
			s[i] = '.'
		}
	}
	for _, p := range points {
		s[p.X-min.X+(p.Y-min.Y)*width] = '#'
	}
	return string(s)
}
