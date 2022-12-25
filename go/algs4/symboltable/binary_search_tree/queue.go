package bst

type Queue[T any] struct {
	items []T
	front int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if  q.front == len(q.items) {
		var v T
		return v, false
	}
	v := (q.items)[q.front]
	q.front++
	return v, true
}
