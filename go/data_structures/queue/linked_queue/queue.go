package queue

type Queue struct {
	head *node
	tail *node
	len  int
}

type node struct {
	v    string
	next *node
}

func (q *Queue) GetSize() int {
	return q.len
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue) Enqueue(v string) {
	n := &node{v: v}
	if q.len == 0 {
		q.head = n
		q.tail = n
		q.len++
	} else {
		q.tail.next = n
		q.tail = q.tail.next
		q.len++
	}
}

func (q *Queue) Dequeue() string {
	v := q.head.v
	q.head = q.head.next
	q.len--
	if q.len == 0 {
		q.tail = nil
	}
	return v
}
