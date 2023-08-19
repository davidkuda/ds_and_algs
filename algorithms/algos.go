package algos

import (
	"cmp"  // new in Go 1.21, see https://stackoverflow.com/a/70562597
	"math/rand"
	"time"
)

// --- SortingAlgorithms

func SelectionSort[T cmp.Ordered](a []T) {
	N := len(a)
	for i := 0; i < N; i++ {
		// identify min
		min := i
		for j := i + 1; j < N; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		// exchange
		a[i], a[min] = a[min], a[i]
	}
}

func InsertionSort[T cmp.Ordered](a []T) {
	N := len(a)
	for i := 0; i < N; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
}

func ShellSort[T cmp.Ordered](a []T) {
	N := len(a)
	h := 1

	for h < N/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, ..
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h = h / 3
	}
}

func Shuffle[T cmp.Ordered](a []T) {
	// rand.Seed will assure random values; otherwise, rand.Intn will yield same value every time
	// see https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		r := rand.Intn(i + 1)
		a[i], a[r] = a[r], a[i]
	}
}

func ShuffleReturn[T cmp.Ordered](a []T) []T {
	for i := range a {
		r := rand.Intn(i + 1)
		a[i], a[r] = a[r], a[i]
	}
	return a
}

/*
Basic Plan:
- Shuffle the Array a (preserving randomness: shuffling is needed for performance guarantee)
- Partition the Array, so that for some j
  - entry a[j] is in place,
  - no larger entry to the left of j
  - no smaller entry to the right of j

- Sort each piece recursively

Complexity:
  - Average time: 1.39NlogN; Worst Case: N*N; If you shuffle first, worst case is very unlikely
  - Space: O(1); Sort in place; but call stack is logN, due to recursive function calls
  - QuickSort is faster than MergeSort, because the implementation is fairly simple:
    compare, increase a pointer, swap;
*/
func QuickSort[T cmp.Ordered](a []T) {
	Shuffle(a)
	quickSort(a, 0, len(a)-1)
}

func quickSort[T cmp.Ordered](a []T, left, right int) {
	if left >= right {
		return
	}
	p := partition(a, left, right)
	quickSort(a, left, p-1)
	quickSort(a, p, right)
}

func partition[T cmp.Ordered](a []T, left, right int) int {
	pivot := a[left]
	i, j := left+1, right

	for {
		for i < right && a[i] < pivot {
			i++
		}

		for j > left && a[j] > pivot {
			j--
		}

		if j <= i {
			break
		}

		a[i], a[j] = a[j], a[i]
	}
	a[left], a[j] = a[j], a[left]
	return j
}

// return the kth largest element of a slice "a"
// e.g. k = len(a)/2 == median; k = len(a)-1 == max(a); k = 0 == min(a)
func QuickSelect[T cmp.Ordered](a []T, k int) T {
	var sol T
	Shuffle(a)
	left, right := 0, len(a)-1
	for left <= right {
		p := partition2(a, left, right)
		if p < k {
			left = p + 1
		} else if p > k {
			right = p - 1
		} else {
			sol = a[k]
			break
		}
	}
	return sol
}

// --- LinkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverse by creating a new Node
func reverseList1(head *ListNode) *ListNode {
	var n *ListNode
	for head != nil {
		n = &ListNode{Val: head.Val, Next: n}
		head = head.Next
	}
	return n
}

// reverse by storing head as tempNext
func reverseList2(head *ListNode) *ListNode {
	var r *ListNode
	for head != nil {
		tempNext := head.Next
		head.Next = r
		r = head
		head = tempNext
	}
	return r
}

func middleNode(head *ListNode) *ListNode {
	mid := head
	for head != nil {
		if head.Next == nil {
			break
		}
		mid = mid.Next
		head = head.Next.Next
	}
	return mid
}

// --- MATH

// Fibonacci:

// The solution I came up with
func fibonacci(n int) int {
    if n == 0 {
        return 0
    }
    a := 0
    b := 1
    for i := 1; i < n; i++ {
        b = a + b
        a = b - a
    }
    return b
}

// recursive Fibonacci
func fib(N int) int {
    if N <= 1 {
        return N
    }
    return fib(N - 1) + fib(N - 2)
}

// with cache-map
var cache = map[int]int{0: 0, 1: 1}

func fib2(N int) int {
    if _, ok := cache[N]; ok {
        return cache[N]
    }
    cache[N] = fib(N - 1) + fib(N - 2)
    return cache[N]
}

// cached int array fibonacci
func fib3(N int) int {
    if N <= 1 {
        return N
    }
    
    cache := make([]int, N + 1)
    cache[1] = 1
    for i := 2; i <= N; i++ {
        cache[i] = cache[i - 1] + cache[i - 2]
    }
                      
    return cache[N]
}
