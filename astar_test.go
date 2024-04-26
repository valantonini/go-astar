package main

import (
	"testing"
)

func TestPath(t *testing.T) {
	grid := NewGrid(5, 5)
	m := []byte{
		1, 1, 1, 1, 1,
		1, 0, 1, 0, 1,
		1, 0, 1, 0, 1,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 1,
	}
	i := 0
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			grid.Set(x, y, m[i])
			i++
		}
	}

	pathfinder := NewPathfinder(grid)
	got := pathfinder.Find(1, 1, 3, 3)
	want := []Point{
		{1, 1},
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 3},
		{3, 2},
		{3, 1},
	}
	if len(got) != len(want) {
		t.Fatalf("len want %d got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("pos %d want %v got %v", i, want[i], got[i])
		}
	}
}
