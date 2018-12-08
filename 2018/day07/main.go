package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type step int

func (s step) String() string {
	return string(rune(s) + 'A')
}

func (s step) GoString() string {
	return string(rune(s) + 'A')
}

func main() {
	input := read()
	fmt.Printf("Answer for part 1: %s\n", part1(input))
	fmt.Printf("Answer for part 2: %d\n", part2(input, 5, 60))
}

func part1(input [][2]step) string {
	graph, max := prepare(input)

	todo := [26]bool{}
	for s := step(0); s <= max; s++ {
		todo[s] = true
	}

	answer := strings.Builder{}
	done := false
loop:
	for !done {
		done = true
		for i := range todo {
			if todo[i] {
				done = false
				ready := true
				for j := range graph[i] {
					if todo[graph[i][j]] {
						ready = false
						break
					}
				}
				if ready {
					answer.WriteRune('A' + rune(i))
					todo[i] = false
					continue loop
				}
			}
		}
	}

	return answer.String()
}

func part2(input [][2]step, workers, duration int) int {
	graph, max := prepare(input)
	fmt.Printf("%v\n", graph)

	todo := [26]int{}
	for s := step(0); s <= max; s++ {
		todo[s] = int(s) + 1 + duration
	}

	answer := 0
	done := false
	work := make([]step, workers)
	for i := range work {
		work[i] = -1
	}
loop:
	for !done {
		done = true
		ready := []step{}
	inner:
		for i := range todo {
			if todo[i] > 0 {
				done = false
				for j := range graph[i] {
					if todo[graph[i][j]] > 0 {
						continue inner
					}
				}
				ready = append(ready, step(i))
			}
		}
		if len(ready) > 0 {
			done = false
			fmt.Printf("%v: ", ready)
		innerinner:
			for _, s := range ready {
				for w := range work {
					if work[w] == s {
						continue innerinner
					}
				}
				for w := range work {
					if work[w] == -1 {
						work[w] = s
						break
					}
				}
			}
			answer++
			count := 0
			for w := range work {
				if work[w] != -1 && count < workers {
					fmt.Print(work[w], " ")
					todo[work[w]]--
					if todo[work[w]] <= 0 {
						work[w] = -1
					}
					count++
				}
			}
			fmt.Println()
			continue loop
		}
	}
	return answer
}

func prepare(input [][2]step) (graph [26][]step, max step) {
	max = step(0)
	for _, constraint := range input {
		graph[constraint[1]] = append(graph[constraint[1]], constraint[0])
		if constraint[0] > max {
			max = constraint[0]
		}
		if constraint[1] > max {
			max = constraint[1]
		}
	}

	return graph, max
}

func read() (input [][2]step) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		var before, after step
		n, err := fmt.Sscanf(s.Text(), "Step %c must be finished before step %c can begin.",
			&before, &after)
		if n != 2 || err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}
		input = append(input, [2]step{before - 'A', after - 'A'})
	}
	return input
}
