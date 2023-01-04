package linkedqueue

import "testing"

func TestQueue(t *testing.T) {
	expected := [3]string{"first", "second", "third"}
	q := Queue[string]{}
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")

	var actual [3]string
	for i := range actual {
		actual[i] = q.Dequeue()
	}

	if actual != expected {
		t.Errorf("Something wrong!\nexpected: %v\nactual: %v", expected, actual)
	}
}
