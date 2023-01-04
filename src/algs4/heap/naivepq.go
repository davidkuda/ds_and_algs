package heap

import (
	"golang.org/x/exp/constraints"
)

type UnorderedMaxPriorityQueue[T constraints.Ordered] struct {
	items []T
}

func (pq *UnorderedMaxPriorityQueue[T]) Size() int {
	return len(pq.items)
}

func (pq *UnorderedMaxPriorityQueue[T]) Insert(key T) {
	pq.items = append(pq.items, key)
}

func (pq *UnorderedMaxPriorityQueue[T]) DelMax() T {
	var max T
	var maxIndex int
	// naive implementation: search max in all items
	for i, v := range pq.items {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	N := len(pq.items)
	pq.items[N-1], pq.items[maxIndex] = pq.items[maxIndex], pq.items[N-1]
	pq.items = pq.items[:len(pq.items)-1]
	return max
}
