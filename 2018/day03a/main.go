package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func main() {
	input := read()

	overlaps := map[[2]int]struct{}{}
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			o := input[i].Intersect(input[j])
			if !o.Empty() {
				for x := o.Min.X; x < o.Max.X; x++ {
					for y := o.Min.Y; y < o.Max.Y; y++ {
						overlaps[[2]int{x, y}] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Printf("Answer: %d\n", len(overlaps))
}

func read() (input []image.Rectangle) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var id, x, y, w, h int
		n, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if n < 5 {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			continue
		}
		input = append(input, image.Rectangle{
			Min: image.Point{x, y},
			Max: image.Point{x + w, y + h},
		})
	}
	return input
}
