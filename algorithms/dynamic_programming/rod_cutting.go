package dp

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// 1: Naive approach with a for loop
func CutRodForLoop(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := MinInt
	for i := 1; i <= n; i++ {
		q = max(q, p[i]+CutRodForLoop(p, n-i))
	}
	return q
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// 2: Naive approach recursion only (no for loop)
func CutRodRecursionOnly(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := MinInt
	q = cr(p, n, q, 1)
	return q
}

func cr(p []int, n, q, i int) int {
	if i > n {
		return q
	}
	// same as in the for loop above:
	q = max(q, p[i]+CutRodRecursionOnly(p, n-i))
	return cr(p, n, q, i+1)
}
