package main

import (
	"fmt"
	"strings"
)

type node struct {
	Marble         int
	previous, next *node
}

func (n *node) move(delta int) *node {
	current := n
	if delta > 0 {
		for i := 0; i < delta; i++ {
			current = current.next
		}
		return current
	}
	for i := 0; i < -delta; i++ {
		current = current.previous
	}
	return current
}

func (n *node) insert(marble int) *node {
	newnode := node{
		Marble:   marble,
		previous: n.previous,
		next:     n,
	}
	n.previous.next = &newnode
	n.previous = &newnode
	return &newnode
}

func (n *node) remove() *node {
	n.previous.next = n.next
	n.next.previous = n.previous
	return n.next
}

func (n *node) String() string {
	s := strings.Builder{}
	current := n
	for ; current.Marble != 0; current = current.next {
	}
	for {
		if current == n {
			s.WriteString(fmt.Sprintf("(% 2d)", current.Marble))
		} else {
			s.WriteString(fmt.Sprintf(" % 2d ", current.Marble))
		}
		current = current.next
		if current.Marble == 0 {
			break
		}
	}
	return s.String()
}

func main() {
	playerCount, lastMarble := read(input)
	fmt.Printf("Answer for part 1: %d\n", part1(playerCount, lastMarble))
	fmt.Printf("Answer for part 2: %d\n", part1(playerCount, 100*lastMarble))
}

func part1(	playerCount, lastMarble int) (highscore int) {
	current := &node{Marble: 0}
	current.next, current.previous = current, current
	player := 1
	scores := make([]int, playerCount)

	for m := 1; m <= lastMarble; m++ {
		if m%23 != 0 {
			current = current.move(2)
			current = current.insert(m)
		} else {
			scores[player] += m
			current = current.move(-7)
			scores[player] += current.Marble
			current = current.remove()
		}
		player = (player + 1) % playerCount
	}

	winner := 0
	for i := range scores {
		if scores[i] > scores[winner] {
			winner = i
		}
	}

	return scores[winner]
}

func read(input string) (playerCount, lastMarble int) {
	fmt.Sscanf(input, "%d players; last marble is worth %d points",
		&playerCount, &lastMarble)
	return playerCount, lastMarble
}
