package heap

import "testing"

func TestMaxPQ(t *testing.T) {
	pq := MaxPriorityQueue[int]{}

	values := []int{10, 7, 12, 20, 5, 3, 21}

	for _, v := range values {
		pq.Insert(v)
	}

	if pq.Size() != 7 {
		t.Errorf("Expected size of 7, but got %d", pq.Size())
	}
	max, ok := pq.Poll()
	if !ok {
		t.Errorf("Tried to poll from empty pq?")
	}

	if max != 21 {
		t.Errorf("Expected max to be 21, but was %d", max)
	}

	if pq.Size() != 6 {
		t.Errorf("Expected len of 6, but got %d", pq.Size())
	}
}

func TestParent(t *testing.T) {
	
	if parentIndex(0) != 0 {
		t.Errorf("Failed to get parentIndex, exp 0, got %d", parentIndex(0))
	}

	if parentIndex(1) != 0 {
		t.Errorf("Failed to get parentIndex, exp 0, got %d", parentIndex(1))
	}

	if parentIndex(2) != 0 {
		t.Errorf("Failed to get parentIndex, exp 0, got %d", parentIndex(2))
	}

	if parentIndex(3) != 1 {
		t.Errorf("Failed to get parentIndex, exp 0, got %d", parentIndex(3))
	}
}