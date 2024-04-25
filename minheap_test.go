package main

import (
	"slices"
	"testing"
)

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

func TestMinHeap_Pop(t *testing.T) {
	cases := []struct {
		name string
		data []int
	}{
		{
			name: "case 1",
			data: []int{5, 7, 3, 9, 1, 6},
		},
		{
			name: "case 2",
			data: []int{4, 3, 2, 1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			heap := &MinHeap{}
			for _, n := range c.data {
				heap.Push(Node{n})
			}

			want := make([]int, 0, len(c.data))
			copy(want, c.data)
			slices.Sort(want)

			for _, n := range want {
				got := heap.Pop()
				if got.F != n {
					t.Errorf("want %v got %v", n, got.F)
				}
			}
		})
	}
}
