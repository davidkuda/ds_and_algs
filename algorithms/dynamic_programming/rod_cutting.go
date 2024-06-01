package dp

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func CutRod(prices []int, rodLength int) int {
	fmt.Println("iterative:", rodLength)
	if rodLength == 0 {
		return 0
	}

	highestPrice := prices[rodLength]

	for i := 0; i < rodLength; i++ {
		highestPrice = max(highestPrice, prices[i]+CutRod(prices, rodLength-i-1))
	}

	return highestPrice
}

func CutRod2(p []int, n int) int {
	fmt.Println("recursive:", n)
	if n == 0 {
		return 0
	}

	q := MinInt

	q = max(q, p[n]+CutRod2(p, n-1))

	if n > 0 {
		q = max(q, p[n]+CutRod2(p, n-1))
	}

	fmt.Println(q)

	return q
}

