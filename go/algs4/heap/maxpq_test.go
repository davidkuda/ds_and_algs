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
