package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := read()

	twos, threes := 0, 0
	counts := [26]byte{}
	for _, id := range input {

		counts = [26]byte{}
		for _, r := range id {
			if r < 'a' || r > 'z' {
				fmt.Fprintf(os.Stderr, "ERROR: not a letter\n")
				continue
			}
			counts[r-'a']++
		}

		has2, has3 := false, false
		for _, c := range counts {
			switch c {
			case 2:
				has2 = true
			case 3:
				has3 = true
			}
		}
		if has2 {
			twos++
		}
		if has3 {
			threes++
		}

	}

	fmt.Printf("Answer: %d\n", twos*threes)
}

func read() (input []string) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input = append(input, s.Text())
	}
	return input
}
