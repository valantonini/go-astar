package astar

// Vec2 is a 2D vector.
type Vec2 struct {
	X int
	Y int
}

// Node is a node in the search space.
type Node struct {
	F      int   // F = G + H
	Pos    Vec2  // Position
	Weight int   // Weight of the node (0 = impassable)
	Parent *Node // Parent node
	G      int   // Cost from start node
	H      int   // Heuristic cost to end node
	Open   bool  // In open list
	Closed bool  // In closed list
}
