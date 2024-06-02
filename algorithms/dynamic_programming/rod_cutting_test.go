package dp

import (
	"testing"
)

// - - - - - - - - 0  1  2  3  4   5   6   7   8   9  10 
var prices = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
var tests = []struct {
	input int
	want  int
}{
	{1, 1},
	{2, 5},
	{3, 8},
	{4, 10},
	{5, 13},
	{6, 17},
	{7, 18},
	{8, 22},
	{9, 25},
	{10, 30},
}

func TestCutRodForLoop(t *testing.T) {
	for _, test := range tests {
		got := CutRodForLoop(prices, test.input)
		if got != test.want {
			t.Errorf("CutRod(%d) = %d (want %d)", test.input, got, test.want)
		}
	}
}

func TestCutRodRecursionOnly(t *testing.T) {
	for _, test := range tests {
		got := CutRodRecursionOnly(prices, test.input)
		if got != test.want {
			t.Errorf("CutRod(%d) = %d (want %d)", test.input, got, test.want)
		}
	}
}
