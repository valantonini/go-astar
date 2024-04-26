package main

import "strings"

type Grid struct {
	Width  int
	Height int
	Cells  []byte
}

func (b *Grid) Set(x, y int, val byte) {
	idx := y*b.Width + x
	b.Cells[idx] = val
}

func (b *Grid) Get(x, y int) byte {
	idx := y*b.Width + x
	return b.Cells[idx]
}

func NewGrid(width, height int) Grid {
	return Grid{
		Width:  width,
		Height: height,
		Cells:  make([]byte, width*height),
	}
}

func RenderAsString(grid *Grid) string {
	sb := &strings.Builder{}
	sb.WriteString("\n")
	for x := range grid.Width {
		for y := range grid.Height {
			val := grid.Get(x, y)
			switch val {
			case 0:
				sb.WriteString(".")
				break
			default:
				sb.WriteRune(rune(val + 48))
				break
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
