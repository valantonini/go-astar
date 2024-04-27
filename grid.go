package astar

// Grid is a 2D grid of any type.
type Grid[T any] struct {
	Width  int
	Height int
	inner  []T
}

// NewGrid creates a new grid with the given width and height.
func NewGrid[T any](width, height int) Grid[T] {
	grid := Grid[T]{
		Width:  width,
		Height: height,
		inner:  make([]T, width*height),
	}
	return grid
}

// Set sets the value at the given position.
func (b *Grid[T]) Set(v Vec2, val T) {
	idx := v.Y*b.Width + v.X
	b.inner[idx] = val
}

// Get returns the value at the given position.
func (b *Grid[T]) Get(v Vec2) T {
	idx := v.Y*b.Width + v.X
	return b.inner[idx]
}
