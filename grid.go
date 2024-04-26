package main

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
