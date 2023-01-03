package heap

import (
	"golang.org/x/exp/constraints"
)

type MaxPriorityQueue[T constraints.Ordered] struct {
	items []T
}

func (pq *MaxPriorityQueue[T]) Size() int {
	return len(pq.items)
}

func (pq *MaxPriorityQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *MaxPriorityQueue[T]) Insert() {}

func (pq *MaxPriorityQueue[T]) DelMax() T {
	var r T
	return r
}

func (pq *MaxPriorityQueue[T]) swim() {}

func (pq *MaxPriorityQueue[T]) sink() {}
