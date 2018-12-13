package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	tracks, carts := read(input)
	x, y := part1(tracks, carts)
	fmt.Printf("Answer for part 1: %d, %d\n", x, y)

	tracks, carts = read(input)
	x, y = part2(tracks, carts)
	fmt.Printf("Answer for part 2: %d, %d\n", x, y)
}

func part1(tracks [][]byte, carts []Cart) (x, y int) {
	for {
		for i := range carts {
			step(tracks, carts, i)
			c, x, y, _ := collides(carts)
			// PrintAll(tracks, carts)
			if c {
				return x, y
			}
			// _, _ = fmt.Scanf("%s")
		}
	}
}

func part2(tracks [][]byte, carts []Cart) (x, y int) {
	for {
		remove := make([]bool, len(carts))
		for i := range carts {
			step(tracks, carts, i)
			_, _, _, rm := collides(carts)
			for i := range remove {
				if rm[i] {
					remove[i] = true
				}
			}
			// PrintAll(tracks, carts)
			// _, _ = fmt.Scanf("%s")
		}
		remaining := []Cart{}
		for i, c := range carts {
			if !remove[i] {
				remaining = append(remaining, c)
			}
		}
		carts = remaining
		if len(carts) == 1 {
			return carts[0].X, carts[0].Y
		}
	}
}

func step(tracks [][]byte, carts []Cart, cart int) {
	i := cart
	switch tracks[carts[i].Y][carts[i].X] {
	case '+':
		switch carts[i].State {
		case 0:
			carts[i].Direction = (carts[i].Direction + 1) % 4
		case 1:
			// NOP
		case 2:
			carts[i].Direction = (carts[i].Direction - 1) % 4
		default:
			panic("impossible state")
		}
		carts[i].State = (carts[i].State + 1) % 3
	case '\\':
		switch carts[i].Direction {
		case Left, Right:
			carts[i].Direction = (carts[i].Direction - 1) % 4
		case Up, Down:
			carts[i].Direction = (carts[i].Direction + 1) % 4
		default:
			panic("impossible direction in turn")
		}
	case '/':
		switch carts[i].Direction {
		case Left, Right:
			carts[i].Direction = (carts[i].Direction + 1) % 4
		case Up, Down:
			carts[i].Direction = (carts[i].Direction - 1) % 4
		default:
			panic("impossible direction in turn")
		}
	case '|', '-':
		// NOP
	default:
		println(carts[i].Y, carts[i].X, tracks[carts[i].Y][carts[i].X])
		panic("impossible track under cart " + string(tracks[carts[i].Y][carts[i].X]))
	}
	switch carts[i].Direction {
	case Up:
		carts[i].Y--
	case Down:
		carts[i].Y++
	case Left:
		carts[i].X--
	case Right:
		carts[i].X++
	default:
		println(carts[i].Direction)
		panic("impossible direction")
	}
}

func collides(carts []Cart) (collision bool, x, y int, remove []bool) {
	remove = make([]bool, len(carts))
	for i := range carts {
		for j := i + 1; j < len(carts); j++ {
			if carts[i].X == carts[j].X && carts[i].Y == carts[j].Y {
				collision, x, y = true, carts[i].X, carts[i].Y
				remove[i] = true
				remove[j] = true
			}
		}
	}
	return collision, x, y, remove
}

func read(input string) (tracks [][]byte, carts []Cart) {
	s := bufio.NewScanner(strings.NewReader(input))
	x, y := 0, 0
	for s.Scan() {
		x = 0
		l := make([]byte, len(s.Bytes()))
		for i, b := range s.Bytes() {
			switch b {
			case '^':
				l[i] = '|'
				carts = append(carts, Cart{X: x, Y: y, Direction: Up})
			case 'v':
				l[i] = '|'
				carts = append(carts, Cart{X: x, Y: y, Direction: Down})
			case '<':
				l[i] = '-'
				carts = append(carts, Cart{X: x, Y: y, Direction: Left})
			case '>':
				l[i] = '-'
				carts = append(carts, Cart{X: x, Y: y, Direction: Right})
			case '|', '-', '+', '/', '\\', ' ':
				l[i] = b
			default:
				log.Printf("unrecognized track part %c", b)
				l[i] = '?'
			}
			x++
		}
		tracks = append(tracks, l)
		y++
	}
	return tracks, carts
}

type Cart struct {
	X, Y      int
	Direction uint8
	State     int
}

const (
	Up = iota
	Left
	Down
	Right
)

func PrintTracks(tracks [][]byte) {
	for _, l := range tracks {
		fmt.Println(string(l))
	}
}

func PrintAll(tracks [][]byte, carts []Cart) {
	s := strings.Builder{}
	s.WriteString("\033[2J")
	for y := range tracks {
	loop:
		for x := range tracks[y] {
			for _, c := range carts {
				if c.Y == y && c.X == x {
					s.WriteByte("^<v>"[c.Direction])
					continue loop
				}
			}
			s.WriteByte(tracks[y][x])
		}
		s.WriteByte('\n')
	}
	fmt.Println(s.String())
}
