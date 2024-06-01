package dp

import (
	"testing"
)

func TestCutRod(t *testing.T) {
	prices := []int{1, 5, 8, 9, 10, 17, 20, 24, 30}
	have := CutRod(prices, 4)
	want := 13
	if have != want {
		t.Fatalf("Failed: Got %v, Want %v", have, want)
	}
}

func TestCutRodChatGPT(t *testing.T) {
	prices := []int{1, 5, 8, 9, 10, 17, 20, 24, 30}
	have := CutRodChatGPT(prices, 0)
	want := 1

	if have != want {
		t.Fatalf("Failed: Got %v, Want %v", have, want)
	}
}

func TestCutRodC(t *testing.T) {
	// doesn't handle 0 correctly
	prices := []int{1, 5, 8, 9, 10, 17, 20, 24, 30}
	have := CutRodC(prices, 4)
	want := 13

	if have != want {
		t.Fatalf("Failed: Got %v, Want %v", have, want)
	}
}
