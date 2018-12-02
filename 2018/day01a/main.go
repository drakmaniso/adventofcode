package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := read()
	answer := 0
	for _, v := range input {
		answer += v
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
