package astar

// MinHeap is a simple min heap implementation.
type MinHeap struct {
	width  int
	height int
	inner  []Node
}

// NewMinHeap creates a new MinHeap with the given width and height.
func NewMinHeap(width, height int) *MinHeap {
	heap := &MinHeap{
		width:  width,
		height: height,
	}

	return heap
}

// Len returns the number of elements in the heap.
func (h *MinHeap) Len() int {
	return len(h.inner)
}

// Push adds an element to the heap.
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

// Peek returns the element at the top of the heap. If the heap is empty, it
// will panic.
func (h *MinHeap) Peek() Node {
	if len(h.inner) == 0 {
		panic("heap empty")
	}
	return h.inner[0]
}

// Pop removes and returns the element at the top of the heap. If the heap is
// empty,it will panic.
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
		pn := p // copy p
		left := 2*p + 1
		right := 2*p + 2

		// choose the smallest child for p
		if len(h.inner) > left && h.inner[left].F < h.inner[p].F {
			p = left
		}

		if len(h.inner) > right && h.inner[right].F < h.inner[p].F {
			p = right
		}

		// no smaller child existed, heap property is satisfied
		if p == pn {
			break
		}

		h.inner[p], h.inner[pn] = h.inner[pn], h.inner[p]
	}

	return result
}
