package heap

import (
	"golang.org/x/exp/constraints"
)

type MaxPriorityQueue[T constraints.Ordered] struct {
	items []T
}

// use this constructor if you know the number of elements that the pq should hold.
// otherwise, the slice may need to be resized.
func NewMaxPriorityQueue[T constraints.Ordered](size int) *MaxPriorityQueue[T] {
	return &MaxPriorityQueue[T]{
		items: make([]T, size),
	}
}

func (pq *MaxPriorityQueue[T]) Size() int {
	return len(pq.items)
}

func (pq *MaxPriorityQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *MaxPriorityQueue[T]) Insert(v int) {}

func (pq *MaxPriorityQueue[T]) DelMax() T {
	var r T
	return r
}

func (pq *MaxPriorityQueue[T]) swim() {}

func (pq *MaxPriorityQueue[T]) sink() {}

func parent(k int) int {
	return (k - 1) / 2
}

func leftChild(k int) int {
	return k*2 + 1
}

func rightChild(k int) int {
	return k*2 + 2
}

