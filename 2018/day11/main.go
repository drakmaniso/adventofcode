package main

import "fmt"

func main() {
	x, y, p := part1(5034, 3)
	fmt.Printf("Answer for part 1: cell %d,%d (power %d)\n", x, y, p)
	var s int
	x, y, s, p = part2(5034)
	fmt.Printf("Answer for part 1: cell %d,%d,%d (power %d)\n", x, y, s, p)
}

func power(x, y int, serial int) int {
	rackID := x + 10
	p := rackID * y
	p += serial
	p *= rackID
	p = (p / 100) % 10
	p -= 5
	return p
}

func part1(serial int, size int) (clusterX, clusterY, clusterPower int) {
	for x := 0; x < 300-size; x++ {
		for y := 0; y < 300-size; y++ {
			p := 0
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					p += power(x+i, y+j, serial)
				}
			}
			if p > clusterPower {
				clusterX, clusterY = x, y
				clusterPower = p
			}
		}
	}
	return clusterX, clusterY, clusterPower
}
func part2(serial int) (clusterX, clusterY, clusterSize, clusterPower int) {
	for s := 1; s < 300; s++ {
		x, y, p := part1(serial, s)
		if p > clusterPower {
			clusterX, clusterY = x, y
			clusterSize = s
			clusterPower = p
		}
	}
	return clusterX, clusterY, clusterSize, clusterPower
}
