package astar

// heapNode is a node in the min heap.
type heapNode struct {
	pos Vec2
	f   int
}

// minHeap is a specialized 2d min heap that orders heapNodes by their f value.
type minHeap struct {
	width  int
	height int
	inner  []heapNode
}

// newMinHeap creates a new MinHeap with the given width and height.
func newMinHeap(width, height int) *minHeap {
	heap := &minHeap{
		width:  width,
		height: height,
		inner:  []heapNode{},
	}

	return heap
}

// len returns the number of elements in the heap.
func (h *minHeap) len() int {
	return len(h.inner)
}

// push adds an element to the heap.
func (h *minHeap) push(elem heapNode) {
	h.inner = append(h.inner, elem)
	curr := len(h.inner) - 1
	for {
		if curr == 0 {
			break
		}

		parent := (curr - 1) / 2

		// min heap property satisfied
		if h.inner[parent].f <= h.inner[curr].f {
			break
		}

		h.inner[curr], h.inner[parent] = h.inner[parent], h.inner[curr]
		curr = parent
	}
}

// peek returns the element at the top of the heap. If the heap is empty, it
// will panic.
func (h *minHeap) peek() heapNode {
	if len(h.inner) == 0 {
		panic("heap empty")
	}
	return h.inner[0]
}

// pop removes and returns the element at the top of the heap. If the heap is
// empty,it will panic.
func (h *minHeap) pop() heapNode {
	if len(h.inner) == 0 {
		panic("heap empty")
	}

	result := h.inner[0]

	// arbitrarily choose last item as head
	h.inner[0] = h.inner[len(h.inner)-1]
	h.inner = h.inner[:len(h.inner)-1]

	p := 0
	for {
		pn := p // copy p
		left := 2*p + 1
		right := 2*p + 2

		// choose the smallest child for p
		if len(h.inner) > left && h.inner[left].f < h.inner[p].f {
			p = left
		}

		if len(h.inner) > right && h.inner[right].f < h.inner[p].f {
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
