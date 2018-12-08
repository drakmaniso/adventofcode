package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coordinates struct {
	X, Y int
}

func main() {
	input := read()
	letter, answer1 := part1(input)
	fmt.Printf("Answer for part 1: %d (coordinates %c)\n", answer1, letter)
	fmt.Printf("Answer for part2: %d\n", part2(input, 10000))
}

func part1(input []coordinates) (rune, int) {
	min, max := boundaries(input)

	areas := make([]int, len(input))
	for x := min.X - 1; x <= max.X+1; x++ {
		for y := min.Y - 1; y <= max.Y+1; y++ {
			nearest := 0
			shared := false
			best := distance(coordinates{x, y}, input[0])
			for i := 1; i < len(input); i++ {
				d := distance(coordinates{x, y}, input[i])
				if d == best {
					shared = true
				}
				if d < best {
					best = d
					nearest = i
					shared = false
				}
			}
			switch {
			case x == min.X-1 || x == max.X+1 || y == min.Y-1 || y == max.Y+1:
				areas[nearest] = -1
			case !shared && areas[nearest] != -1:
				areas[nearest]++
			}
		}
	}

	answer := 0
	area := areas[0]
	for i := 1; i < len(areas); i++ {
		if areas[i] > area {
			area = areas[i]
			answer = i
		}
	}

	return 'A' + rune(answer), area
}

func part2(input []coordinates, dist int) int {
	min, max := boundaries(input)

	size := 0
	for x := min.X - 1; x <= max.X+1; x++ {
		for y := min.Y - 1; y <= max.Y+1; y++ {
			sum := 0
			for _, c := range input {
				sum += distance(coordinates{x, y}, c)
			}
			if sum < dist {
				size++
			}
		}
	}
	return size
}

func boundaries(input []coordinates) (min, max coordinates) {
	min, max = input[0], input[0]
	for _, c := range input {
		if c.X < min.X {
			min.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
	}
	return min, max
}

func distance(a, b coordinates) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func read() (input []coordinates) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Unable to open input: %v", err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		var c coordinates
		n, err := fmt.Sscanf(s.Text(), "%d, %d", &c.X, &c.Y)
		if n != 2 || err != nil {
			log.Printf("Parsing error: %v", err)
			continue
		}
		input = append(input, c)
	}
	return input
}
