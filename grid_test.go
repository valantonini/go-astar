package astar

import (
	"testing"
)

func TestGrid(t *testing.T) {
	t.Parallel()
	board := NewGrid[int](3, 3)

	i := 0
	for x := range board.Width {
		for y := range board.Height {
			i++
			board.Set(Vec2{x, y}, i)
		}
	}

	want := 0
	for x := range board.Width {
		for y := range board.Height {
			want++
			got := board.Get(Vec2{x, y})
			if got != want {
				t.Errorf("pos %d,%d want %d got %d", x, y, want, got)
			}
		}
	}
}

func TestGridFromSlice(t *testing.T) {
	t.Parallel()
	board := NewGridFromSlice(3, 3, []int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})

	want := 0
	for y := range board.Width {
		for x := range board.Height {
			want++
			got := board.Get(Vec2{x, y})
			if got != want {
				t.Errorf("pos %d,%d want %d got %d", x, y, want, got)
			}
		}
	}
}
