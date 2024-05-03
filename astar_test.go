package astar

import (
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestGetSuccessors_Cardinal(t *testing.T) {
	weights := []int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	grid := NewGridFromSlice(3, 3, weights)

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
	weights := []int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
	grid := NewGridFromSlice(3, 3, weights)
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
	equal(t, got, want, &grid)
}

func TestPath_NoDiagonal1(t *testing.T) {
	weights := []int{
		0, 0, 0, 0, 0,
		0, 1, 0, 1, 0,
		0, 1, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(5, 5, weights)

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
	equal(t, got, want, &grid)
}

func TestPath_NoPath(t *testing.T) {
	weights := []int{
		0, 0, 0, 0, 0,
		0, 1, 0, 1, 0,
		0, 0, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(5, 5, weights)

	pathfinder := NewPathfinder(grid)
	got := pathfinder.Find(Vec2{1, 1}, Vec2{3, 1})

	if len(got) != 0 {
		t.Logf(renderWithPathAsString(&grid, got))
		t.Logf("got: %v", got)
		t.Fatalf("len want %d got %d", 0, len(got))
	}
}

func TestPath_NoDiagonal2(t *testing.T) {
	weights := []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 1, 1, 1, 0, 0,
		0, 1, 1, 1, 0, 1, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(8, 4, weights)

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
	equal(t, got, want, &grid)
}

func TestPath_Diagonal1(t *testing.T) {
	weights := []int{
		0, 0, 0, 0, 0,
		0, 1, 0, 1, 0,
		0, 1, 0, 1, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(5, 5, weights)

	pathfinder := NewPathfinder(grid, WithDiagonals())
	got := pathfinder.Find(Vec2{1, 1}, Vec2{3, 1})

	want := []Vec2{
		{1, 1},
		{1, 2},
		{2, 3},
		{3, 2},
		{3, 1},
	}
	equal(t, got, want, &grid)
}

func TestPath_Diagonal2(t *testing.T) {
	weights := []int{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 1, 1, 1, 0, 0,
		0, 1, 1, 1, 0, 1, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}
	grid := NewGridFromSlice(8, 4, weights)

	pathfinder := NewPathfinder(grid, WithDiagonals())
	got := pathfinder.Find(Vec2{1, 1}, Vec2{6, 2})

	want := []Vec2{
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
	}

	if len(got) != len(want) {
		t.Errorf("len want %d got %d", len(want), len(got))
	}

	for i := range want {
		if i >= len(got) {
			break
		}
		if got[i] != want[i] {
			t.Errorf("pos %d want %v got %v", i, want[i], got[i])
		}
	}

	equal(t, got, want, &grid)
}

func TestEuclideanDistance(t *testing.T) {
	tests := []struct {
		name     string
		v1, v2   Vec2
		expected int
	}{
		{
			name:     "distance between same points",
			v1:       Vec2{X: 1, Y: 1},
			v2:       Vec2{X: 1, Y: 1},
			expected: 0,
		},
		{
			name:     "distance between different points",
			v1:       Vec2{X: 1, Y: 1},
			v2:       Vec2{X: 4, Y: 5},
			expected: 5,
		},
		{
			name:     "distance with negative coordinates",
			v1:       Vec2{X: -1, Y: -1},
			v2:       Vec2{X: -4, Y: -5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := euclideanDistance(tt.v1, tt.v2); got != tt.expected {
				t.Errorf("euclideanDistance() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDiagonalDistance(t *testing.T) {
	tests := []struct {
		name string
		v1   Vec2
		v2   Vec2
		want int
	}{
		{
			name: "same point",
			v1:   Vec2{0, 0},
			v2:   Vec2{0, 0},
			want: 0,
		},
		{
			name: "one step diagonal",
			v1:   Vec2{0, 0},
			v2:   Vec2{1, 1},
			want: 1,
		},
		{
			name: "two steps diagonal",
			v1:   Vec2{0, 0},
			v2:   Vec2{2, 2},
			want: 2,
		},
		{
			name: "one step straight, one step diagonal",
			v1:   Vec2{0, 0},
			v2:   Vec2{2, 1},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalDistance(tt.v1, tt.v2); got != tt.want {
				t.Errorf("diagonalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_Larger(t *testing.T) {
	t.Skip("experimental playground")
	weights := []int{
		0, 1, 1, 1, 1, 1, 1, 1,
		0, 0, 0, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 0, 0, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
	}
	grid := NewGridFromSlice(8, 8, weights)

	pathfinder := NewPathfinder(grid, WithDiagonals(), EuclideanDistance())
	got := pathfinder.Find(Vec2{4, 6}, Vec2{1, 0})

	want := []Vec2{
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
	}

	if len(got) != len(want) {
		t.Errorf("len want %d got %d", len(want), len(got))
	}

	for i := range want {
		if i >= len(got) {
			break
		}
		if got[i] != want[i] {
			t.Errorf("pos %d want %v got %v", i, want[i], got[i])
		}
	}

	equal(t, got, want, &grid)
}

func TestPath_PunishChangeDirection(t *testing.T) {
	weights := []int{
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 1, 1, 1,
	}
	grid := NewGridFromSlice(5, 5, weights)

	pathfinder := NewPathfinder(grid, PunishChangeDirection())
	got := pathfinder.Find(Vec2{1, 3}, Vec2{3, 1})

	want := []Vec2{
		{1, 3},
		{1, 2},
		{1, 1},
		{2, 1},
		{3, 1},
	}
	equal(t, got, want, &grid)
}

func TestPunishChangeDirection_Algo(t *testing.T) {
	end := Vec2{7, 2}
	cases := []struct {
		name string
		q    node
		succ Vec2
		want int
	}{
		{
			name: "x adjacent",
			q: node{
				pos: Vec2{1, 2},
				parent: &node{
					pos: Vec2{1, 1},
				},
			},
			succ: Vec2{1, 3},
			want: 0,
		},
		{
			name: "x change dir",
			q: node{
				pos: Vec2{1, 2},
				parent: &node{
					pos: Vec2{1, 1},
				},
			},
			succ: Vec2{2, 3},
			want: 6,
		},
		{
			name: "y adjacent",
			q: node{
				pos: Vec2{2, 1},
				parent: &node{
					pos: Vec2{1, 1},
				},
			},
			succ: Vec2{3, 1},
			want: 0,
		},
		{
			name: "y change dir",
			q: node{
				pos: Vec2{2, 1},
				parent: &node{
					pos: Vec2{1, 1},
				},
			},
			succ: Vec2{3, 2},
			want: 4,
		},
		{
			name: "diag adj",
			q: node{
				pos: Vec2{2, 2},
				parent: &node{
					pos: Vec2{1, 1},
				},
			},
			succ: Vec2{3, 3},
			want: 0,
		},
		{
			name: "diag change dir",
			q: node{
				pos: Vec2{2, 2},
				parent: &node{
					pos: Vec2{1, 1},
				},
			},
			succ: Vec2{3, 4},
			want: 6,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := punishChangeDirection(c.q, c.succ, end)

			if got != c.want {
				t.Errorf("want: %d got: %d", c.want, got)
			}
		})
	}
}

func equal(t *testing.T, got, want []Vec2, grid *Grid[int]) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("len want %d got %d", len(want), len(got))
	}

	for i := range want {
		if i >= len(got) {
			break
		}
		if got[i] != want[i] {
			t.Errorf("pos %d want %v got %v", i, want[i], got[i])
		}
	}

	if t.Failed() {
		t.Logf("want: %v", want)
		t.Logf(renderWithPathAsString(grid, want))
		t.Logf("got: %v", got)
		t.Logf(renderWithPathAsString(grid, got))
	}
}

var _ = renderAsString // suppress unused

// renderAsString returns a string representation of the grid.
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

// renderWithPathAsString returns a string representation of the grid with the
// path drawn.
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
				// sb.WriteString(strconv.Itoa(val))
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
