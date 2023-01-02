package heap

import (
	"testing"
)

func TestNaivePQ(t *testing.T) {
	pq := UnorderedMaxPriorityQueue[int]{}
	pq.Insert(108)
	pq.Insert(42)
	pq.Insert(27)
	pq.Insert(54)
	pq.Insert(81)

	if pq.N != 5 {
		t.Errorf("Expected size of pq to be 5, but was %d", pq.N)
	}

	max := pq.DelMax()
	if max != 108 {
		t.Errorf("Got wrong max, expeced 108, got %d", max)
	}

	if pq.N != 4 {
		t.Errorf("Expected size of pq to be 4, but was %d", pq.N)
	}
}