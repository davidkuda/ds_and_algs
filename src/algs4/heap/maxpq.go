package heap

import (
	"golang.org/x/exp/constraints"
)

// sorts all items of data in place
func HeapSort[T constraints.Ordered](data *[]T) {
	pq := NewMaxPriorityQueue[T](len(*data), len(*data))
	pq.items = *data
	sorted := data
	N := len(pq.items) - 1
	// first pass: order items so that the array represents a MaxHeap
	for k := N / 2; k >= 0; k-- {
		pq.sink(k)
	}
	// second pass: push max to last element
	for N > 0 {
		pq.swap(0, N)
		pq.items = pq.items[:N]
		N--
		pq.sink(0)
	}
	pq.items = *sorted
}

// ? what about capacity? e.g., what if we want to keep only 20 values?
type MaxPriorityQueue[T constraints.Ordered] struct {
	items []T
}

// use this constructor if you know the number of elements that the pq should hold.
// otherwise, the slice may need to be resized.
func NewMaxPriorityQueue[T constraints.Ordered](size, cap int) *MaxPriorityQueue[T] {
	return &MaxPriorityQueue[T]{
		items: make([]T, size, cap),
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
	N := len(pq.items) - 1
	pq.swap(0, N)
	pq.items = pq.items[:N]
	pq.sink(0)
	return max, true
}

func (pq *MaxPriorityQueue[T]) Insert(item T) {
	pq.items = append(pq.items, item)
	N := pq.Size() - 1
	pq.swim(N) // aka heapifyUp
}

// --- Helper Functions

// aka "heapifyUp"
func (pq *MaxPriorityQueue[T]) swim(k int) {
	for k > 0 && pq.items[k] > pq.parent(k) {
		pq.swap(parentIndex(k), k)
		k = parentIndex(k)
	}
}

// aka "heapifyDown": swap parent with bigger child
func (pq *MaxPriorityQueue[T]) sink(k int) {
	// if no left child, then certainly no right child
	for pq.hasLeftChild(k) {
		largerChildIndex := leftChildIndex(k)
		if pq.hasRightChild(k) && pq.rightChild(k) > pq.leftChild(k) {
			largerChildIndex = rightChildIndex(k)
		}

		if pq.items[k] > pq.items[largerChildIndex] {
			break
		}

		pq.swap(k, largerChildIndex)
		k = largerChildIndex
	}
}

func (pq *MaxPriorityQueue[T]) swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

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
