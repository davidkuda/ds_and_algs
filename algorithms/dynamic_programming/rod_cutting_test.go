package dp

import (
	"fmt"
	"testing"
)

func TestCutRod(t *testing.T) {
	fmt.Println("Iterative:")
	prices := []int{1, 5, 8, 9, 10, 17, 20, 24, 30}

	l := 4

	have := CutRod(prices, l)

	want := 13

	fmt.Println()
	if have != want {
		t.Fatalf("Failed: Got %v, Want %v", have, want)
	}
}

func TestCutRod2(t *testing.T) {
	fmt.Println("Recursive:")
	prices := []int{1, 5, 8, 9, 10, 17, 20, 24, 30}
	have := CutRod2(prices, 4)
	want := 13

	if have != want {
		t.Fatalf("Failed: Got %v, Want %v", have, want)
	}
}
