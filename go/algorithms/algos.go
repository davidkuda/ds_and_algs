package algos

// --- SortingAlgorithms

// func selectionSort[T comparable](a []T) []T {
func selectionSort(a []int) {
	N := len(a)
	for i := 0; i < N; i++ {
		// identify min
		min := i
		for j := i+1; j < N; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		// exchange
		a[i], a[min] = a[min], a[i]
	}
}

// --- LinkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

// --- reverse linked list

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
