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

// return max and true if pq.Size > 0; else, return default T val and false
func (pq *MaxPriorityQueue[T]) Peek() (max T, ok bool) {
	if pq.IsEmpty() {
		var defaultVal T
		return defaultVal, false
	} else {
		return pq.items[0], true
	}
}

// aka pq.DelMax -> Remove Max from PQ
func (pq *MaxPriorityQueue[T]) Poll() (max T, ok bool) {
	if pq.IsEmpty() {
		var defaultVal T
		return defaultVal, false
	}
	max = pq.items[0]
	pq.swap(0, len(pq.items)-1)
	heapifyDown()
	return max, true
}


func (pq *MaxPriorityQueue[T]) DelMax() T {
	pq.swim(len(pq.items))
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

