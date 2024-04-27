package astar

type Vec2 struct {
	X int
	Y int
}

type Node struct {
	F      int
	Pos    Vec2
	Weight int
	Parent *Node
	G      int
	H      int
	Open   bool
	Closed bool
}

type MinHeap struct {
	width  int
	height int
	inner  []Node
}

func NewMinHeap(width, height int) *MinHeap {
	heap := &MinHeap{
		width:  width,
		height: height,
	}

	return heap
}

func (h *MinHeap) Len() int {
	return len(h.inner)
}

func (h *MinHeap) Push(elem Node) {
	h.inner = append(h.inner, elem)
	curr := len(h.inner) - 1
	for {
		if curr == 0 {
			break
		}

		parent := (curr - 1) / 2
		if h.inner[curr].F < h.inner[parent].F {
			h.inner[curr], h.inner[parent] = h.inner[parent], h.inner[curr]
			curr = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) Peek() Node {
	if len(h.inner) == 0 {
		panic("heap empty")
	}
	return h.inner[0]
}

func (h *MinHeap) Pop() (result Node) {
	if len(h.inner) == 0 {
		panic("heap empty")
	}

	result = h.inner[0]

	// arbitrarily choose last item has head
	h.inner[0] = h.inner[len(h.inner)-1]
	h.inner = h.inner[:len(h.inner)-1]

	p := 0
	for {
		pn := p
		left := 2*p + 1
		right := 2*p + 2

		if len(h.inner) > left && h.inner[left].F < h.inner[p].F {
			p = left
		}

		if len(h.inner) > right && h.inner[right].F < h.inner[p].F {
			p = right
		}

		if p == pn {
			break
		}

		h.inner[p], h.inner[pn] = h.inner[pn], h.inner[p]
	}

	return result
}
