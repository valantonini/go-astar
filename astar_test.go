package astar

import (
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestGetSuccessors_Cardinal(t *testing.T) {
	w := []int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	grid := NewGridFromSlice(3, 3, w)

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
			got := getSuccessors(c.pos, grid.Width, grid.Height, cardinalSuccessors)
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

func TestGetSuccessors_Diagonal(t *testing.T) {
	w := []int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	grid := NewGridFromSlice(3, 3, w)
	got := getSuccessors(Vec2{1, 1}, grid.Width, grid.Height, diagonalSuccessors)

	want := []Vec2{
		{1, 0},
		{2, 0},
		{2, 1},
		{2, 2},
		{1, 2},
		{0, 2},
		{0, 1},
		{0, 0},
	}
	if len(got) != len(want) {
		t.Fatalf("len want %d got %d", len(want), len(got))
	}
	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Errorf("pos %d want %v got %v", i, want[i], got[i])
		}
	}
}

func TestPath_NoDiagonal1(t *testing.T) {
	w := []int{
		0, 0, 0, 0, 0,
		0, 1, 0, 1, 0,
		0, 1, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(5, 5, w)

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
		t.Logf(renderAsString(&grid))
		t.Logf("want: %v", want)
		t.Logf("got: %v", got)
	}
}

func TestPath_NoPath(t *testing.T) {
	w := []int{
		0, 0, 0, 0, 0,
		0, 1, 0, 1, 0,
		0, 0, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(5, 5, w)

	pathfinder := NewPathfinder(grid)
	got := pathfinder.Find(Vec2{1, 1}, Vec2{3, 1})

	if len(got) != 0 {
		t.Logf(renderAsString(&grid))
		t.Logf("got: %v", got)
		t.Fatalf("len want %d got %d", 0, len(got))
	}
}

func TestPath_NoDiagonal2(t *testing.T) {
	w := []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 1, 1, 1, 0, 0,
		0, 1, 1, 1, 0, 1, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(8, 4, w)

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
		t.Log(renderAsString(&grid))
		t.Logf("want: %v", want)
		t.Logf("got: %v", got)
	}
}

func renderAsString(grid *Grid[int]) string {
	sb := &strings.Builder{}
	sb.WriteString("\n")
	for y := range grid.Height {
		for x := range grid.Width {
			val := grid.Get(Vec2{x, y})
			switch val {
			case 0:
				sb.WriteRune('\u2588') // block █
			default:
				sb.WriteString(strconv.Itoa(val))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func renderWithPathAsString(grid *Grid[int], path []Vec2) string {
	sb := &strings.Builder{}
	sb.WriteString("\n")
	for y := range grid.Height {
		for x := range grid.Width {
			val := grid.Get(Vec2{x, y})

			if slices.Contains(path, Vec2{x, y}) {
				sb.WriteRune('\u25e6')
				continue
			}

			switch val {
			case 0:
				sb.WriteRune('\u2588') // block █
			default:
				sb.WriteString(strconv.Itoa(val))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func TestPath_Diagonal1(t *testing.T) {
	w := []int{
		0, 0, 0, 0, 0,
		0, 1, 0, 1, 0,
		0, 1, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(5, 5, w)

	pathfinder := NewDiagonalPathfinder(grid)
	got := pathfinder.Find(Vec2{1, 1}, Vec2{3, 1})

	want := []Vec2{
		{1, 1},
		{1, 2},
		{2, 3},
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
		t.Logf(renderAsString(&grid))
		t.Logf("want: %v", want)
		t.Logf("got: %v", got)
	}
}
