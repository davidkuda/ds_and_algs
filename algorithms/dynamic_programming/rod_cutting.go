package dp

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func CutRod(prices []int, rodLength int) int {
	if rodLength == 0 {
		return 0
	}

	highestPrice := prices[rodLength]

	for i := 0; i < rodLength; i++ {
		fmt.Println(i)
		highestPrice = max(highestPrice, prices[i]+CutRod(prices, rodLength-i-1))
	}

	return highestPrice
}
