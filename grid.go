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

type Grid[T any] struct {
	Width  int
	Height int
	inner  []T
}

func NewGrid[T any](width, height int) Grid[T] {
	grid := Grid[T]{
		Width:  width,
		Height: height,
		inner:  make([]T, width*height),
	}
	return grid
}

func NewMaxGrid[T int](width, height int) Grid[T] {
	grid := NewGrid[T](width, height)
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			grid.Set(x, y, math.MaxInt)
		}
	}
	return grid
}

func (b *Grid[T]) Set(x, y int, val T) {
	idx := y*b.Width + x
	b.inner[idx] = val
}

func (b *Grid[T]) Get(v Vec2) T {
	idx := v.Y*b.Width + v.X
	return b.inner[idx]
}

func RenderAsString(grid *Grid[int]) string {
	sb := &strings.Builder{}
	sb.WriteString("\n")
	for y := range grid.Height {
		for x := range grid.Width {
			val := grid.Get(Vec2{x, y})
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
