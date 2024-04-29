package astar

// Vec2 is a 2D vector.
type Vec2 struct {
	X int
	Y int
}

// isHorizAdj returns true if a and b are horizontally adjacent.
func isHorizAdj(a, b Vec2) bool {
	return a.Y-b.Y == 0
}

// isVertAdj returns true if a and b are vertically adjacent.
func isVertAdj(a, b Vec2) bool {
	return a.X-b.X == 0
}

// isDiagAdj returns true if a and b are diagonally adjacent.
func isDiagAdj(a, b Vec2) bool {
	return abs(a.X-b.X) == abs(a.Y-b.Y)
}
