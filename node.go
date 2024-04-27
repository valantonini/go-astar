package astar

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
