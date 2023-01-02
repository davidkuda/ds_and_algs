package algos

import "testing"

func TestSelectionSort(t *testing.T) {
	nums := []int{10, 21, 7, 5, 3, 12, 8}
	selectionSort(nums)
	expected := []int{3, 5, 7, 8, 10, 12, 21}
	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("Nums were not sorted correctly;\nexpected: %v\nactual: %v", expected, nums)
			break
		}
	}

}