package main

import (
	"reflect"
	"testing"
)

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

func TestGrid_Neighbours(t *testing.T) {
	g := []byte{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	grid := NewGrid(3, 3)
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
		want []Node
	}{
		{
			name: "4 cardinal neighbours",
			pos:  Vec2{1, 1},
			want: []Node{
				{Pos: Vec2{1, 0}, Weight: 2},
				{Pos: Vec2{2, 1}, Weight: 6},
				{Pos: Vec2{1, 2}, Weight: 8},
				{Pos: Vec2{0, 1}, Weight: 4},
			},
		},
		{
			name: "bounded right",
			pos:  Vec2{2, 1},
			want: []Node{
				{Pos: Vec2{2, 0}, Weight: 3},
				{Pos: Vec2{2, 2}, Weight: 9},
				{Pos: Vec2{1, 1}, Weight: 5},
			},
		},
		{
			name: "bounded left",
			pos:  Vec2{0, 1},
			want: []Node{
				{Pos: Vec2{0, 0}, Weight: 1},
				{Pos: Vec2{1, 1}, Weight: 5},
				{Pos: Vec2{0, 2}, Weight: 7},
			},
		},
		{
			name: "bounded top",
			pos:  Vec2{1, 0},
			want: []Node{
				{Pos: Vec2{2, 0}, Weight: 3},
				{Pos: Vec2{1, 1}, Weight: 5},
				{Pos: Vec2{0, 0}, Weight: 1},
			},
		},
		{
			name: "bounded bottom",
			pos:  Vec2{1, 2},
			want: []Node{
				{Pos: Vec2{1, 1}, Weight: 5},
				{Pos: Vec2{2, 2}, Weight: 9},
				{Pos: Vec2{0, 2}, Weight: 7},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := grid.Neighbours(c.pos.X, c.pos.Y)
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
