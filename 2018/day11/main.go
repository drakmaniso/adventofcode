package main

import "fmt"

func main() {
	x, y, p := part1(5034)
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

func part1(serial int) (clusterX, clusterY, clusterPower int) {
	const size = 3
	for x := 1; x <= 300-size; x++ {
		for y := 1; y <= 300-size; y++ {
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
	powers := [301][301]int{}
	for s := 1; s <= 300; s++ {
		for x := 1; x <= 300-s; x++ {
			for y := 1; y <= 300-s; y++ {
				powers[x][y] += power(x+s-1,y+s-1, serial)
				for i := 0; i < s-1; i++ {
					powers[x][y] += power(x+i,y+s-1, serial)
					powers[x][y] += power(x+s-1,y+i, serial)
				}
				if powers[x][y] > clusterPower {
					clusterX, clusterY = x, y
					clusterSize = s
					clusterPower = powers[x][y]
				}
			}
		}
	}
	return clusterX, clusterY, clusterSize, clusterPower
}
