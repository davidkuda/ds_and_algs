package dp

// iterative approach
func CutRod(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := -1
	for i := 0; i < n; i++ {
		q = max(q, p[i]+CutRod(p, n-i-1))
	}
	return q
}

// avoid iteration, failed attempt
func CutRod2(p []int, n int) int {
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
		p[i]+CutRod(p, n-i-1),
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

	currentMax := p[i] + CutRod(p, n-i)
	if currentMax > q {
		q = currentMax
	}

	return recursiveCut(p, n, q, i+1)
}
