package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input := read()
	fmt.Printf("Answer for part1: %d\n", len(part1(input)))
	u, l := part2(input)
	fmt.Printf("Answer for part2: unit %c -> length %d\n", u, l)
}

func part1(input []byte) []byte {
	in := append(make([]byte, 0, len(input)), input...)
	out := make([]byte, 0, len(input))

	for {
		i := 0
		for i < len(in) {
			if i < len(in)-1 && in[i] == opposite(in[i+1]) {
				i += 2
				continue
			}
			out = append(out, in[i])
			i++
		}
		if len(in) == len(out) {
			break
		}
		in, out = out, in[:0]
	}

	return out
}

func part2(input []byte) (unit byte, length int) {
	length = len(input)
	for u := byte('a'); u <= 'z'; u++ {
		in := strip(input, u)
		l := len(part1(in))
		if l < length {
			length = l
			unit = u
		}
	}
	return unit, length
}

func opposite(r byte) byte {
	if r < 'a' {
		return r + 'a' - 'A'
	}
	return r - 'a' + 'A'
}

func strip(input []byte, b byte) []byte {
	out := make([]byte, 0, len(input))
	for i := range input {
		if input[i] == b || input[i] == opposite(b) {
			continue
		}
		out = append(out, input[i])
	}
	return out
}

func read() (input []byte) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("ERROR: open input: %v", err)
	}
	defer f.Close()
	n, err := fmt.Fscanf(f, "%s", &input)
	if n != 1 || err != nil {
		log.Printf("ERROR: scan input: %v", err)
	}
	return input
}
