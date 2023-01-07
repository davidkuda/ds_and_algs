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

	max2, ok := pq.Peek()
	if !ok {
		t.Errorf("pq should not be empty")
	}
	if max2 != 20 {
		t.Errorf("Expected 20, got %d", max2)
	}
}

func TestHeapSort(t *testing.T) {
	nums := []int8{10, 7, 21, 3, 5}
	expectedNums := []int8{3, 5, 7, 10, 21}
	HeapSort(&nums)
	for i, num := range nums {
		if num != expectedNums[i] {
			t.Errorf("Sort was unsuccesful, got %v", nums)
			break
		}
	}

	word := "heapsort"
	data := []rune{}
	for _, c := range word {
		data = append(data, c)
	}
	HeapSort(&data)
	expected := []rune{'a', 'e', 'h', 'o', 'p', 'r', 's', 't'}
	for i := range data {
		if data[i] != expected[i] {
			t.Errorf("Sort was unsuccesful, got %c", data)
			break
		}
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
