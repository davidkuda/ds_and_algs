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
