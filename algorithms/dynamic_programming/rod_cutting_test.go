package dp

import (
	"testing"
)

func TestCutRod(t *testing.T) {
	prices := []int{1, 5, 8, 9, 10, 17, 20, 24, 30}

	l := 4

	have := CutRod(prices, l)

	want := 13

	if have != want {
		t.Fatalf("Failed: Got %v, Want %v", have, want)
	}

}
