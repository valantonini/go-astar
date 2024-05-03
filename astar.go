package astar

import (
	"log"
	"math"
	"slices"
)

// cardinalSuccessors are the offsets for the 4 cardinal directions.
var cardinalSuccessors = []Vec2{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

// diagonalSuccessors are the offsets for the 8 cardinal and diagonal
// directions.
var diagonalSuccessors = []Vec2{
	{0, -1},  // Up
	{1, -1},  // Up-Right
	{1, 0},   // Right
	{1, 1},   // Right-Down
	{0, 1},   // Down
	{-1, 1},  // Down-Left
	{-1, 0},  // Left
	{-1, -1}, // Up-Left
}

// heuristicFunc is a function that calculates the distance between two vectors.
type heuristicFunc func(Vec2, Vec2) int

// getSuccessorsFunc is a function that returns the successors of a vector for
// a given search space.
type getSuccessorsFunc func(v Vec2) []Vec2

// node is a node in the search space.
type node struct {
	pos    Vec2  // Position
	parent *node // Parent node
	g      int   // Cost from start node
	h      int   // Heuristic cost to end node
	f      int   // F = G + H
	weight int   // Weight of the node (0 = impassable)
	open   bool  // In open list
	closed bool  // In closed list
}

// Pathfinder is a simple A* pathfinding algorithm implementation.
type Pathfinder struct {
	weights       Grid[int]
	options       option
	heuristic     heuristicFunc
	getSuccessors getSuccessorsFunc
}

// NewPathfinder creates a new Pathfinder with the given weights and options.
// The weights are used to determine the cost of traversing a cell. A weight of
// 0 means the cell is not traversable. A weight of 1 or higher means the cell
// is traversable.
func NewPathfinder(weights Grid[int], opts ...Option) Pathfinder {
	opt := option{
		heuristic: man,
	}

	for _, optFunc := range opts {
		optFunc(&opt)
	}

	pf := Pathfinder{
		weights:   weights,
		options:   opt,
		heuristic: manhattan,
		getSuccessors: func(v Vec2) []Vec2 {
			return getSuccessors(v, weights.Width, weights.Height, cardinalSuccessors)
		},
	}

	switch opt.heuristic {
	case man:
		pf.heuristic = manhattan
	case dd:
		pf.heuristic = diagonalDistance
	case euc:
		pf.heuristic = euclideanDistance
	default:
		log.Printf("unknown heuristic: %d", opt.heuristic)
	}

	if opt.diagonals {
		if opt.heuristic == man {
			pf.heuristic = diagonalDistance
		}
		pf.getSuccessors = func(v Vec2) []Vec2 {
			return getSuccessors(v, weights.Width, weights.Height, diagonalSuccessors)
		}
	}

	return pf
}

// Find returns a path from start to end. If no path is found, an empty slice
// is returned.
func (p Pathfinder) Find(startPos, endPos Vec2) []Vec2 {
	searchSpace := newSearchSpace(p.weights)                  // tracks the open, closed and f values of each node
	open := newMinHeap(searchSpace.Width, searchSpace.Height) // prioritised queue of f

	start := searchSpace.Get(startPos)
	start.f = 0
	start.open = true

	open.push(heapNode{pos: startPos, f: start.f})
	searchSpace.Set(startPos, start)

	for open.len() > 0 {
		qPos := open.pop().pos
		q := searchSpace.Get(qPos)
		for _, succPos := range p.getSuccessors(qPos) {
			successor := searchSpace.Get(succPos)

			// not traversable
			if successor.weight == 0 {
				continue
			}

			successor.parent = &q
			successor.g = q.g + p.heuristic(qPos, succPos)
			if p.options.punishChangeDirection {
				successor.g += punishChangeDirection(q, succPos, endPos)
			}
			successor.h = p.heuristic(succPos, endPos)
			successor.f = successor.g + successor.h
			successor.open = true

			// found
			if succPos == endPos {
				path := []Vec2{}
				curr := &successor
				for curr != nil {
					path = append(path, curr.pos)
					curr = curr.parent
				}
				slices.Reverse(path)
				return path
			}

			// check if more optimal path to successor was already encountered
			existingSuccessor := searchSpace.Get(succPos)
			if existingSuccessor.open && existingSuccessor.f < successor.f {
				continue
			}
			if existingSuccessor.closed && existingSuccessor.f < successor.f {
				continue
			}

			searchSpace.Set(succPos, successor)
			open.push(heapNode{pos: succPos, f: successor.f})
		}
		q.closed = true
		searchSpace.Set(qPos, q)
	}

	// not found
	return []Vec2{}
}

// punishChangeDirection returns a punishment for changing direction that can be applied to g.
func punishChangeDirection(q node, successor, end Vec2) int {
	if q.parent == nil {
		return 0
	}
	punishment := abs(successor.X-end.X) + abs(successor.Y-end.Y)

	if !isHorizAdj(q.pos, successor) {
		if isHorizAdj(q.pos, q.parent.pos) {
			return punishment
		}
	}

	if !isVertAdj(q.pos, successor) {
		if isVertAdj(q.pos, q.parent.pos) {
			return punishment
		}
	}

	if !isDiagAdj(q.pos, successor) {
		if isDiagAdj(q.pos, q.parent.pos) {
			return punishment
		}
	}

	return 0
}

// manhattan calculates the Manhattan distance between two vectors by summing
// the absolute values of the differences of their components. It does not
// support diagonal movement.
func manhattan(v1, v2 Vec2) int {
	return abs(v1.X-v2.X) + abs(v1.Y-v2.Y)
}

// diagonalDistance calculates the diagonal distance between two vectors.
func diagonalDistance(v1, v2 Vec2) int {
	dx := abs(v1.X - v2.X)
	dy := abs(v1.Y - v2.Y)
	diagonal := float64(min(dx, dy))
	straight := float64(dx + dy)
	return int(straight + (float64(math.Sqrt2)-2)*diagonal)
}

// euclideanDistance calculates the Euclidean distance between two vectors.
func euclideanDistance(v1, v2 Vec2) int {
	// pythagorean
	h := math.Sqrt(
		math.Pow(float64(v1.X-v2.X), 2) + math.Pow(float64(v1.Y-v2.Y), 2),
	)
	return int(h)
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// min returns the minimum of x1 and x2.
func min(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}

// getSuccessors returns the successors of a vector. If a successor is outside
// of the grid, it is not included.
func getSuccessors(vec Vec2, width, height int, offsets []Vec2) []Vec2 {
	results := make([]Vec2, 0, len(offsets))
	for _, n := range offsets {
		x := vec.X + n.X
		y := vec.Y + n.Y

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
