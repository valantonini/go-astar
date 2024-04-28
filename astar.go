package astar

import (
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
	open.push(heapNode{pos: start})
	searchSpace.Set(start, origin)

	for open.len() > 0 {
		qPos := open.pop().pos
		q := searchSpace.Get(qPos)
		for _, succPos := range getSuccessors(qPos, searchSpace.Width, searchSpace.Height) {
			successor := searchSpace.Get(succPos)

			// cell is not traversable
			if successor.weight == 0 {
				continue
			}

			successor.parent = &q
			successor.g = q.g + manhattan(qPos, successor.pos)
			successor.h = manhattan(succPos, end)
			successor.f = successor.g + successor.h
			successor.open = true

			// found
			if succPos == end {
				path := []Vec2{}
				var curr *node = &successor
				for curr != nil {
					path = append(path, curr.pos)
					curr = curr.parent
				}
				slices.Reverse(path)
				return path
			}

			// current successor
			existingSuccessor := searchSpace.Get(succPos)

			// better successor with same position in open list
			if existingSuccessor.open && existingSuccessor.f < successor.f {
				continue
			}

			// better successor with same position in closed list
			if existingSuccessor.closed && existingSuccessor.f < successor.f {
				continue
			}

			searchSpace.Set(succPos, successor)
			open.push(heapNode{pos: succPos, f: successor.f})
		}
		q.closed = true
		searchSpace.Set(qPos, q)
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
			}
			grid.Set(node.pos, node)
		}
	}
	return grid
}
