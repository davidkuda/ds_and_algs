package linkedqueue

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	len  int
}

type node[T any] struct {
	v    T
	next *node[T]
}

func (q *Queue[T]) GetSize() int {
	return q.len
}

func (q *Queue[T]) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue[T]) Enqueue(v T) {
	n := &node[T]{v: v}
	if q.IsEmpty() {
		q.head = n
		q.tail = n
		q.len++
	} else {
		q.tail.next = n
		q.tail = q.tail.next
		q.len++
	}
}

func (q *Queue[T]) Dequeue() T {
	v := q.head.v
	q.head = q.head.next
	q.len--
	if q.len == 0 {
		q.tail = nil
	}
	return v
}
