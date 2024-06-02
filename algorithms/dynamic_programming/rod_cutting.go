package dp

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// 1: Naive approach with a for loop
func CutRodForLoop(p []int, n int) int {
	if n == 0 {
		return 0
	}
	var q int
	// the textbook uses: q := MinInt, but since prices are always > 0,
	// we can use 0 as base q.
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
	var q int
	// the textbook uses: q := MinInt
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

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// 3. a. Dynamic Programming: Memoization / Top Down Approach
func MemoizedCutRod(p []int, n int) int {
	r := make([]int, n+1, n+1)
	return MemoizedCutRodAux(p, n, r)
}

func MemoizedCutRodAux(p []int, n int, r []int) int {
	if r[n] > 0 {
		return r[n]
	}

	var q int
	if n == 0 {
		q = 0
	} else {
		for i := 1; i <= n; i++ {
			q = max(q, p[i]+MemoizedCutRodAux(p, n-i, r))
		}
	}
	r[n] = q
	return q
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// 3. b. Dynamic Programming: Bottom Up Approach
func BottomUpCutRod(p []int, n int) int {
	r := make([]int, n+1, n+1)
	// textbook initialises r[0] = 0, but not necessary in Golang
	var q int
	for j := 1; j <= n; j++ {
		for i := 1; i <= j; i++ {
			q = max(q, p[i]+r[j-i])
		}
		r[j] = q
	}
	return q
}
