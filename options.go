package astar

type heuristic int

const (
	// Manhattan heuristic.
	man heuristic = iota
	// Diagonal distance heuristic.
	dd
	// Euclidean distance heuristic
	euc
)

// Option is a functional option for the pathfinder.
type Option func(o *option)

// option contains the options for the pathfinder.
type option struct {
	diagonals             bool
	heuristic             heuristic
	punishChangeDirection bool
}

// WithDiagonals enables diagonal movement in the search space.
func WithDiagonals() Option {
	return Option(func(o *option) {
		o.diagonals = true
	})
}

// PunishChangeDirection punishes changing direction when calculating G.
func PunishChangeDirection() Option {
	return Option(func(o *option) {
		o.punishChangeDirection = true
	})
}

// EuclideanDistance sets the heuristic to Euclidean distance.
func EuclideanDistance() Option {
	return Option(func(o *option) {
		o.heuristic = euc
	})
}
