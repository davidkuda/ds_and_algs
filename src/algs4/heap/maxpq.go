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

// --- Helper Functions

// Parent

func parentIndex(k int) int {
	return (k - 1) / 2
}

func hasParent(k int) bool {
	return parentIndex(k) >= 0
}

func (pq *MaxPriorityQueue[T]) parent(k int) T {
	return pq.items[parentIndex(k)]
}

// Left Child

func leftChildIndex(k int) int {
	return k*2 + 1
}

func (pq *MaxPriorityQueue[T]) hasLeftChild(k int) bool {
	return leftChildIndex(k) < len(pq.items)
}

func (pq *MaxPriorityQueue[T]) leftChild(k int) T {
	return pq.items[leftChildIndex(k)]
}

// Right Child

func rightChildIndex(k int) int {
	return k*2 + 2
}

func (pq *MaxPriorityQueue[T]) hasRightChild(k int) bool {
	return rightChildIndex(k) < len(pq.items)
}

func (pq *MaxPriorityQueue[T]) rightChild(k int) T {
	return pq.items[rightChildIndex(k)]
}

