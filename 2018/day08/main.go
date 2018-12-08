package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func main() {
	in := read(strings.NewReader(input))
	_, answer1 := part1(in, 0)
	fmt.Printf("Answer for part 1: %d\n", answer1)
	fmt.Printf("Answer for part 2: %d\n", part2(in))
}

func part1(input []int, start int) (next int, sum int) {
	if start >= len(input) {
		log.Printf("out of range: %d\n", start)
		return -1, -1
	}
	next = start + 2
	for i := 0; i < input[start]; i++ {
		s := 0
		next, s = part1(input, next)
		sum += s
	}
	for i := 0; i < input[start+1]; i++ {
		sum += input[next]
		next++
	}
	return next, sum
}

type node struct {
	children []int
	metadata []int
}

func part2(input []int) int {
	_, nodes := parseNode(input, 0, []node{})
	return license(nodes, 0)
}

func parseNode(input []int, start int, nodes []node) (int, []node) {
	if start >= len(input) {
		log.Printf("out of range: %d\n", start)
		return -1, nil
	}
	next := start + 2
	nodes = append(nodes, node{})
	n := len(nodes) - 1
	for i := 0; i < input[start]; i++ {
		nodes[n].children = append(nodes[n].children, len(nodes))
		next, nodes = parseNode(input, next, nodes)
	}
	for i := 0; i < input[start+1]; i++ {
		nodes[n].metadata = append(nodes[n].metadata, input[next])
		next++
	}
	return next, nodes
}

func license(nodes []node, n int) int {
	sum := 0
	if len(nodes[n].children) == 0 {
		for _, v := range nodes[n].metadata {
			sum += v
		}
	} else {
		for _, v := range nodes[n].metadata {
			child := v-1
			if child < len(nodes[n].children) {
				sum += license(nodes, nodes[n].children[child])
			}
		}
	}
	return sum
}

func read(r io.Reader) (input []int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Printf("ERROR: read: unable to parse %#v", s.Text())
			continue
		}
		input = append(input, v)
	}
	return input
}
