package bst

import "testing"

func TestQueue(t *testing.T) {
	expected := [3]string{"first", "second", "third"}
	q := NewQueue[string]()
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")

	var actual [3]string
	var ok bool
	for i := range actual {
		actual[i], ok = q.Dequeue()
		if !ok {
			t.Errorf("Tried to access more items than present in queue")
		}
	}

	if actual != expected {
		t.Errorf("Something wrong!\nexpected: %v\nactual: %v", expected, actual)
	}
}
