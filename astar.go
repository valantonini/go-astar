package main

type Point struct {
	X int
	Y int
}
type Pathfinder struct {
	grid Grid
}

func NewPathfinder(grid Grid) Pathfinder {
	return Pathfinder{
		grid: grid,
	}
}

func (p Pathfinder) Find(x1, y1, x2, y2 int) []Point {
	return []Point{}
}
