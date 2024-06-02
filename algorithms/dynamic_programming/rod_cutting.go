package dp


const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// 1: Naive approach with a for loop
func CutRodForLoop(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := MinInt
	for i := 0; i < n; i++ {
		q = max(q, p[i]+CutRodForLoop(p, n-i-1))
	}
	return q
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// 2: Naive approach recursion only (no for loop)
func CutRodRecursionOnly(p []int, n int) int {
	q := cr(p, n, 0)
	return q
}

func cr(p []int, n, i int) int {
	if n == 0 {
		return 0
	}
	q := -1
	q = max(q, p[i]+cr(p, n-1, i+1))
	if n > 0 {
		q = max(q, p[i]+cr(p, n-1, i+1))
	}
	return q
}

// ChatGPT version (fails at 0)
func CutRodC(p []int, n int) int {
	if n == 0 {
		return 0
	}

	q := -1
	q = recursiveCut(p, n, q, 0)

	return q
}

func recursiveCut(p []int, n int, q int, i int) int {
	if i == n {
		return q
	}

	q = max(
		q,
		p[i]+CutRodC(p, n-i-1),
	)

	return recursiveCut(p, n, q, i+1)
}

// ChatGPT version (treats 0 as rod length 1)
func CutRodChatGPT(p []int, n int) int {
	if n == 0 {
		return p[0]
	}

	q := -1
	q = recursiveCut0(p, n, q, 1) // start from 1 to handle the correct indexing

	return q
}

func recursiveCut0(p []int, n int, q int, i int) int {
	if i > n {
		return q
	}

	currentMax := p[i] + CutRodChatGPT(p, n-i)
	if currentMax > q {
		q = currentMax
	}

	return recursiveCut0(p, n, q, i+1)
}
