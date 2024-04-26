package main

type Pathfinder struct {
	grid Grid
}

func NewPathfinder(grid Grid) Pathfinder {
	return Pathfinder{
		grid: grid,
	}
}

func (p Pathfinder) Find(x1, y1, x2, y2 int) []Vec2 {
	open := &MinHeap{}
	open.Push(Node{Pos: Vec2{x1, y1}, F: 0})
	for open.Len() > 0 {
		q := open.Pop()
		_ = q
	}
	return []Vec2{}
}
