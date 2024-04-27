package astar

// node is a node in the search space.
type node struct {
	f      int   // F = G + H
	pos    Vec2  // Position
	weight int   // Weight of the node (0 = impassable)
	parent *node // Parent node
	g      int   // Cost from start node
	h      int   // Heuristic cost to end node
	open   bool  // In open list
	closed bool  // In closed list
}
