package astar

type heuristic int

const (
	// Manhattan heuristic.
	man heuristic = iota
	// Euclidean heuristic.
	dd
)

// Option is a functional option for the pathfinder.
type Option func(o *Options)

// Options contains the options for the pathfinder.
type Options struct {
	diagonals             bool
	heuristic             heuristic
	punishChangeDirection bool
}

// WithDiagonals enables diagonal movement in the search space.
func WithDiagonals() Option {
	return Option(func(o *Options) {
		o.diagonals = true
	})
}

// PunishChangeDirection punishes changing direction when calculating G.
func PunishChangeDirection() Option {
	return Option(func(o *Options) {
		o.punishChangeDirection = true
	})
}
