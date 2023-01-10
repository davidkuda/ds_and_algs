package algos

import "testing"

func TestSelectionSort(t *testing.T) {
	nums := []int{10, 21, 7, 5, 3, 12, 8}
	SelectionSort(nums)
	expected := []int{3, 5, 7, 8, 10, 12, 21}
	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("Nums were not sorted correctly;\nexpected: %v\nactual: %v", expected, nums)
			break
		}
	}
}

func TestInsertionSort(t *testing.T) {
	nums := []int{10, 21, 7, 5, 3, 12, 8}
	InsertionSort(nums)
	expected := []int{3, 5, 7, 8, 10, 12, 21}
	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("Nums were not sorted correctly;\nexpected: %v\nactual: %v", expected, nums)
			break
		}
	}
}

func TestShellSort(t *testing.T) {
	nums := []int{10, 21, 7, 5, 3, 12, 8}
	ShellSort(nums)
	expected := []int{3, 5, 7, 8, 10, 12, 21}
	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("Nums were not sorted correctly;\nexpected: %v\nactual: %v", expected, nums)
			break
		}
	}
}

func TestShuffle(t *testing.T) {
	nums := []int{3, 5, 7, 8, 10, 12, 21}
	Shuffle(nums)
	notExpected := []int{3, 5, 7, 8, 10, 12, 21}
	var different bool
	for i := range nums {
		if nums[i] == notExpected[i] {
			different = true
			break
		}
	}
	if !different {
		t.Errorf("Nums were not sorted correctly;\nexpected: %v\nactual: %v", notExpected, nums)
	}
}

func TestShuffleReturn(t *testing.T) {
	nums := []int{3, 5, 7, 8, 10, 12, 21}
	originalOrder := []int{3, 5, 7, 8, 10, 12, 21}
	shuffledNums := ShuffleReturn(nums)

	var different bool
	for i := range nums {
		if shuffledNums[i] == originalOrder[i] {
			different = true
			break
		}
	}
	if !different {
		t.Errorf("Nums were not sorted correctly;\nexpected: %v\nactual: %v", originalOrder, nums)
	}

	// nums should remain same after shuffle, i.e. in place changes should not happen
	var sameValues bool
	for i := range nums {
		if nums[i] == originalOrder[i] {
			sameValues = true
		} else {
			sameValues = false
			break
		}
	}
	if !sameValues {
		t.Errorf("Shuffle changed slice in place!")
	}
}

func TestQuickSort(t *testing.T) {
	nums := []int{10, 21, 7, 5, 3, 12, 8, 42, 3, 25, 3, 30, 24, 4, 8, 9, 10, 18}
	QuickSort(nums)
	expected := []int{3, 3, 3, 4, 5, 7, 8, 8, 9, 10, 10, 12, 18, 21, 24, 25, 30, 42}
	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("Nums were not sorted correctly;\nexpected: %v\nactual: %v", expected, nums)
			break
		}
	}
}
