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
	closed := NewGrid(p.grid.Width, p.grid.Height)

	open.Push(Node{Pos: Vec2{x1, y1}, F: 0})
	for open.Len() > 0 {
		q := open.Pop()
		neighbours := p.grid.Neighbours(q.Pos.X, q.Pos.Y)
		for _, successor := range neighbours {
			successor.Parent = &Vec2{q.Pos.X, q.Pos.Y}
			if successor.Pos.X == x2 && successor.Pos.Y == y2 {
				// found
				panic("found")
			}
			successor.G = q.G + manhattan(q.Pos.X, q.Pos.Y, successor.Pos.X, successor.Pos.Y)
			successor.H = manhattan(successor.Pos.X, successor.Pos.Y, x2, y2)
			successor.F = successor.G + successor.H

			open.Push(successor)
		}
		closed.Set(q.Pos.X, q.Pos.Y, q.F)
	}
	return []Vec2{}
}

func manhattan(x1, y1, x2, y2 int) int {
	return (x1 - x2) + (y1 - y2)
}
