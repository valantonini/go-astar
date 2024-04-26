package main

import "testing"

func TestGrid(t *testing.T) {
	board := NewGrid(3, 3)

	i := 0
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			i++
			board.Set(x, y, byte(i))
		}
	}

	want := 0
	for x := 0; x < board.Width; x++ {
		for y := 0; y < board.Height; y++ {
			want++
			got := board.Get(x, y)
			if got != byte(want) {
				t.Errorf("pos %d,%d want %d got %d", x, y, want, got)
			}
		}
	}
}
