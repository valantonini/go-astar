package main

import (
	"math"
	"strconv"
	"strings"
)

var neighbours = [4][2]int{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

type Grid struct {
	Width  int
	Height int
	Cells  []int
}

func (b *Grid) Set(x, y int, val int) {
	idx := y*b.Width + x
	b.Cells[idx] = val
}

func (b *Grid) Get(x, y int) int {
	idx := y*b.Width + x
	return b.Cells[idx]
}

func (b *Grid) Neighbours(x, y int) []Node {
	results := make([]Node, 0, len(neighbours))
	for _, n := range neighbours {
		x := x + n[0]
		y := y + n[1]

		if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
			continue
		}

		n := Node{
			F:      0,
			Pos:    Vec2{x, y},
			Weight: b.Get(x, y),
		}
		results = append(results, n)
	}
	return results
}

func NewGrid(width, height int) Grid {
	grid := Grid{
		Width:  width,
		Height: height,
		Cells:  make([]int, width*height),
	}
	return grid
}

func NewMaxGrid(width, height int) Grid {
	grid := NewGrid(width, height)
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			grid.Set(x, y, math.MaxInt)
		}
	}
	return grid
}

func RenderAsString(grid *Grid) string {
	sb := &strings.Builder{}
	sb.WriteString("\n")
	for y := range grid.Height {
		for x := range grid.Width {
			val := grid.Get(x, y)
			switch val {
			case 0:
				sb.WriteString(".")
				break
			default:
				sb.WriteString(strconv.Itoa(val))
				break
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
