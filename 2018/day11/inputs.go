package main

var input = 5034

var subexamples = []struct {
	x, y int
	serial int
	power int
}{
	{3, 5, 8, 4},
	{122, 79, 57, -5},
	{217, 196, 39, 0},
	{101, 153, 71, 4},
}

var examples = []struct {
	serial int
	x,y int
	power int
}{
	{18, 33, 45, 29},
	{42, 21, 61, 30},
}

var examples2 = []struct {
	serial int
	x,y int
	size int
	power int
}{
	{18, 90,269,16,113},
	{42,232,251,12,119},
}
