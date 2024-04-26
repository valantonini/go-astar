package main

import (
	"slices"
)

type Pathfinder struct {
	grid Grid
}

func NewPathfinder(grid Grid) Pathfinder {
	return Pathfinder{
		grid: grid,
	}
}

func (p Pathfinder) Find(x1, y1, x2, y2 int) []Vec2 {
	open := NewMinHeap(p.grid.Width, p.grid.Height)
	closed := NewMaxGrid(p.grid.Width, p.grid.Height)

	node := Node{Pos: Vec2{x1, y1}, F: 0, Weight: p.grid.Get(x1, y1)}
	open.Push(node)

	for open.Len() > 0 {
		q := open.Pop()
		neighbours := p.grid.Neighbours(q.Pos.X, q.Pos.Y)
		for _, successor := range neighbours {
			if p.grid.Get(successor.Pos.X, successor.Pos.Y) != 0 {
				// successor blocked
				continue
			}

			successor.Parent = &q

			if successor.Pos.X == x2 && successor.Pos.Y == y2 {
				// found
				path := []Vec2{}
				var n *Node = &successor
				for n != nil {
					path = append(path, n.Pos)
					n = n.Parent
				}
				slices.Reverse(path)
				return path
			}

			successor.G = q.G + manhattan(q.Pos.X, q.Pos.Y, successor.Pos.X, successor.Pos.Y)
			successor.H = manhattan(successor.Pos.X, successor.Pos.Y, x2, y2)
			successor.F = successor.G + successor.H
			successor.Weight = p.grid.Get(successor.Pos.X, successor.Pos.Y)

			if open.FValAt(successor.Pos.X, successor.Pos.Y) < successor.F {
				continue
			}

			if closed.Get(successor.Pos.X, successor.Pos.Y) < successor.F {
				// already found better
				continue
			}
			open.Push(successor)
		}
		closed.Set(q.Pos.X, q.Pos.Y, q.F)
	}
	return []Vec2{}
}

func manhattan(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
