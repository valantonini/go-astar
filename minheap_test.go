package main

import "testing"

type Node struct {
	F int
}
type MinHeap struct {
	inner []Node
}

func (h *MinHeap) Push(elem Node) {
	h.inner = append(h.inner, elem)
	curr := len(h.inner) - 1

	for {
		if len(h.inner) == 1 {
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
		return Node{}
	}
	return h.inner[0]
}

func TestMinHeap_Peek(t *testing.T) {
	heap := &MinHeap{}
	heap.Push(Node{5})
	heap.Push(Node{7})
	heap.Push(Node{3})
	heap.Push(Node{9})
	heap.Push(Node{1})
	heap.Push(Node{6})

	got := heap.Peek()

	if got.F != 1 {
		t.Errorf("want %v got %v", 1, got.F)
	}
}
