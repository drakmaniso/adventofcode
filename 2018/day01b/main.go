package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := read()
	previous := map[int]bool{0: true}
	answer := 0
loop:
	for {
		for _, v := range input {
			answer += v
			if previous[answer] {
				break loop
			}
			previous[answer] = true
		}
	}
	fmt.Printf("Answer: %d\n", answer)
}

func read() (input []int) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Printf("ERROR: read input: %v\n", err)
			continue
		}
		input = append(input, v)
	}
	return input
}
