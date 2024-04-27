package astar

import (
	"math"
	"slices"
)

type Pathfinder struct {
	weights Grid[int]
}

func NewPathfinder(weights Grid[int]) Pathfinder {
	return Pathfinder{
		weights: weights,
	}
}

func (p Pathfinder) Find(start, end Vec2) []Vec2 {
	open := NewMinHeap(p.weights.Width, p.weights.Height)
	searchSpace := newSearchSpace(p.weights)

	origin := Node{
		Pos:    start,
		F:      0,
		Weight: p.weights.Get(start),
		Open:   true,
	}
	open.Push(origin)
	searchSpace.Set(origin.Pos, origin)

	for open.Len() > 0 {
		q := open.Pop()
		for _, succ := range getSuccessors(q.Pos, p.weights.Width, p.weights.Height) {
			// cell is not traversable
			if p.weights.Get(succ) == 0 {
				continue
			}

			successor := Node{
				Pos:    Vec2{succ.X, succ.Y},
				Weight: p.weights.Get(succ),
				Parent: &q,
				Open:   true,
			}

			// found
			if successor.Pos == end {
				path := []Vec2{}
				var curr *Node = &successor
				for curr != nil {
					path = append(path, curr.Pos)
					curr = curr.Parent
				}
				slices.Reverse(path)
				return path
			}

			successor.G = q.G + manhattan(q.Pos, successor.Pos)
			successor.H = manhattan(successor.Pos, end)
			successor.F = successor.G + successor.H
			successor.Weight = p.weights.Get(successor.Pos)

			ss := searchSpace.Get(successor.Pos)

			// better node with same position in open list
			if ss.Open && ss.F < successor.F {
				continue
			}

			// better node with same position in closed list
			if ss.Closed && ss.F < successor.F {
				continue
			}

			searchSpace.Set(successor.Pos, successor)
			open.Push(successor)
		}

		s := searchSpace.Get(q.Pos)
		s.Closed = true
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

var neighbours = [4][2]int{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
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

func newSearchSpace(weights Grid[int]) Grid[Node] {
	g := NewGrid[Node](weights.Width, weights.Height)
	for x := range weights.Width {
		for y := range weights.Height {
			c := Node{
				Pos:    Vec2{x, y},
				Weight: weights.Get(Vec2{x, y}),
				F:      math.MaxInt,
			}
			g.Set(Vec2{x, y}, c)
		}
	}
	return g
}
