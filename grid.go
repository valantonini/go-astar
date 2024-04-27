package main

type Grid[T any] struct {
	Width  int
	Height int
	inner  []T
}

func NewGrid[T any](width, height int) Grid[T] {
	grid := Grid[T]{
		Width:  width,
		Height: height,
		inner:  make([]T, width*height),
	}
	return grid
}

func (b *Grid[T]) Set(v Vec2, val T) {
	idx := v.Y*b.Width + v.X
	b.inner[idx] = val
}

func (b *Grid[T]) Get(v Vec2) T {
	idx := v.Y*b.Width + v.X
	return b.inner[idx]
}
