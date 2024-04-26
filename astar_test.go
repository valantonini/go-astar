package main

import (
	"testing"
)

func TestPath_NoDiagonal1(t *testing.T) {
	grid := NewGrid(5, 5)
	m := []int{
		1, 1, 1, 1, 1,
		1, 0, 1, 0, 1,
		1, 0, 1, 0, 1,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 1,
	}
	i := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			grid.Set(x, y, m[i])
			i++
		}
	}

	t.Log(RenderAsString(&grid))

	pathfinder := NewPathfinder(grid)
	got := pathfinder.Find(1, 1, 3, 1)
	want := []Vec2{
		{1, 1},
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 3},
		{3, 2},
		{3, 1},
	}
	if len(got) != len(want) {
		t.Logf("got: %v", got)
		t.Fatalf("len want %d got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("pos %d want %v got %v", i, want[i], got[i])
		}
	}
	if t.Failed() {
		t.Logf("want: %v", want)
		t.Logf("got: %v", got)
	}
}
func TestPath_NoDiagonal2(t *testing.T) {
	grid := NewGrid(8, 4)
	m := []int{
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 0, 1, 0, 0, 0, 1, 1,
		1, 0, 0, 0, 1, 0, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}
	i := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			grid.Set(x, y, m[i])
			i++
		}
	}

	t.Log(RenderAsString(&grid))

	pathfinder := NewPathfinder(grid)
	got := pathfinder.Find(1, 1, 6, 2)
	want := []Vec2{
		{1, 1},
		{1, 2},
		{2, 2},
		{3, 2},
		{3, 1},
		{4, 1},
		{5, 1},
		{5, 2},
		{6, 2},
	}
	if len(got) != len(want) {
		t.Logf("got: %v", got)
		t.Fatalf("len want %d got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("pos %d want %v got %v", i, want[i], got[i])
		}
	}
	if t.Failed() {
		t.Logf("want: %v", want)
		t.Logf("got: %v", got)
	}
}
