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
	max := pq.DelMax()

	if max != 21 {
		t.Errorf("Expected max to be 21, but was %d", max)
	}
}

func TestParent(t *testing.T) {
	
	if parent(0) != 0 {
		t.Errorf("Failed to get parent, exp 0, got %d", parent(0))
	}

	if parent(1) != 0 {
		t.Errorf("Failed to get parent, exp 0, got %d", parent(1))
	}

	if parent(2) != 0 {
		t.Errorf("Failed to get parent, exp 0, got %d", parent(2))
	}

	if parent(3) != 1 {
		t.Errorf("Failed to get parent, exp 0, got %d", parent(3))
	}
}