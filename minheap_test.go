package main

import (
	"slices"
	"testing"
)

func TestMinHeap_Peek(t *testing.T) {
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
				heap.Push(Node{F: n})
			}

			want := make([]int, len(c.data))
			copy(want, c.data)
			slices.Sort(want)

			got := heap.Peek()

			if got.F != want[0] {
				t.Errorf("want %v got %v", want[0], got.F)
			}
		})
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
				heap.Push(Node{F: n})
			}

			want := make([]int, len(c.data))
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
