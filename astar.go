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

func (p Pathfinder) Find(start, end Vec2) []Vec2 {
	open := NewMinHeap(p.grid.Width, p.grid.Height)
	closed := NewMaxGrid(p.grid.Width, p.grid.Height)

	open.Push(Node{
		Pos:    start,
		F:      0,
		Weight: p.grid.Get(start),
	})

	for open.Len() > 0 {
		q := open.Pop()
		for _, succ := range getSuccessors(q.Pos, p.grid.Width, p.grid.Height) {
			// cell is not open
			if p.grid.Get(succ) != 0 {
				continue
			}

			successor := Node{
				Pos:    Vec2{succ.X, succ.Y},
				Weight: p.grid.Get(succ),
				Parent: &q,
			}

			// found
			if successor.Pos == end {
				path := []Vec2{}
				var n *Node = &successor
				for n != nil {
					path = append(path, n.Pos)
					n = n.Parent
				}
				slices.Reverse(path)
				return path
			}

			successor.G = q.G + manhattan(q.Pos, successor.Pos)
			successor.H = manhattan(successor.Pos, end)
			successor.F = successor.G + successor.H
			successor.Weight = p.grid.Get(successor.Pos)

			// already found better
			if closed.Get(successor.Pos) < successor.F {
				continue
			}

			open.Push(successor)
		}
		closed.Set(q.Pos.X, q.Pos.Y, q.F)
	}
	return []Vec2{}
}

func manhattan(v1, v2 Vec2) int {
	return abs(v1.X-v2.X) + abs(v1.Y-v2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getSuccessors(vec Vec2, width, height int) []Vec2 {
	results := make([]Vec2, 0, len(neighbours))
	for _, n := range neighbours {
		x := vec.X + n[0]
		y := vec.Y + n[1]

		if x < 0 || x >= width || y < 0 || y >= height {
			continue
		}
		n := Vec2{x, y}
		results = append(results, n)
	}
	return results
}
