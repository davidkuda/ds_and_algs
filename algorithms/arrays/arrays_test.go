package arrays

import "testing"

func reverse_slice_test(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	reverse(a)
	exp := []int{5, 4, 3, 2, 1}

	for i := range a {
		if a[i] != exp[i] {
			t.Fatalf("Reverse failed: got: %v; expected: %v", a, exp)
		}
	}
}
