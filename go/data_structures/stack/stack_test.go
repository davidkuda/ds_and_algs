package stack

import "testing"

func TestStack(t *testing.T) {
	expected := [3]int{3, 2, 1}
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	var actual [3]int
	var ok bool
	for i := range actual {
		actual[i], ok = s.Pop()
		if !ok {
			t.Errorf("Tried to access more items than present in queue")
		}
	}

	if actual != expected {
		t.Errorf("Something wrong!\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestQueueIteration(t *testing.T) {
	expected := [3]int{3,2,1}
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	var actual [3]int
	var count int
	for v, ok := s.Pop(); ok != false; v, ok = s.Pop() {
		actual[count] = v
		count++
	}

	if actual != expected {
		t.Errorf("Something wrong!\nexpected: %v\nactual: %v", expected, actual)
	}
}
