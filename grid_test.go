package main

import (
	"testing"
)

func TestGrid(t *testing.T) {
	board := NewGrid[int](3, 3)

	i := 0
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			i++
			board.Set(x, y, i)
		}
	}

	want := 0
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			want++
			got := board.Get(Vec2{x, y})
			if got != want {
				t.Errorf("pos %d,%d want %d got %d", x, y, want, got)
			}
		}
	}
}
