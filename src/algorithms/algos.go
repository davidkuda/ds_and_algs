package algos

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

// --- SortingAlgorithms

type ordered constraints.Ordered

func SelectionSort[T constraints.Ordered](a []T) {
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

func InsertionSort[T constraints.Ordered](a []T) {
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

func ShellSort[T constraints.Ordered](a []T) {
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

func Shuffle[T constraints.Ordered](a []T) {
	// rand.Seed will assure random values; otherwise, rand.Intn will yield same value every time
	// see https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		r := rand.Intn(i + 1)
		a[i], a[r] = a[r], a[i]
	}
}

func ShuffleReturn[T constraints.Ordered](a []T) []T {
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
func QuickSort[T ordered](a []T) {
	Shuffle(a)
	quickSort(a, 0, len(a)-1)
}

func quickSort[T ordered](a []T, left, right int) {
	if left >= right {
		return
	}
	partitionPoint := partition(a, left, right)
	quickSort(a, left, partitionPoint-1)
	quickSort(a, partitionPoint, right)
}

func partition[T ordered](a []T, left, right int) int {
	pi := (left + right) / 2
	pivot := a[pi]
	for left < right {
		for a[left] < pivot {
			left++
		}
		for a[right] > pivot {
			right--
		}
		if left <= right {
			a[left], a[right] = a[right], a[left]
			left++
			right--
		}
	}
	// return partition point -> where elements on left of partition points are
	// smaller than the pivot, and the elements to the right are larger.
	return left
}

// return the kth largest element of a slice "a"
// e.g. k = len(a)/2 == median; k = 1 == max(a)
func QuickSelect[T ordered](a []T, k int) T {
	var sol T
	Shuffle(a)
	left, right := 0, len(a)-1
	for left <= right {
		i := partition(a, left, right)
		if i < k {
			left = i+1
		} else if i > k {
			right = i-1
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
