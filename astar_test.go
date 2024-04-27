package main

import (
	"reflect"
	"testing"
)

func TestGetSuccessors(t *testing.T) {
	g := []int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	grid := NewGrid[int](3, 3)
	i := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			grid.Set(x, y, g[i])
			i++
		}
	}
	cases := []struct {
		name string
		pos  Vec2
		want []Vec2
	}{
		{
			name: "4 cardinal neighbours",
			pos:  Vec2{1, 1},
			want: []Vec2{
				{1, 0},
				{2, 1},
				{1, 2},
				{0, 1},
			},
		},
		{
			name: "bounded right",
			pos:  Vec2{2, 1},
			want: []Vec2{
				{2, 0},
				{2, 2},
				{1, 1},
			},
		},
		{
			name: "bounded left",
			pos:  Vec2{0, 1},
			want: []Vec2{
				{0, 0},
				{1, 1},
				{0, 2},
			},
		},
		{
			name: "bounded top",
			pos:  Vec2{1, 0},
			want: []Vec2{
				{2, 0},
				{1, 1},
				{0, 0},
			},
		},
		{
			name: "bounded bottom",
			pos:  Vec2{1, 2},
			want: []Vec2{
				{1, 1},
				{2, 2},
				{0, 2},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := getSuccessors(c.pos, grid.Width, grid.Height)
			if len(got) != len(c.want) {
				t.Fatalf("len want %d got %d", len(c.want), len(got))
			}
			for i := range c.want {
				if !reflect.DeepEqual(got[i], c.want[i]) {
					t.Errorf("pos %d want %v got %v", i, c.want[i], got[i])
				}
			}
		})
	}
}
func TestPath_NoDiagonal1(t *testing.T) {
	grid := NewGrid[int](5, 5)
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

	pathfinder := NewPathfinder(grid)
	got := pathfinder.Find(Vec2{1, 1}, Vec2{3, 1})
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
		t.Logf(RenderAsString(&grid))
		t.Logf("want: %v", want)
		t.Logf("got: %v", got)
	}
}
func TestPath_NoDiagonal2(t *testing.T) {
	grid := NewGrid[int](8, 4)
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

	pathfinder := NewPathfinder(grid)
	got := pathfinder.Find(Vec2{1, 1}, Vec2{6, 2})
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
		t.Log(RenderAsString(&grid))
		t.Logf("want: %v", want)
		t.Logf("got: %v", got)
	}
}
