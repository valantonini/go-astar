package astar

import (
	"math"
	"slices"
)

// Pathfinder is a simple A* pathfinding algorithm implementation.
type Pathfinder struct {
	weights Grid[int]
}

// NewPathfinder creates a new Pathfinder with the given weights. The weights
// are used to determine the cost of traversing a cell. A weight of 0 means the
// cell is not traversable. A weight of 1 or higher means the cell is
// traversable.
func NewPathfinder(weights Grid[int]) Pathfinder {
	return Pathfinder{
		weights: weights,
	}
}

// Find returns a path from start to end. If no path is found, an empty slice.
func (p Pathfinder) Find(start, end Vec2) []Vec2 {
	open := newMinHeap(p.weights.Width, p.weights.Height)
	searchSpace := newSearchSpace(p.weights) // tracks the open, closed and f values of each node

	origin := node{
		pos:    start,
		f:      0,
		weight: p.weights.Get(start),
		open:   true,
	}
	open.push(origin)
	searchSpace.Set(origin.pos, origin)

	for open.len() > 0 {
		q := open.pop()
		for _, succ := range getSuccessors(q.pos, p.weights.Width, p.weights.Height) {
			// cell is not traversable
			if p.weights.Get(succ) == 0 {
				continue
			}

			successor := node{
				pos:    Vec2{succ.X, succ.Y},
				weight: p.weights.Get(succ),
				parent: &q,
				open:   true,
			}

			// found
			if successor.pos == end {
				path := []Vec2{}
				var curr *node = &successor
				for curr != nil {
					path = append(path, curr.pos)
					curr = curr.parent
				}
				slices.Reverse(path)
				return path
			}

			successor.g = q.g + manhattan(q.pos, successor.pos)
			successor.h = manhattan(successor.pos, end)
			successor.f = successor.g + successor.h
			successor.weight = p.weights.Get(successor.pos)

			ss := searchSpace.Get(successor.pos)

			// better node with same position in open list
			if ss.open && ss.f < successor.f {
				continue
			}

			// better node with same position in closed list
			if ss.closed && ss.f < successor.f {
				continue
			}

			searchSpace.Set(successor.pos, successor)
			open.push(successor)
		}

		s := searchSpace.Get(q.pos)
		s.closed = true
		searchSpace.Set(s.pos, s)
	}
	return []Vec2{}
}

// manhattan calculates the Manhattan distance between two vectors by summing
// the absolute values of the differences of their components. It does not
// support diagonal movement.
func manhattan(v1, v2 Vec2) int {
	return abs(v1.X-v2.X) + abs(v1.Y-v2.Y)
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var successors = [4][2]int{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

// getSuccessors returns the successors of a vector. If a successor is outside
// of the grid, it is not included.
func getSuccessors(vec Vec2, width, height int) []Vec2 {
	results := make([]Vec2, 0, len(successors))
	for _, n := range successors {
		x := vec.X + n[0]
		y := vec.Y + n[1]

		if x < 0 || x >= width || y < 0 || y >= height {
			continue
		}
		results = append(results, Vec2{x, y})
	}
	return results
}

// newSearchSpace creates a new search space from the given weights.
func newSearchSpace(weights Grid[int]) Grid[node] {
	grid := NewGrid[node](weights.Width, weights.Height)
	for x := range weights.Width {
		for y := range weights.Height {
			node := node{
				pos:    Vec2{x, y},
				weight: weights.Get(Vec2{x, y}),
				f:      math.MaxInt,
			}
			grid.Set(node.pos, node)
		}
	}
	return grid
}
