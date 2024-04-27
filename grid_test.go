package astar

import (
	"testing"
)

func TestGrid(t *testing.T) {
	board := NewGrid[int](3, 3)

	i := 0
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			i++
			board.Set(Vec2{x, y}, i)
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

func TestGridFromSlice(t *testing.T) {
	board := NewGridFromSlice(3, 3, []int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})

	want := 0
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			want++
			got := board.Get(Vec2{x, y})
			if got != want {
				t.Errorf("pos %d,%d want %d got %d", x, y, want, got)
			}
		}
	}
}
