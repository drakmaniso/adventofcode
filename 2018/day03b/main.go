package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func main() {
	input, claims := read()

loop:
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i != j && input[i].Overlaps(input[j]) {
				continue loop
			}
		}
		fmt.Printf("Answer: %d\n", claims[i])
		return
	}
}

func read() (input []image.Rectangle, claims []int) {
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
		claims = append(claims, id)
	}
	return input, claims
}
